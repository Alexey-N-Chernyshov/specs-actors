package storage_market

import (
	addr "github.com/filecoin-project/go-address"

	abi "github.com/filecoin-project/specs-actors/actors/abi"
	big "github.com/filecoin-project/specs-actors/actors/abi/big"
	builtin "github.com/filecoin-project/specs-actors/actors/builtin"
	vmr "github.com/filecoin-project/specs-actors/actors/runtime"
	exitcode "github.com/filecoin-project/specs-actors/actors/runtime/exitcode"
	. "github.com/filecoin-project/specs-actors/actors/util"
	"github.com/filecoin-project/specs-actors/actors/util/adt"
)

type StorageMarketActor struct{}

type Runtime = vmr.Runtime

////////////////////////////////////////////////////////////////////////////////
// Actor methods
////////////////////////////////////////////////////////////////////////////////

type WithdrawBalanceParams struct {
	Address addr.Address
	Amount  abi.TokenAmount
}

// Attempt to withdraw the specified amount from the balance held in escrow.
// If less than the specified amount is available, yields the entire available balance.
func (a *StorageMarketActor) WithdrawBalance(rt Runtime, params *WithdrawBalanceParams) *adt.EmptyValue {
	amountSlashedTotal := abi.NewTokenAmount(0)

	if params.Amount.LessThan(big.Zero()) {
		rt.Abort(exitcode.ErrIllegalArgument, "negative amount %v", params.Amount)
	}

	recipientAddr := builtin.MarketAddress(rt, params.Address)

	var amountExtracted abi.TokenAmount
	var st StorageMarketActorState
	rt.State().Transaction(&st, func() interface{} {
		st.abortIfAddressEntryDoesNotExist(rt, params.Address)

		// Before any operations that check the balance tables for funds, execute all deferred
		// deal state updates.
		//
		// Note: as an optimization, implementations may cache efficient data structures indicating
		// which of the following set of updates are redundant and can be skipped.
		amountSlashedTotal = big.Add(amountSlashedTotal, st.updatePendingDealStatesForParty(rt, params.Address))

		minBalance := st.getLockedReqBalance(params.Address)
		newTable, ex, ok := BalanceTable_WithExtractPartial(st.EscrowTable, params.Address, params.Amount, minBalance)
		Assert(ok)

		st.EscrowTable = newTable
		amountExtracted = ex
		return nil
	})

	_, code := rt.Send(builtin.BurntFundsActorAddr, builtin.MethodSend, nil, amountSlashedTotal)
	builtin.RequireSuccess(rt, code, "failed to burn slashed funds")
	_, code = rt.Send(recipientAddr, builtin.MethodSend, nil, amountExtracted)
	builtin.RequireSuccess(rt, code, "failed to send funds")
	return &adt.EmptyValue{}
}

// Deposits the specified amount into the balance held in escrow.
// Note: the amount is included implicitly in the message.
func (a *StorageMarketActor) AddBalance(rt Runtime, address *addr.Address) *adt.EmptyValue {
	builtin.MarketAddress(rt, *address)

	var st StorageMarketActorState
	rt.State().Transaction(&st, func() interface{} {
		st.abortIfAddressEntryDoesNotExist(rt, *address)

		msgValue := rt.ValueReceived()
		newTable, ok := BalanceTable_WithAdd(st.EscrowTable, *address, msgValue)
		if !ok {
			// Entry not found; create implicitly.
			newTable, ok = BalanceTable_WithNewAddressEntry(st.EscrowTable, *address, msgValue)
			Assert(ok)
		}
		st.EscrowTable = newTable
		return nil
	})
	return &adt.EmptyValue{}
}

type PublishStorageDealsParams struct {
	Deals []StorageDealProposal
}

