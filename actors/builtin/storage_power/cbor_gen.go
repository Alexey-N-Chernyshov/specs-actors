// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package storage_power

import (
	"fmt"
	"io"

	abi "github.com/filecoin-project/specs-actors/actors/abi"
	builtin "github.com/filecoin-project/specs-actors/actors/builtin"
	peer "github.com/libp2p/go-libp2p-core/peer"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf

func (t *AddBalanceParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{129}); err != nil {
		return err
	}

	// t.Miner (address.Address) (struct)
	if err := t.Miner.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *AddBalanceParams) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Miner (address.Address) (struct)

	{

		if err := t.Miner.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	return nil
}

func (t *CreateMinerParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{131}); err != nil {
		return err
	}

	// t.Worker (address.Address) (struct)
	if err := t.Worker.MarshalCBOR(w); err != nil {
		return err
	}

	// t.SectorSize (abi.SectorSize) (int64)
	if t.SectorSize >= 0 {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.SectorSize))); err != nil {
			return err
		}
	} else {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajNegativeInt, uint64(-t.SectorSize)-1)); err != nil {
			return err
		}
	}

	// t.Peer (peer.ID) (string)
	if len(t.Peer) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Peer was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len(t.Peer)))); err != nil {
		return err
	}
	if _, err := w.Write([]byte(t.Peer)); err != nil {
		return err
	}
	return nil
}

func (t *CreateMinerParams) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Worker (address.Address) (struct)

	{

		if err := t.Worker.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.SectorSize (abi.SectorSize) (int64)
	{
		maj, extra, err := cbg.CborReadHeader(br)
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.SectorSize = abi.SectorSize(extraI)
	}
	// t.Peer (peer.ID) (string)

	{
		sval, err := cbg.ReadString(br)
		if err != nil {
			return err
		}

		t.Peer = peer.ID(sval)
	}
	return nil
}

func (t *DeleteMinerParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{129}); err != nil {
		return err
	}

	// t.Miner (address.Address) (struct)
	if err := t.Miner.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *DeleteMinerParams) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Miner (address.Address) (struct)

	{

		if err := t.Miner.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	return nil
}

func (t *WithdrawBalanceParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{130}); err != nil {
		return err
	}

	// t.Miner (address.Address) (struct)
	if err := t.Miner.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Requested (big.Int) (struct)
	if err := t.Requested.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *WithdrawBalanceParams) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Miner (address.Address) (struct)

	{

		if err := t.Miner.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.Requested (big.Int) (struct)

	{

		if err := t.Requested.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	return nil
}

func (t *EnrollCronEventParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{130}); err != nil {
		return err
	}

	// t.EventEpoch (abi.ChainEpoch) (int64)
	if t.EventEpoch >= 0 {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.EventEpoch))); err != nil {
			return err
		}
	} else {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajNegativeInt, uint64(-t.EventEpoch)-1)); err != nil {
			return err
		}
	}

	// t.Payload ([]uint8) (slice)
	if len(t.Payload) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Payload was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajByteString, uint64(len(t.Payload)))); err != nil {
		return err
	}
	if _, err := w.Write(t.Payload); err != nil {
		return err
	}
	return nil
}

func (t *EnrollCronEventParams) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.EventEpoch (abi.ChainEpoch) (int64)
	{
		maj, extra, err := cbg.CborReadHeader(br)
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.EventEpoch = abi.ChainEpoch(extraI)
	}
	// t.Payload ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Payload: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}
	t.Payload = make([]byte, extra)
	if _, err := io.ReadFull(br, t.Payload); err != nil {
		return err
	}
	return nil
}

func (t *OnSectorTerminateParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{130}); err != nil {
		return err
	}

	// t.TerminationType (builtin.SectorTermination) (int64)
	if t.TerminationType >= 0 {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.TerminationType))); err != nil {
			return err
		}
	} else {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajNegativeInt, uint64(-t.TerminationType)-1)); err != nil {
			return err
		}
	}

	// t.Weight (storage_power.SectorStorageWeightDesc) (struct)
	if err := t.Weight.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *OnSectorTerminateParams) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.TerminationType (builtin.SectorTermination) (int64)
	{
		maj, extra, err := cbg.CborReadHeader(br)
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.TerminationType = builtin.SectorTermination(extraI)
	}
	// t.Weight (storage_power.SectorStorageWeightDesc) (struct)

	{

		if err := t.Weight.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	return nil
}

func (t *OnSectorModifyWeightDescParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{130}); err != nil {
		return err
	}

	// t.PrevWeight (storage_power.SectorStorageWeightDesc) (struct)
	if err := t.PrevWeight.MarshalCBOR(w); err != nil {
		return err
	}

	// t.NewWeight (storage_power.SectorStorageWeightDesc) (struct)
	if err := t.NewWeight.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *OnSectorModifyWeightDescParams) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.PrevWeight (storage_power.SectorStorageWeightDesc) (struct)

	{

		if err := t.PrevWeight.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.NewWeight (storage_power.SectorStorageWeightDesc) (struct)

	{

		if err := t.NewWeight.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	return nil
}

func (t *OnSectorProveCommitParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{129}); err != nil {
		return err
	}

	// t.Weight (storage_power.SectorStorageWeightDesc) (struct)
	if err := t.Weight.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *OnSectorProveCommitParams) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Weight (storage_power.SectorStorageWeightDesc) (struct)

	{

		if err := t.Weight.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	return nil
}

func (t *ReportConsensusFaultParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{133}); err != nil {
		return err
	}

	// t.BlockHeader1 ([]uint8) (slice)
	if len(t.BlockHeader1) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.BlockHeader1 was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajByteString, uint64(len(t.BlockHeader1)))); err != nil {
		return err
	}
	if _, err := w.Write(t.BlockHeader1); err != nil {
		return err
	}

	// t.BlockHeader2 ([]uint8) (slice)
	if len(t.BlockHeader2) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.BlockHeader2 was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajByteString, uint64(len(t.BlockHeader2)))); err != nil {
		return err
	}
	if _, err := w.Write(t.BlockHeader2); err != nil {
		return err
	}

	// t.Target (address.Address) (struct)
	if err := t.Target.MarshalCBOR(w); err != nil {
		return err
	}

	// t.FaultEpoch (abi.ChainEpoch) (int64)
	if t.FaultEpoch >= 0 {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.FaultEpoch))); err != nil {
			return err
		}
	} else {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajNegativeInt, uint64(-t.FaultEpoch)-1)); err != nil {
			return err
		}
	}

	// t.FaultType (storage_power.ConsensusFaultType) (int64)
	if t.FaultType >= 0 {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.FaultType))); err != nil {
			return err
		}
	} else {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajNegativeInt, uint64(-t.FaultType)-1)); err != nil {
			return err
		}
	}
	return nil
}

func (t *ReportConsensusFaultParams) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 5 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.BlockHeader1 ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.BlockHeader1: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}
	t.BlockHeader1 = make([]byte, extra)
	if _, err := io.ReadFull(br, t.BlockHeader1); err != nil {
		return err
	}
	// t.BlockHeader2 ([]uint8) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.BlockHeader2: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}
	t.BlockHeader2 = make([]byte, extra)
	if _, err := io.ReadFull(br, t.BlockHeader2); err != nil {
		return err
	}
	// t.Target (address.Address) (struct)

	{

		if err := t.Target.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.FaultEpoch (abi.ChainEpoch) (int64)
	{
		maj, extra, err := cbg.CborReadHeader(br)
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.FaultEpoch = abi.ChainEpoch(extraI)
	}
	// t.FaultType (storage_power.ConsensusFaultType) (int64)
	{
		maj, extra, err := cbg.CborReadHeader(br)
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.FaultType = ConsensusFaultType(extraI)
	}
	return nil
}

func (t *OnMinerSurprisePoStFailureParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{129}); err != nil {
		return err
	}

	// t.NumConsecutiveFailures (int64) (int64)
	if t.NumConsecutiveFailures >= 0 {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.NumConsecutiveFailures))); err != nil {
			return err
		}
	} else {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajNegativeInt, uint64(-t.NumConsecutiveFailures)-1)); err != nil {
			return err
		}
	}
	return nil
}

