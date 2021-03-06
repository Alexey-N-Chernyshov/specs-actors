// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package storage_market

import (
	"fmt"
	"io"

	abi "github.com/filecoin-project/specs-actors/actors/abi"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf

func (t *WithdrawBalanceParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{130}); err != nil {
		return err
	}

	// t.Address (address.Address) (struct)
	if err := t.Address.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Amount (big.Int) (struct)
	if err := t.Amount.MarshalCBOR(w); err != nil {
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

	// t.Address (address.Address) (struct)

	{

		if err := t.Address.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.Amount (big.Int) (struct)

	{

		if err := t.Amount.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	return nil
}

func (t *PublishStorageDealsParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{129}); err != nil {
		return err
	}

	// t.Deals ([]storage_market.StorageDeal) (slice)
	if len(t.Deals) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Deals was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajArray, uint64(len(t.Deals)))); err != nil {
		return err
	}
	for _, v := range t.Deals {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}
	return nil
}

func (t *PublishStorageDealsParams) UnmarshalCBOR(r io.Reader) error {
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

	// t.Deals ([]storage_market.StorageDeal) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Deals: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}
	if extra > 0 {
		t.Deals = make([]StorageDeal, extra)
	}
	for i := 0; i < int(extra); i++ {

		var v StorageDeal
		if err := v.UnmarshalCBOR(br); err != nil {
			return err
		}

		t.Deals[i] = v
	}

	return nil
}

func (t *VerifyDealsOnSectorProveCommitParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{130}); err != nil {
		return err
	}

	// t.DealIDs ([]abi.DealID) (slice)
	if len(t.DealIDs) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.DealIDs was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajArray, uint64(len(t.DealIDs)))); err != nil {
		return err
	}
	for _, v := range t.DealIDs {
		if v >= 0 {
			if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(v))); err != nil {
				return err
			}
		} else {
			if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajNegativeInt, uint64(-v)-1)); err != nil {
				return err
			}
		}
	}

	// t.SectorExpiry (abi.ChainEpoch) (int64)
	if t.SectorExpiry >= 0 {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.SectorExpiry))); err != nil {
			return err
		}
	} else {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajNegativeInt, uint64(-t.SectorExpiry)-1)); err != nil {
			return err
		}
	}
	return nil
}

func (t *VerifyDealsOnSectorProveCommitParams) UnmarshalCBOR(r io.Reader) error {
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

	// t.DealIDs ([]abi.DealID) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.DealIDs: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}
	if extra > 0 {
		t.DealIDs = make([]abi.DealID, extra)
	}
	for i := 0; i < int(extra); i++ {
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

			t.DealIDs[i] = abi.DealID(extraI)
		}
	}

	// t.SectorExpiry (abi.ChainEpoch) (int64)
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

		t.SectorExpiry = abi.ChainEpoch(extraI)
	}
	return nil
}

func (t *GetPieceInfosForDealIDsParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{129}); err != nil {
		return err
	}

	// t.DealIDs ([]abi.DealID) (slice)
	if len(t.DealIDs) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.DealIDs was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajArray, uint64(len(t.DealIDs)))); err != nil {
		return err
	}
	for _, v := range t.DealIDs {
		if v >= 0 {
			if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(v))); err != nil {
				return err
			}
		} else {
			if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajNegativeInt, uint64(-v)-1)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (t *GetPieceInfosForDealIDsParams) UnmarshalCBOR(r io.Reader) error {
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

	// t.DealIDs ([]abi.DealID) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.DealIDs: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}
	if extra > 0 {
		t.DealIDs = make([]abi.DealID, extra)
	}
	for i := 0; i < int(extra); i++ {
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

			t.DealIDs[i] = abi.DealID(extraI)
		}
	}

	return nil
}

func (t *OnMinerSectorsTerminateParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{129}); err != nil {
		return err
	}

	// t.DealIDs ([]abi.DealID) (slice)
	if len(t.DealIDs) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.DealIDs was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajArray, uint64(len(t.DealIDs)))); err != nil {
		return err
	}
	for _, v := range t.DealIDs {
		if v >= 0 {
			if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(v))); err != nil {
				return err
			}
		} else {
			if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajNegativeInt, uint64(-v)-1)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (t *OnMinerSectorsTerminateParams) UnmarshalCBOR(r io.Reader) error {
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

	// t.DealIDs ([]abi.DealID) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.DealIDs: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}
	if extra > 0 {
		t.DealIDs = make([]abi.DealID, extra)
	}
	for i := 0; i < int(extra); i++ {
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

			t.DealIDs[i] = abi.DealID(extraI)
		}
	}

	return nil
}

func (t *GetPieceInfosForDealIDsReturn) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{129}); err != nil {
		return err
	}

	// t.Pieces ([]abi.PieceInfo) (slice)
	if len(t.Pieces) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Pieces was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajArray, uint64(len(t.Pieces)))); err != nil {
		return err
	}
	for _, v := range t.Pieces {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}
	return nil
}

func (t *GetPieceInfosForDealIDsReturn) UnmarshalCBOR(r io.Reader) error {
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

	// t.Pieces ([]abi.PieceInfo) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Pieces: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}
	if extra > 0 {
		t.Pieces = make([]abi.PieceInfo, extra)
	}
	for i := 0; i < int(extra); i++ {

		var v abi.PieceInfo
		if err := v.UnmarshalCBOR(br); err != nil {
			return err
		}

		t.Pieces[i] = v
	}

	return nil
}

