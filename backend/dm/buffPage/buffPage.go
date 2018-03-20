package pcache

import (
	"sync"
	"github.com/fangker/gbdb/backend/dm/page"
)

type BuffPager interface {
	NewBuffPage()
	SetType()
}

type BuffPage struct {
	dirty bool
	rwLock *sync.RWMutex
	data   page.PageData
	pType  uint16
	Page   *page.Page
	//IndexPage
	//Inode
	Fsp    *page.FSPage
}

func NewBuffPage() *BuffPage {
	return &BuffPage{data: page.PageData{}}
}

func (bp *BuffPage) SetType(pType uint16) {
	bp.pType = pType
	switch pType {
	case page.PAGE_TYPE_PAGE:
		bp.Page = page.NewPage(&bp.data)
	case page.PAGE_TYPE_FSP:
		bp.Fsp = page.NewFSPage(&bp.data)
	}
}

func (bp *BuffPage) Dirty() {
	bp.dirty=true
}

func (bp *BuffPage) RLock(){
	bp.rwLock.RLock()
}

func (bp *BuffPage) WLock(){
	bp.rwLock.Lock()
}

func (bp *BuffPage)GetPosition() (uint32,uint32){
	switch bp.pType {
	case page.PAGE_TYPE_PAGE:
		return bp.Page.FH.Space,bp.Page.FH.Offset
	case page.PAGE_TYPE_FSP:
		return bp.Page.FH.Space,bp.Page.FH.Offset
	default:
		return 0,0
	}
}

func (bp *BuffPage)Date() *page.PageData{
	return &bp.data
}
