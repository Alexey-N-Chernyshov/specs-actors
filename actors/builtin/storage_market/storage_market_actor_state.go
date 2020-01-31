package storage_market

import (
	"io"

	addr "github.com/filecoin-project/go-address"

	abi "github.com/filecoin-project/specs-actors/actors/abi"
	big "github.com/filecoin-project/specs-actors/actors/abi/big"
	exitcode "github.com/filecoin-project/specs-actors/actors/runtime/exitcode"
	indices "github.com/filecoin-project/specs-actors/actors/runtime/indices"
	. "github.com/filecoin-project/specs-actors/actors/util"
	actorutil "github.com/filecoin-project/specs-actors/actors/util"
)

const epochUndefined = abi.ChainEpoch(-1)

// TODO H/AMT, depending on sparseness and mutation pattern.
type DealsById map[abi.DealID]OnChainDeal

// TODO HAMT
type DealsByParty map[addr.Address]DealIDSet
type DealIDSet map[abi.DealID]bool // TODO: HAMT or bitfield, depending how often it mutates

// TODO HAMT (probably of AMTs, i.e. a multimap).
type DealExpirationQueue map[abi.ChainEpoch]DealIDQueue

type StorageMarketActorState struct {
	Deals DealsById

	// Total amount held in escrow, indexed by actor address (including both locked and unlocked amounts).
	EscrowTable actorutil.BalanceTableHAMT

	// Amount locked, indexed by actor address.
	// Note: the amounts in this table do not affect the overall amount in escrow:
	// only the _portion_ of the total escrow amount that is locked.
	LockedReqTable actorutil.BalanceTableHAMT

	NextID abi.DealID

	// Metadata cached for efficient iteration over deals.
	DealIDsByParty             DealsByParty
	ExpirationsPending         DealExpirationQueue
	ExpirationsNextProcEpoch   abi.ChainEpoch
	CurrEpochNumDealsPublished int
}

func (st *StorageMarketActorState) UnmarshalCBOR(r io.Reader) error {
	panic("replace with cbor-gen")
}

func (st *StorageMarketActorState) MarshalCBOR(w io.Writer) error {
	panic("replace with cbor-gen")
}

////////////////////////////////////////////////////////////////////////////////
// Deal state operations
////////////////////////////////////////////////////////////////////////////////

func (st *StorageMarketActorState) updatePendingDealStates(dealIDs []abi.DealID, epoch abi.ChainEpoch) (
	amountSlashedTotal abi.TokenAmount) {

	amountSlashedTotal = abi.NewTokenAmount(0)

	for _, dealID := range dealIDs {
		amountSlashedCurr := st.updatePendingDealState(dealID, epoch)
		amountSlashedTotal = big.Add(amountSlashedTotal, amountSlashedCurr)
	}

	return
}

func (st *StorageMarketActorState) updatePendingDealState(dealID abi.DealID, epoch abi.ChainEpoch) (
	amountSlashed abi.TokenAmount) {

	amountSlashed = abi.NewTokenAmount(0)

	deal, dealP := st.getOnChainDealAssert(dealID)

	everUpdated := deal.LastUpdatedEpoch != epochUndefined
	everSlashed := deal.SlashEpoch != epochUndefined

	Assert(!everUpdated || (deal.LastUpdatedEpoch <= epoch))
	if deal.LastUpdatedEpoch == epoch {
		return
	}

	if deal.SectorStartEpoch == epochUndefined {
		// Not yet appeared in proven sector; check for timeout.
		if dealP.StartEpoch >= epoch {
			st.processDealInitTimedOut(dealID)
		}
		return
	}

	Assert(dealP.StartEpoch <= epoch)

	dealEnd := dealP.EndEpoch
	if everSlashed {
		Assert(deal.SlashEpoch <= dealEnd)
		dealEnd = deal.SlashEpoch
	}

	elapsedStart := dealP.StartEpoch
	if everUpdated && deal.LastUpdatedEpoch > elapsedStart {
		elapsedStart = deal.LastUpdatedEpoch
	}

	elapsedEnd := dealEnd
	if epoch < elapsedEnd {
		elapsedEnd = epoch
	}

	numEpochsElapsed := elapsedEnd - elapsedStart
	st._processDealPaymentEpochsElapsed(dealID, numEpochsElapsed)

	if everSlashed {
		amountSlashed = st.processDealSlashed(dealID)
		return
	}

	if epoch >= dealP.EndEpoch {
		st.processDealExpired(dealID)
		return
	}

	ocd := st.Deals[dealID]
	ocd.LastUpdatedEpoch = epoch
	st.Deals[dealID] = ocd
	return
}

