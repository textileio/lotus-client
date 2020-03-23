package store

import (
	"github.com/textileio/lotus-client/chain/types"
)

const (
	HCRevert  = "revert"
	HCApply   = "apply"
	HCCurrent = "current"
)

type HeadChange struct {
	Type string
	Val  *types.TipSet
}
