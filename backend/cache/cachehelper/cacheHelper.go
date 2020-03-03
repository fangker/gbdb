package cachehelper

import (
	"github.com/fangker/gbdb/backend/cache/buffPage"
)

var CpHelper ICpHelper;

type ICpHelper interface {
	BlockPageAlign(b *byte) *pcache.BlockPage
	BlockOffsetAlign(b *byte) uint64
}

func BlockPageAlign(b *byte) *pcache.BlockPage {
	return CpHelper.BlockPageAlign(b)
}
func BlockOffsetAlign(b *byte) uint64 {
	return CpHelper.BlockOffsetAlign(b)
}
