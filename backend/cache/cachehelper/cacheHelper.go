package cachehelper

import (
	"github.com/fangker/gbdb/backend/cache/buffPage"
)

var CpHelper ICpHelper;

type ICpHelper interface {
	PosInBlockAlign(b *byte) *pcache.BlockPage
	OffsetInBlockAlign(b *byte) uint64
}

func PosInBlockAlign(b *byte) *pcache.BlockPage {
	return CpHelper.PosInBlockAlign(b)
}
func OffsetInBlockAlign(b *byte) uint64 {
	return CpHelper.OffsetInBlockAlign(b)
}