// Publish a new set of storage deals (not yet included in a sector).
func (a *StorageMarketActor) PublishStorageDeals(rt Runtime, params *PublishStorageDealsParams) *adt.EmptyValue {
	amountSlashedTotal := abi.NewTokenAmount(0)

	// Deal message must have a From field identical to the provider of all the deals.
	// This allows us to retain and verify only the client's signature in each deal proposal itself.
	rt.ValidateImmediateCallerType(builtin.CallerTypesSignable...)

	var st StorageMarketActorState
	rt.State().Transaction(&st, func() interface{} {
		// All storage deals will be added in an atomic transaction; this operation will be unrolled if any of them fails.
		for _, deal := range params.Deals {
			if deal.Provider != rt.ImmediateCaller() {
				rt.Abort(exitcode.ErrForbidden, "caller is not provider %v", deal.Provider)
			}

			validateDeal(rt, deal)

			// Before any operations that check the balance tables for funds, execute all deferred
			// deal state updates.
			//
			// Note: as an optimization, implementations may cache efficient data structures indicating
			// which of the following set of updates are redundant and can be skipped.
			amountSlashedTotal = big.Add(amountSlashedTotal, st.updatePendingDealStatesForParty(rt, deal.Client))
			amountSlashedTotal = big.Add(amountSlashedTotal, st.updatePendingDealStatesForParty(rt, deal.Provider))

			st.lockBalanceOrAbort(rt, deal.Client, deal.ClientBalanceRequirement())
			st.lockBalanceOrAbort(rt, deal.Provider, deal.ProviderBalanceRequirement())

			id := st.generateStorageDealID()

			onchainDeal := OnChainDeal{
				ID:               id,
				Proposal:         deal,
				SectorStartEpoch: epochUndefined,
			}

			st.Deals[id] = onchainDeal
		}

		return nil
	})

	_, code := rt.Send(builtin.BurntFundsActorAddr, builtin.MethodSend, nil, amountSlashedTotal)
	builtin.RequireSuccess(rt, code, "failed to burn funds")
	return &adt.EmptyValue{}
}

type VerifyDealsOnSectorProveCommitParams struct {
	DealIDs      []abi.DealID
	SectorExpiry abi.ChainEpoch
}

// Verify that a given set of storage deals is valid for a sector currently being ProveCommitted,
// update the market's internal state accordingly, and return DealWeight of the set of storage deals given.
// Note: in the case of a capacity-commitment sector (one with zero deals), this function should succeed vacuously.
// The weight is defined as the sum, over all deals in the set, of the product of its size
// with its duration. This quantity may be an input into the functions specifying block reward,
// sector power, collateral, and/or other parameters.
func (a *StorageMarketActor) VerifyDealsOnSectorProveCommit(rt Runtime, params *VerifyDealsOnSectorProveCommitParams) *abi.DealWeight {
	rt.ValidateImmediateCallerType(builtin.StorageMinerActorCodeID)
	minerAddr := rt.ImmediateCaller()
	totalWeight := big.Zero()

	var st StorageMarketActorState
	rt.State().Transaction(&st, func() interface{} {
		// if there are no dealIDs, it is a CommittedCapacity sector
		// and the totalWeight should be zero
		for _, dealID := range params.DealIDs {
			deal, dealP := st.getOnChainDealOrAbort(rt, dealID)
			validateDealCanActivate(rt, minerAddr, params.SectorExpiry, deal)
			ocd := st.Deals[dealID]
			ocd.SectorStartEpoch = rt.CurrEpoch()
			st.Deals[dealID] = ocd

			// Compute deal weight
			dur := big.NewInt(int64(dealP.Duration()))
			siz := big.NewInt(dealP.PieceSize.Total())
			weight := big.Mul(dur, siz)
			totalWeight = big.Add(totalWeight, weight)
		}
		return nil
	})
	return &totalWeight
}

type GetPieceInfosForDealIDsParams struct {
	DealIDs []abi.DealID
}

type GetPieceInfosForDealIDsReturn struct {
	Pieces []abi.PieceInfo
}

func (a *StorageMarketActor) GetPieceInfosForDealIDs(rt Runtime, params *GetPieceInfosForDealIDsParams) *GetPieceInfosForDealIDsReturn {
	rt.ValidateImmediateCallerType(builtin.StorageMinerActorCodeID)

	var ret []abi.PieceInfo
	var st StorageMarketActorState
	rt.State().Transaction(&st, func() interface{} {
		for _, dealID := range params.DealIDs {
			_, dealP := st.getOnChainDealOrAbort(rt, dealID)
			ret = append(ret, abi.PieceInfo{
				PieceCID: dealP.PieceCID,
				Size:     dealP.PieceSize.Total(),
			})
		}
		return nil
	})

	return &GetPieceInfosForDealIDsReturn{Pieces: ret}
}

type OnMinerSectorsTerminateParams struct {
	DealIDs []abi.DealID
}

// Terminate a set of deals in response to their containing sector being terminated.
// Slash provider collateral, refund client collateral, and refund partial unpaid escrow
// amount to client.
func (a *StorageMarketActor) OnMinerSectorsTerminate(rt Runtime, params *OnMinerSectorsTerminateParams) *adt.EmptyValue {
	rt.ValidateImmediateCallerType(builtin.StorageMinerActorCodeID)
	minerAddr := rt.ImmediateCaller()

	var st StorageMarketActorState
	rt.State().Transaction(&st, func() interface{} {
		for _, dealID := range params.DealIDs {
			_, dealP := st.getOnChainDealOrAbort(rt, dealID)
			Assert(dealP.Provider == minerAddr)

			// Note: we do not perform the balance transfers here, but rather simply record the flag
			// to indicate that processDealSlashed should be called when the deferred state computation
			// is performed.
			ocd := st.Deals[dealID]
			ocd.SlashEpoch = rt.CurrEpoch()
			st.Deals[dealID] = ocd
		}
		return nil
	})
	return &adt.EmptyValue{}
}