func (t *StorageDealProposal) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{138}); err != nil {
		return err
	}

	// t.PieceCID (cid.Cid) (struct)

	if err := cbg.WriteCid(w, t.PieceCID); err != nil {
		return xerrors.Errorf("failed to write cid field t.PieceCID: %w", err)
	}

	// t.PieceSize (abi.PieceSize) (struct)
	if err := t.PieceSize.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Client (address.Address) (struct)
	if err := t.Client.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Provider (address.Address) (struct)
	if err := t.Provider.MarshalCBOR(w); err != nil {
		return err
	}

	// t.ClientSignature (crypto.Signature) (struct)
	if err := t.ClientSignature.MarshalCBOR(w); err != nil {
		return err
	}

	// t.StartEpoch (abi.ChainEpoch) (int64)
	if t.StartEpoch >= 0 {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.StartEpoch))); err != nil {
			return err
		}
	} else {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajNegativeInt, uint64(-t.StartEpoch)-1)); err != nil {
			return err
		}
	}

	// t.EndEpoch (abi.ChainEpoch) (int64)
	if t.EndEpoch >= 0 {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.EndEpoch))); err != nil {
			return err
		}
	} else {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajNegativeInt, uint64(-t.EndEpoch)-1)); err != nil {
			return err
		}
	}

	// t.StoragePricePerEpoch (big.Int) (struct)
	if err := t.StoragePricePerEpoch.MarshalCBOR(w); err != nil {
		return err
	}

	// t.ProviderCollateral (big.Int) (struct)
	if err := t.ProviderCollateral.MarshalCBOR(w); err != nil {
		return err
	}

	// t.ClientCollateral (big.Int) (struct)
	if err := t.ClientCollateral.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *StorageDealProposal) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 10 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.PieceCID (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.PieceCID: %w", err)
		}

		t.PieceCID = c

	}
	// t.PieceSize (abi.PieceSize) (struct)

	{

		if err := t.PieceSize.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.Client (address.Address) (struct)

	{

		if err := t.Client.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.Provider (address.Address) (struct)

	{

		if err := t.Provider.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.ClientSignature (crypto.Signature) (struct)

	{

		if err := t.ClientSignature.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.StartEpoch (abi.ChainEpoch) (int64)
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

		t.StartEpoch = abi.ChainEpoch(extraI)
	}
	// t.EndEpoch (abi.ChainEpoch) (int64)
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

		t.EndEpoch = abi.ChainEpoch(extraI)
	}
	// t.StoragePricePerEpoch (big.Int) (struct)

	{

		if err := t.StoragePricePerEpoch.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.ProviderCollateral (big.Int) (struct)

	{

		if err := t.ProviderCollateral.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.ClientCollateral (big.Int) (struct)

	{

		if err := t.ClientCollateral.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	return nil
}

func (t *StorageDeal) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{129}); err != nil {
		return err
	}

	// t.Proposal (storage_market.StorageDealProposal) (struct)
	if err := t.Proposal.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *StorageDeal) UnmarshalCBOR(r io.Reader) error {
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

	// t.Proposal (storage_market.StorageDealProposal) (struct)

	{

		if err := t.Proposal.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	return nil
}

func (t *OnChainDeal) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{133}); err != nil {
		return err
	}

	// t.ID (abi.DealID) (int64)
	if t.ID >= 0 {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.ID))); err != nil {
			return err
		}
	} else {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajNegativeInt, uint64(-t.ID)-1)); err != nil {
			return err
		}
	}

	// t.Deal (storage_market.StorageDeal) (struct)
	if err := t.Deal.MarshalCBOR(w); err != nil {
		return err
	}

	// t.SectorStartEpoch (abi.ChainEpoch) (int64)
	if t.SectorStartEpoch >= 0 {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.SectorStartEpoch))); err != nil {
			return err
		}
	} else {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajNegativeInt, uint64(-t.SectorStartEpoch)-1)); err != nil {
			return err
		}
	}

	// t.LastUpdatedEpoch (abi.ChainEpoch) (int64)
	if t.LastUpdatedEpoch >= 0 {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.LastUpdatedEpoch))); err != nil {
			return err
		}
	} else {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajNegativeInt, uint64(-t.LastUpdatedEpoch)-1)); err != nil {
			return err
		}
	}

	// t.SlashEpoch (abi.ChainEpoch) (int64)
	if t.SlashEpoch >= 0 {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.SlashEpoch))); err != nil {
			return err
		}
	} else {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajNegativeInt, uint64(-t.SlashEpoch)-1)); err != nil {
			return err
		}
	}
	return nil
}

func (t *OnChainDeal) UnmarshalCBOR(r io.Reader) error {
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

	// t.ID (abi.DealID) (int64)
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

		t.ID = abi.DealID(extraI)
	}
	// t.Deal (storage_market.StorageDeal) (struct)

	{

		if err := t.Deal.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.SectorStartEpoch (abi.ChainEpoch) (int64)
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

		t.SectorStartEpoch = abi.ChainEpoch(extraI)
	}
	// t.LastUpdatedEpoch (abi.ChainEpoch) (int64)
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

		t.LastUpdatedEpoch = abi.ChainEpoch(extraI)
	}
	// t.SlashEpoch (abi.ChainEpoch) (int64)
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

		t.SlashEpoch = abi.ChainEpoch(extraI)
	}
	return nil
}
