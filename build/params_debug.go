// +build debug

package build

func init() {
	InsecurePoStValidation = true

	power.ConsensusMinerMinPower = big.NewInt(2048)
}

var SectorSizes = []abi.SectorSize{2048}

// Seconds
const BlockDelay = 6

const PropagationDelay = 3

// SlashablePowerDelay is the number of epochs after ElectionPeriodStart, after
// which the miner is slashed
//
// Epochs
const SlashablePowerDelay = 20

// Epochs
const InteractivePoRepConfidence = 6