type HandleExpiredDealsParams struct {
	Deals []abi.DealID // TODO: RLE
}

func (a *StorageMarketActor) HandleExpiredDeals(rt Runtime, params *HandleExpiredDealsParams) *adt.EmptyValue {
	rt.ValidateImmediateCallerType(builtin.CallerTypesSignable...)
	var slashed abi.TokenAmount
	var st StorageMarketActorState
	rt.State().Transaction(&st, func() interface{} {
		slashed = st.updatePendingDealStates(params.Deals, rt.CurrEpoch())
		return nil
	})

	// TODO: award some small portion of slashed to caller as incentive

	_, code := rt.Send(builtin.BurntFundsActorAddr, builtin.MethodSend, nil, slashed)
	builtin.RequireSuccess(rt, code, "failed to burn funds")
	return &adt.EmptyValue{}
}

func (a *StorageMarketActor) Constructor(rt Runtime, _ *adt.EmptyValue) *adt.EmptyValue {
	rt.ValidateImmediateCallerIs(builtin.SystemActorAddr)
	var st StorageMarketActorState
	rt.State().Transaction(&st, func() interface{} {
		st.Deals = DealsAMT_Empty()
		st.EscrowTable = BalanceTableHAMT_Empty()
		st.LockedTable = BalanceTableHAMT_Empty()
		st.NextID = abi.DealID(0)
		st.DealIDsByParty = CachedDealIDsByPartyHAMT_Empty()
		return nil
	})
	return &adt.EmptyValue{}
}

////////////////////////////////////////////////////////////////////////////////
// Checks
////////////////////////////////////////////////////////////////////////////////

func validateDealCanActivate(rt Runtime, minerAddr addr.Address, sectorExpiration abi.ChainEpoch, deal OnChainDeal) {
	if deal.Proposal.Provider != minerAddr {
		rt.Abort(exitcode.ErrIllegalArgument, "Deal has incorrect miner as its provider.")
	}

	if deal.SectorStartEpoch != epochUndefined {
		rt.Abort(exitcode.ErrIllegalArgument, "Deal has already appeared in proven sector.")
	}

	if rt.CurrEpoch() > deal.Proposal.StartEpoch {
		rt.Abort(exitcode.ErrIllegalArgument, "Deal start epoch has already elapsed.")
	}

	if deal.Proposal.EndEpoch > sectorExpiration {
		rt.Abort(exitcode.ErrIllegalArgument, "Deal would outlive its containing sector.")
	}
}

func validateDeal(rt Runtime, deal StorageDealProposal) {
	if !dealProposalIsInternallyValid(rt, deal) {
		rt.Abort(exitcode.ErrIllegalArgument, "Invalid deal proposal.")
	}

	if rt.CurrEpoch() > deal.StartEpoch {
		rt.Abort(exitcode.ErrIllegalArgument, "Deal start epoch has already elapsed.")
	}

	inds := rt.CurrIndices()

	minDuration, maxDuration := inds.StorageDeal_DurationBounds(deal.PieceSize, deal.StartEpoch)
	if deal.Duration() < minDuration || deal.Duration() > maxDuration {
		rt.Abort(exitcode.ErrIllegalArgument, "Deal duration out of bounds.")
	}

	minPrice, maxPrice := inds.StorageDeal_StoragePricePerEpochBounds(deal.PieceSize, deal.StartEpoch, deal.EndEpoch)
	if deal.StoragePricePerEpoch.LessThan(minPrice) || deal.StoragePricePerEpoch.GreaterThan(maxPrice) {
		rt.Abort(exitcode.ErrIllegalArgument, "Storage price out of bounds.")
	}

	minProviderCollateral, maxProviderCollateral := inds.StorageDeal_ProviderCollateralBounds(deal.PieceSize, deal.StartEpoch, deal.EndEpoch)
	if deal.ProviderCollateral.LessThan(minProviderCollateral) || deal.ProviderCollateral.GreaterThan(maxProviderCollateral) {
		rt.Abort(exitcode.ErrIllegalArgument, "Provider collateral out of bounds.")
	}

	minClientCollateral, maxClientCollateral := inds.StorageDeal_ClientCollateralBounds(deal.PieceSize, deal.StartEpoch, deal.EndEpoch)
	if deal.ClientCollateral.LessThan(minClientCollateral) || deal.ClientCollateral.GreaterThan(maxClientCollateral) {
		rt.Abort(exitcode.ErrIllegalArgument, "Client collateral out of bounds.")
	}
}
