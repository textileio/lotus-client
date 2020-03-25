package store

import (
	"github.com/filecoin-project/lotus/chain/types"
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
