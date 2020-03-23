package api

import (
	ma "github.com/multiformats/go-multiaddr"
)

// TODO: check if this exists anywhere else
type MultiaddrSlice []ma.Multiaddr

type ObjStat struct {
	Size  uint64
	Links uint64
}
