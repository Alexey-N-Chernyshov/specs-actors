package abi

import (
	"github.com/filecoin-project/specs-actors/actors/abi/big"
)

// The abi package contains definitions of all types that cross the VM boundary and are used
// within actor code.
//
// Primitive types include numerics and opaque array types.

// Epoch number of the chain state, which acts as a proxy for time within the VM.
type ChainEpoch int64

// A sequential number assigned to an actor when created by the InitActor.
// This ID is embedded in ID-type addresses.
type ActorID int64

// MethodNum is an integer that represents a particular method
// in an actor's function table. These numbers are used to compress
// invocation of actor code, and to decouple human language concerns
// about method names from the ability to uniquely refer to a particular
// method.
//
// Consider MethodNum numbers to be similar in concerns as for
// offsets in function tables (in programming languages), and for
// tags in ProtocolBuffer fields. Tags in ProtocolBuffers recommend
// assigning a unique tag to a field and never reusing that tag.
// If a field is no longer used, the field name may change but should
// still remain defined in the code to ensure the tag number is not
// reused accidentally. The same should apply to the MethodNum
// associated with methods in Filecoin VM Actors.
type MethodNum int64

// TokenAmount is an amount of Filecoin tokens. This type is used within
// the VM in message execution, to account movement of tokens, payment
// of VM gas, and more.
//
// BigInt types are aliases rather than new types because the latter introduce incredible amounts of noise converting to
// and from types in order to manipulate values. We give up some type safety for ergonomics.
type TokenAmount = big.Int

func NewTokenAmount(t int64) TokenAmount {
	return big.NewInt(t)
}

// The randomness seed is a string of byte, distinguished from Randomness
// for expressiveness: it hasn't been given the needed entropy
type RandomnessSeed []byte

// Randomness is a string of random bytes
type Randomness []byte
