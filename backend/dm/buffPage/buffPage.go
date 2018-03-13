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
	data := bp.data
	bp.pType = utils.GetUint16(data[page.FIL_PAGE_TYPE:page.FIL_PAGE_TYPE+page.FIL_PAGE_TYPE_SIZE])
	bp.lastLSN = utils.GetUint64(data[page.FIL_PAGE_LSN:page.FIL_PAGE_LSN+page.FIL_PAGE_LSN_SIZE])
	bp.space = utils.GetUint32(data[page.FIL_PAGE_SPACE:page.FIL_PAGE_SPACE+page.FIL_PAGE_SPACE_SIZE])
}
