package storagemarket

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/specs-actors/actors/abi"
	"github.com/filecoin-project/specs-actors/actors/crypto"
	"github.com/ipfs/go-cid"
)

type StorageDealStatus = uint64

const (
	StorageDealUnknown = StorageDealStatus(iota)
	StorageDealProposalNotFound
	StorageDealProposalRejected
	StorageDealProposalAccepted
	StorageDealStaged
	StorageDealSealing
	StorageDealActive
	StorageDealFailing
	StorageDealNotFound

	// Internal

	StorageDealFundsEnsured   // Deposited funds as neccesary to create a deal, ready to move forward
	StorageDealValidating     // Verifying that deal parameters are good
	StorageDealTransferring   // Moving data
	StorageDealWaitingForData // Manual transfer
	StorageDealVerifyData     // Verify transferred data - generate CAR / piece data
	StorageDealPublishing     // Publishing deal to chain
	StorageDealError          // deal failed with an unexpected error
	StorageDealCompleted      // on provider side, indicates deal is active and info for retrieval is recorded
)

var DealStates = []string{
	"StorageDealUnknown",
	"StorageDealProposalNotFound",
	"StorageDealProposalRejected",
	"StorageDealProposalAccepted",
	"StorageDealStaged",
	"StorageDealSealing",
	"StorageDealActive",
	"StorageDealFailing",
	"StorageDealNotFound",

	"StorageDealFundsEnsured",
	"StorageDealValidating",
	"StorageDealTransferring",
	"StorageDealWaitingForData",
	"StorageDealVerifyData",
	"StorageDealPublishing",
	"StorageDealError",
	"StorageDealCompleted",
}

type SignedStorageAsk struct {
	Ask       *StorageAsk
	Signature *crypto.Signature
}

var SignedStorageAskUndefined = SignedStorageAsk{}

type StorageAsk struct {
	// Price per GiB / Epoch
	Price abi.TokenAmount

	MinPieceSize abi.PaddedPieceSize
	Miner        address.Address
	Timestamp    abi.ChainEpoch
	Expiry       abi.ChainEpoch
	SeqNo        uint64
}

type DataRef struct {
	TransferType string
	Root         cid.Cid

	PieceCid  *cid.Cid // Optional, will be recomputed from the data if not given
	PieceSize abi.UnpaddedPieceSize
}
