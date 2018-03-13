package pcache

import (
	"sync"
	"github.com/fangker/gbdb/backend/dm/page"
	"github.com/fangker/gbdb/backend/utils"
)

type BuffPager interface {
}

type BuffPage struct {
	mux *sync.Mutex
	//page
	pType    uint16
	data     [16384]byte
	lastLSN  uint64
	space    uint32
	offset   uint32
	prev     uint32
	next     uint32
	flushLSN uint32
}

func NewBuffPage() *BuffPage {
	return &BuffPage{data: [16384]byte{}}
}

func (bp *BuffPage) parseHeader() {
 bp.pType = bp.data[page.FIL_PAGE_TYPE:]
}
