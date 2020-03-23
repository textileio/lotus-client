package build

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/prometheus/common/log"
)

func MaybeGenesis() []byte {
	builtinGen, err := rice.FindBox("genesis")
	if err != nil {
		log.Warnf("loading built-in genesis: %s", err)
		return nil
	}
	genBytes, err := builtinGen.Bytes("devnet.car")
	if err != nil {
		log.Warnf("loading built-in genesis: %s", err)
	}

	return genBytes
}
