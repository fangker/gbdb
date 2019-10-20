package pcache

import (
	"sync"
	"github.com/fangker/gbdb/backend/def/cType"
	"github.com/fangker/gbdb/backend/wrapper"
	"unsafe"
)
type BpLockType uint

const (
	BP_S_LOCK BpLockType = 0
	BP_X_LOCK BpLockType = 1
)

type BlockPage struct {
	pageNo  uint64
	spaceId uint64
	dirty   bool
	rwLock  sync.RWMutex
	Ptr     *cType.PageData
	pType   uint16
	loaded  bool
}

// Init block memory
func InitBlockPage(uptr uintptr) *BlockPage {
	return &BlockPage{Ptr: (* cType.PageData)(unsafe.Pointer(uptr))}
}

func getBlockPage(spaceId, page uint32, uptr uintptr) *BlockPage {
	return &BlockPage{Ptr: (* cType.PageData)(unsafe.Pointer(uptr))}
}

func (bp *BlockPage) GetPos() (spaceId, pageNum uint64) {
	return bp.spaceId, bp.pageNo
}

func (bp BlockPage) Wp() wp.Wrapper {
	return bp.Wp()
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

func (bp *BlockPage) PageNo() uint64 {
	return bp.pageNo
}

func (bp *BlockPage) SetSpaceId(spaceId uint64) {
	bp.spaceId = spaceId
}

func (bp *BlockPage) SetPageNo(pNo uint64) {
	bp.pageNo = pNo
}
