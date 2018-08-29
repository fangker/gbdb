package pcache

import (
	"sync"
	"github.com/fangker/gbdb/backend/constants/cType"
	"os"
)


type BuffPage struct {
	tableId uint32
	pageNo  uint32
	dirty   bool
	rwLock  sync.RWMutex
	data    cType.PageData
	pType   uint16
	File *os.File
}

func NewBuffPage(tId uint32, pNo uint32) *BuffPage {
	return &BuffPage{tableId: tId, pageNo: pNo, data: cType.PageData{}}
}

func (bp *BuffPage) SetDirty() {
	bp.dirty = true
}

func (bp *BuffPage) Dirty() bool {
	return bp.dirty
}
func (bp *BuffPage) RLock() {
	bp.rwLock.RLock()
}

func (bp *BuffPage) Lock() {
	bp.rwLock.Lock()
}
func (bp *BuffPage) Unlock() {
	bp.rwLock.Unlock()
}

func (bp *BuffPage) GetData() *cType.PageData {
	return &bp.data
}
func (bp *BuffPage) SetData(data cType.PageData) {
	bp.data = data
}
func (bp *BuffPage) getPtype() uint16 {
	return bp.pType
}

func (bp *BuffPage) TableId() uint32 {
	return bp.tableId
}

func (bp *BuffPage) PageNo() uint32 {
	return bp.pageNo
}

func (bp *BuffPage) SetTableId(tId uint32) {
	bp.tableId = tId
}

func (bp *BuffPage) SetPageNo(pNo uint32) {
	bp.pageNo = pNo
}