func (t *OnMinerSurprisePoStFailureParams) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.NumConsecutiveFailures (int64) (int64)
	{
		maj, extra, err := cbg.CborReadHeader(br)
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.NumConsecutiveFailures = int64(extraI)
	}
	return nil
}

func (t *OnSectorTemporaryFaultEffectiveEndParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{129}); err != nil {
		return err
	}

	// t.Weight (storage_power.SectorStorageWeightDesc) (struct)
	if err := t.Weight.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *OnSectorTemporaryFaultEffectiveEndParams) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Weight (storage_power.SectorStorageWeightDesc) (struct)

	{

		if err := t.Weight.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	return nil
}

func (t *OnSectorTemporaryFaultEffectiveBeginParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{129}); err != nil {
		return err
	}

	// t.Weight (storage_power.SectorStorageWeightDesc) (struct)
	if err := t.Weight.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *OnSectorTemporaryFaultEffectiveBeginParams) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Weight (storage_power.SectorStorageWeightDesc) (struct)

	{

		if err := t.Weight.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	return nil
}

func (t *CreateMinerReturn) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{130}); err != nil {
		return err
	}

	// t.IDAddress (address.Address) (struct)
	if err := t.IDAddress.MarshalCBOR(w); err != nil {
		return err
	}

	// t.RobustAddress (address.Address) (struct)
	if err := t.RobustAddress.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *CreateMinerReturn) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.IDAddress (address.Address) (struct)

	{

		if err := t.IDAddress.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.RobustAddress (address.Address) (struct)

	{

		if err := t.RobustAddress.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	return nil
}

func (t *MinerConstructorParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{132}); err != nil {
		return err
	}

	// t.OwnerAddr (address.Address) (struct)
	if err := t.OwnerAddr.MarshalCBOR(w); err != nil {
		return err
	}

	// t.WorkerAddr (address.Address) (struct)
	if err := t.WorkerAddr.MarshalCBOR(w); err != nil {
		return err
	}

	// t.SectorSize (abi.SectorSize) (int64)
	if t.SectorSize >= 0 {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.SectorSize))); err != nil {
			return err
		}
	} else {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajNegativeInt, uint64(-t.SectorSize)-1)); err != nil {
			return err
		}
	}

	// t.PeerId (peer.ID) (string)
	if len(t.PeerId) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.PeerId was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len(t.PeerId)))); err != nil {
		return err
	}
	if _, err := w.Write([]byte(t.PeerId)); err != nil {
		return err
	}
	return nil
}

func (t *MinerConstructorParams) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 4 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.OwnerAddr (address.Address) (struct)

	{

		if err := t.OwnerAddr.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.WorkerAddr (address.Address) (struct)

	{

		if err := t.WorkerAddr.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.SectorSize (abi.SectorSize) (int64)
	{
		maj, extra, err := cbg.CborReadHeader(br)
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.SectorSize = abi.SectorSize(extraI)
	}
	// t.PeerId (peer.ID) (string)

	{
		sval, err := cbg.ReadString(br)
		if err != nil {
			return err
		}

		t.PeerId = peer.ID(sval)
	}
	return nil
}

func (t *SectorStorageWeightDesc) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{131}); err != nil {
		return err
	}

	// t.SectorSize (abi.SectorSize) (int64)
	if t.SectorSize >= 0 {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.SectorSize))); err != nil {
			return err
		}
	} else {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajNegativeInt, uint64(-t.SectorSize)-1)); err != nil {
			return err
		}
	}

	// t.Duration (abi.ChainEpoch) (int64)
	if t.Duration >= 0 {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Duration))); err != nil {
			return err
		}
	} else {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajNegativeInt, uint64(-t.Duration)-1)); err != nil {
			return err
		}
	}

	// t.DealWeight (big.Int) (struct)
	if err := t.DealWeight.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *SectorStorageWeightDesc) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.SectorSize (abi.SectorSize) (int64)
	{
		maj, extra, err := cbg.CborReadHeader(br)
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.SectorSize = abi.SectorSize(extraI)
	}
	// t.Duration (abi.ChainEpoch) (int64)
	{
		maj, extra, err := cbg.CborReadHeader(br)
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.Duration = abi.ChainEpoch(extraI)
	}
	// t.DealWeight (big.Int) (struct)

	{

		if err := t.DealWeight.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	return nil
}
