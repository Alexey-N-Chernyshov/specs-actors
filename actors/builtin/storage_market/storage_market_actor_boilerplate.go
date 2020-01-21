package storage_market

import (
	abi "github.com/filecoin-project/specs-actors/actors/abi"
	vmr "github.com/filecoin-project/specs-actors/actors/runtime"
	autil "github.com/filecoin-project/specs-actors/actors/util"
)

type BalanceTableHAMT = autil.BalanceTableHAMT
type DealIDQueue = autil.DealIDQueue

////////////////////////////////////////////////////////////////////////////////
// Boilerplate
//
// This boilerplate should be essentially identical for all actors, and
// conceptually belongs in the runtime/VM. It is only duplicated here as a
// workaround due to the lack of generics support in Go.
////////////////////////////////////////////////////////////////////////////////

type Runtime = vmr.Runtime
type Bytes = abi.Bytes

var Assert = autil.Assert
var IMPL_FINISH = autil.IMPL_FINISH

func Release(rt Runtime, h vmr.ActorStateHandle, st StorageMarketActorState) {
	checkCID := abi.ActorSubstateCID(rt.IpldPut(&st))
	h.Release(checkCID)
}
func UpdateRelease(rt Runtime, h vmr.ActorStateHandle, st StorageMarketActorState) {
	newCID := abi.ActorSubstateCID(rt.IpldPut(&st))
	h.UpdateRelease(newCID)
}
func DealsAMT_Empty() DealsAMT {
	IMPL_FINISH()
	panic("")
}

func CachedDealIDsByPartyHAMT_Empty() CachedDealIDsByPartyHAMT {
	IMPL_FINISH()
	panic("")
}

func CachedExpirationsPendingHAMT_Empty() CachedExpirationsPendingHAMT {
	IMPL_FINISH()
	panic("")
}
