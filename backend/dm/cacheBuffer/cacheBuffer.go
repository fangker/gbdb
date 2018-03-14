package cacheBuffer

import (
	"github.com/fangker/gbdb/backend/dm/buffPage"
	"container/list"
	"sync"
)

type CacheBuffer struct {
	pagePool    map[uint32]map[uint64]*pcache.BuffPage
	maxCacheNum uint32
	freeList    *list.List
	flushList   *list.List
	mux         *sync.Mutex
}

func NewCacheBuffer(maxCacheNum uint32) *CacheBuffer {
	cb:= &CacheBuffer{
		maxCacheNum: maxCacheNum,
		freeList:    list.New(),
		flushList:   list.New(),
		pagePool: make(map[uint32]map[uint64]*pcache.BuffPage),
	}
	cb.init()
	return cb
}

func (cb *CacheBuffer) GetPage(tsID uint32, pageNo uint64, isNew bool) {

}

func (cb *CacheBuffer) init()  {
	num:=int(cb.maxCacheNum)
	for i:=0;i<num;i++ {
		cb.freeList.PushBack(pcache.NewBuffPage())
	}
}

func (cb *CacheBuffer) GetFreePage()  *pcache.BuffPage{
	return cb.freeList.Front().Value.(*pcache.BuffPage)
}


//func (cb *CacheBuffer) SetPage(page  *pcache.BuffPage) *pcache.BuffPage {
//	tsID,pageNo = page.GetGPos()
//	if tbPool, exist := cb.pagePool[tsID]; exist {
//		tbPool[pageNo] = page
//	} else {
//		pn := make(map[uint64]*pcache.PCacher)
//		pn[pageNo]=page
//		cb.pagePool[tsID]=pn
//	}
//	cb.flushList.PushBack(page)
//	return page
//}