func (st *StorageMarketActorState) _deleteDeal(dealID abi.DealID) {
	_, dealP := st.getOnChainDealAssert(dealID)
	delete(st.Deals, dealID)
	delete(st.DealIDsByParty[dealP.Provider], dealID)
	delete(st.DealIDsByParty[dealP.Client], dealID)
}

// Note: only processes deal payments, not deal expiration (even if the deal has expired).
func (st *StorageMarketActorState) _processDealPaymentEpochsElapsed(dealID abi.DealID, numEpochsElapsed abi.ChainEpoch) {
	deal, dealP := st.getOnChainDealAssert(dealID)
	Assert(deal.SectorStartEpoch != epochUndefined)

	// Process deal payment for the elapsed epochs.
	totalPayment := big.Mul(big.NewInt(int64(numEpochsElapsed)), dealP.StoragePricePerEpoch)
	st.transferBalance(dealP.Client, dealP.Provider, abi.TokenAmount(totalPayment))
}

func (st *StorageMarketActorState) processDealSlashed(dealID abi.DealID) (amountSlashed abi.TokenAmount) {
	deal, dealP := st.getOnChainDealAssert(dealID)
	Assert(deal.SectorStartEpoch != epochUndefined)

	slashEpoch := deal.SlashEpoch
	Assert(slashEpoch != epochUndefined)

	// unlock client collateral and locked storage fee
	clientCollateral := dealP.ClientCollateral
	paymentRemaining := dealGetPaymentRemaining(deal, slashEpoch)
	st.unlockBalance(dealP.Client, big.Add(clientCollateral, paymentRemaining))

	// slash provider collateral
	amountSlashed = dealP.ProviderCollateral
	st.slashBalance(dealP.Provider, amountSlashed)

	st._deleteDeal(dealID)
	return
}

// Deal start deadline elapsed without appearing in a proven sector.
// Delete deal, slash a portion of provider's collateral, and unlock remaining collaterals
// for both provider and client.
func (st *StorageMarketActorState) processDealInitTimedOut(dealID abi.DealID) (amountSlashed abi.TokenAmount) {
	deal, dealP := st.getOnChainDealAssert(dealID)
	Assert(deal.SectorStartEpoch == epochUndefined)

	st.unlockBalance(dealP.Client, dealP.ClientBalanceRequirement())

	amountSlashed = indices.StorageDeal_ProviderInitTimedOutSlashAmount(deal.Deal.Proposal.ProviderCollateral)
	amountRemaining := big.Sub(dealP.ProviderBalanceRequirement(), amountSlashed)

	st.slashBalance(dealP.Provider, amountSlashed)
	st.unlockBalance(dealP.Provider, amountRemaining)

	st._deleteDeal(dealID)
	return
}

// Normal expiration. Delete deal and unlock collaterals for both miner and client.
func (st *StorageMarketActorState) processDealExpired(dealID abi.DealID) {
	deal, dealP := st.getOnChainDealAssert(dealID)
	Assert(deal.SectorStartEpoch != epochUndefined)

	// Note: payment has already been completed at this point (_rtProcessDealPaymentEpochsElapsed)
	st.unlockBalance(dealP.Provider, dealP.ProviderCollateral)
	st.unlockBalance(dealP.Client, dealP.ClientCollateral)

	st._deleteDeal(dealID)
}

func (st *StorageMarketActorState) generateStorageDealID() abi.DealID {
	ret := st.NextID
	st.NextID = st.NextID + abi.DealID(1)
	return ret
}

////////////////////////////////////////////////////////////////////////////////
// Balance table operations
////////////////////////////////////////////////////////////////////////////////

func (st *StorageMarketActorState) addressEntryExists(address addr.Address) bool { // TODO: This is expensive, drop where not needed
	_, foundEscrow := actorutil.BalanceTable_GetEntry(st.EscrowTable, address)
	_, foundLocked := actorutil.BalanceTable_GetEntry(st.LockedReqTable, address)
	// Check that the tables are consistent (i.e. the address is found in one
	// if and only if it is found in the other).
	Assert(foundEscrow == foundLocked)
	return foundEscrow
}

func (st *StorageMarketActorState) getTotalEscrowBalance(a addr.Address) abi.TokenAmount {
	Assert(st.addressEntryExists(a))
	ret, ok := actorutil.BalanceTable_GetEntry(st.EscrowTable, a)
	Assert(ok)
	return ret
}

