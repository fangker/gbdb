package pcache

import (
	"sync"
	"github.com/fangker/gbdb/backend/dm/constants/cType"
	"os"
)

//
//cType BuffPager interface {
//	NewBuffPage()
//	SetType()
//}

type BuffPage struct {
	dirty  bool
	rwLock sync.RWMutex
	data   cType.PageData
	pType  uint16
	//Page   *page.Page
	////IndexPage
	////Inode
	//Fsp    *page.FSPage
	File *os.File
}

func NewBuffPage() *BuffPage {
	return &BuffPage{data: cType.PageData{}}
}

func (bp *BuffPage) Dirty() {
	bp.dirty = true
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
