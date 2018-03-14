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
	mLock  *sync.Mutex
	rwLock *sync.RWMutex
	data   page.PageData
	pType  uint16
	Page   page.Page
	//IndexPage
	//Inode
}

func NewBuffPage() *BuffPage {
	return &BuffPage{data: page.PageData{}}
}

func (bp *BuffPage) SetType(pType uint16) {
	switch pType {
	case page.PAGE_TYPE_PAGE:
		bp.Page = *page.NewPage(&bp.data)
	}
}