func (st *StorageMarketActorState) getLockedReqBalance(a addr.Address) abi.TokenAmount {
	Assert(st.addressEntryExists(a))
	ret, ok := actorutil.BalanceTable_GetEntry(st.LockedReqTable, a)
	Assert(ok)
	return ret
}

func (st *StorageMarketActorState) lockBalanceMaybe(addr addr.Address, amount abi.TokenAmount) (lockBalanceOK bool) {
	Assert(amount.GreaterThanEqual(big.Zero()))
	Assert(st.addressEntryExists(addr))

	prevLocked := st.getLockedReqBalance(addr)
	if big.Add(prevLocked, amount).GreaterThan(st.getTotalEscrowBalance(addr)) {
		lockBalanceOK = false
		return
	}

	newLockedReqTable, ok := actorutil.BalanceTable_WithAdd(st.LockedReqTable, addr, amount)
	Assert(ok)
	st.LockedReqTable = newLockedReqTable

	lockBalanceOK = true
	return
}

func (st *StorageMarketActorState) unlockBalance(addr addr.Address, unlockAmountRequested abi.TokenAmount) {
	Assert(unlockAmountRequested.GreaterThanEqual(big.Zero()))
	Assert(st.addressEntryExists(addr))

	st.LockedReqTable = st.tableWithDeductBalanceExact(st.LockedReqTable, addr, unlockAmountRequested)
}

func (st *StorageMarketActorState) tableWithAddBalance(table actorutil.BalanceTableHAMT, toAddr addr.Address, amountToAdd abi.TokenAmount) actorutil.BalanceTableHAMT {

	Assert(amountToAdd.GreaterThanEqual(big.Zero()))

	newTable, ok := actorutil.BalanceTable_WithAdd(table, toAddr, amountToAdd)
	Assert(ok)
	return newTable
}

func (st *StorageMarketActorState) tableWithDeductBalanceExact(table actorutil.BalanceTableHAMT, fromAddr addr.Address, amountRequested abi.TokenAmount) actorutil.BalanceTableHAMT {
	Assert(amountRequested.GreaterThanEqual(big.Zero()))

	newTable, amountDeducted, ok := actorutil.BalanceTable_WithSubtractPreservingNonnegative(
		table, fromAddr, amountRequested)
	Assert(ok)
	Assert(amountDeducted == amountRequested)
	return newTable
}

// move funds from locked in client to available in provider
func (st *StorageMarketActorState) transferBalance(fromAddr addr.Address, toAddr addr.Address, transferAmountRequested abi.TokenAmount) {

	Assert(transferAmountRequested.GreaterThanEqual(big.Zero()))
	Assert(st.addressEntryExists(fromAddr))
	Assert(st.addressEntryExists(toAddr))

	st.EscrowTable = st.tableWithDeductBalanceExact(st.EscrowTable, fromAddr, transferAmountRequested)
	st.LockedReqTable = st.tableWithDeductBalanceExact(st.LockedReqTable, fromAddr, transferAmountRequested)
	st.EscrowTable = st.tableWithAddBalance(st.EscrowTable, toAddr, transferAmountRequested)
}

func (st *StorageMarketActorState) slashBalance(addr addr.Address, slashAmount abi.TokenAmount) {
	Assert(st.addressEntryExists(addr))
	Assert(slashAmount.GreaterThanEqual(big.Zero()))

	st.EscrowTable = st.tableWithDeductBalanceExact(st.EscrowTable, addr, slashAmount)
	st.LockedReqTable = st.tableWithDeductBalanceExact(st.LockedReqTable, addr, slashAmount)
}

////////////////////////////////////////////////////////////////////////////////
// Method utility functions
////////////////////////////////////////////////////////////////////////////////

func (st *StorageMarketActorState) abortIfAddressEntryDoesNotExist(rt Runtime, entryAddr addr.Address) {
	if !st.addressEntryExists(entryAddr) {
		rt.Abort(exitcode.ErrNotFound, "no entry for %v", entryAddr)
	}
}

func (st *StorageMarketActorState) updatePendingDealStatesForParty(rt Runtime, addr addr.Address) (amountSlashedTotal abi.TokenAmount) {
	// For consistency with OnEpochTickEnd, only process updates up to the end of the _previous_ epoch.
	epoch := rt.CurrEpoch() - 1

	cachedRes, ok := st.DealIDsByParty[addr]
	Assert(ok)
	var extractedDealIDs []abi.DealID
	for cachedDealID := range cachedRes {
		extractedDealIDs = append(extractedDealIDs, cachedDealID)
	}

	amountSlashedTotal = st.updatePendingDealStates(extractedDealIDs, epoch)
	return
}

