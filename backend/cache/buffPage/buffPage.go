package pcache

import (
	"sync"
	"github.com/fangker/gbdb/backend/def/cType"
	"github.com/fangker/gbdb/backend/wrapper"
	"unsafe"
)

type BlockPage struct {
	pageNo uint64
	spaceId uint64
	dirty  bool
	rwLock sync.RWMutex
	Ptr    *cType.PageData
	pType  uint16
	loaded bool
	wp     wp.Wrapper
}

func NewBlockPage(uptr uintptr) *BlockPage {
	return &BlockPage{Ptr: (* cType.PageData)(unsafe.Pointer(uptr))}
}

func (bp BlockPage) Wp() wp.Wrapper {
	return bp.wp
}

func (bp *BlockPage) SetDirty() {
	bp.dirty = true
}

func (bp *BlockPage) Dirty() bool {
	return bp.dirty
}
func (bp *BlockPage) RLock() {
	bp.rwLock.RLock()
}

func (bp *BlockPage) Lock() {
	bp.rwLock.Lock()
}
func (bp *BlockPage) Unlock() {
	bp.rwLock.Unlock()
}

func (bp *BlockPage) GetData() *cType.PageData {
	return bp.Ptr
}

func (bp *BlockPage) SetData(data cType.PageData) {
	*bp.Ptr = data
}

func (bp *BlockPage) getPtype() uint16 {
	return bp.pType
}

func (bp *BlockPage) SetWrapper(wp wp.Wrapper) {
	bp.wp = wp;
}


func (bp *BlockPage) PageNo() uint64 {
	return bp.pageNo
}


func (bp *BlockPage) SetSpaceId(spaceId uint64) {
	bp.spaceId = spaceId
}

func (bp *BlockPage) SetPageNo(pNo uint64) {
	bp.pageNo = pNo
}
