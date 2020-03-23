package actors

import (
	"github.com/filecoin-project/go-address"
	"github.com/textileio/lotus-client/chain/types"
)

type OnChainDeal struct {
	PieceRef  []byte // cid bytes // TODO: spec says to use cid.Cid, probably not a good idea
	PieceSize uint64

	Client   address.Address
	Provider address.Address

	ProposalExpiration uint64
	Duration           uint64 // TODO: spec

	StoragePricePerEpoch types.BigInt
	StorageCollateral    types.BigInt
	ActivationEpoch      uint64 // 0 = inactive
}
type StorageParticipantBalance struct {
	Locked    types.BigInt
	Available types.BigInt
}