func (st *StorageMarketActorState) getOnChainDealOrAbort(rt Runtime, dealID abi.DealID) (deal OnChainDeal, dealP StorageDealProposal) {
	var found bool
	deal, dealP, found = st.getOnChainDeal(dealID)
	if !found {
		rt.Abort(exitcode.ErrNotFound, "dealID not found in Deals.")
	}
	return
}

func (st *StorageMarketActorState) lockBalanceOrAbort(rt Runtime, addr addr.Address, amount abi.TokenAmount) {
	if amount.LessThan(big.Zero()) {
		rt.Abort(exitcode.ErrIllegalArgument, "negative amount %v", amount)
	}

	st.abortIfAddressEntryDoesNotExist(rt, addr)

	if !st.lockBalanceMaybe(addr, amount) {
		rt.Abort(exitcode.ErrInsufficientFunds, "Insufficient funds available to lock")
	}
}

////////////////////////////////////////////////////////////////////////////////
// State utility functions
////////////////////////////////////////////////////////////////////////////////

func dealProposalIsInternallyValid(rt Runtime, dealP StorageDealProposal) bool {
	if dealP.EndEpoch <= dealP.StartEpoch {
		return false
	}

	if dealP.Duration() != dealP.EndEpoch-dealP.StartEpoch {
		return false
	}

	IMPL_FINISH()
	// Determine which subset of DealProposal to use as the message to be signed by the client.
	var m []byte

	// Note: we do not verify the provider signature here, since this is implicit in the
	// authenticity of the on-chain message publishing the deal.
	sigVerified := rt.Syscalls().VerifySignature(dealP.ClientSignature, dealP.Client, m)
	if !sigVerified {
		return false
	}

	return true
}

func dealGetPaymentRemaining(deal OnChainDeal, epoch abi.ChainEpoch) abi.TokenAmount {
	dealP := deal.Deal.Proposal
	Assert(epoch <= dealP.EndEpoch)

	durationRemaining := dealP.EndEpoch - (epoch - 1)
	Assert(durationRemaining > 0)

	return big.Mul(big.NewInt(int64(durationRemaining)), dealP.StoragePricePerEpoch)
}

func (st *StorageMarketActorState) getOnChainDeal(dealID abi.DealID) (deal OnChainDeal, dealP StorageDealProposal, ok bool) {
	deal, ok = st.Deals[dealID]
	if ok {
		dealP = deal.Deal.Proposal
	}
	return
}

func (st *StorageMarketActorState) getOnChainDealAssert(dealID abi.DealID) (deal OnChainDeal, dealP StorageDealProposal) {
	var ok bool
	deal, dealP, ok = st.getOnChainDeal(dealID)
	Assert(ok)
	return
}

///// DealIDQueue /////
// TODO: replace with an AMT
type DealIDQueue struct {
	Values     IndexedDealIDs
	StartIndex int64
	EndIndex   int64
}

type IndexedDealIDs map[int64]abi.DealID

func (x *DealIDQueue) Enqueue(dealID abi.DealID) {
	nextIndex := x.EndIndex
	x.Values[nextIndex] = dealID
	x.EndIndex = nextIndex + 1
}

func (x *DealIDQueue) Dequeue() (dealID abi.DealID, ok bool) {
	actorutil.AssertMsg(x.StartIndex <= x.EndIndex, "index %d > end %d", x.StartIndex, x.EndIndex)

	if x.StartIndex == x.EndIndex {
		dealID = abi.DealID(-1)
		ok = false
		return
	} else {
		dealID = x.Values[x.StartIndex]
		delete(x.Values, x.StartIndex)
		x.StartIndex += 1
		ok = true
		return
	}
}

func NewDealIDQueue() DealIDQueue {
	return DealIDQueue{
		Values:     make(IndexedDealIDs),
		StartIndex: 0,
		EndIndex:   0,
	}
}

func DealsAMT_Empty() DealsById {
	IMPL_FINISH()
	panic("")
}

func CachedDealIDsByPartyHAMT_Empty() DealsByParty {
	IMPL_FINISH()
	panic("")
}

func CachedExpirationsPendingHAMT_Empty() DealExpirationQueue {
	IMPL_FINISH()
	panic("")
}
