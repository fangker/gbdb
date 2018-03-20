package cacheBuffer

import (
	"github.com/fangker/gbdb/backend/dm/buffPage"
	"container/list"
	"sync"
	"github.com/fangker/gbdb/backend/dm/spaceManage"
)

type CacheBuffer struct {
	spaceManage spaceManage.SpaceManage
	pagePool    map[uint32]map[uint64]*pcache.BuffPage
	maxCacheNum uint32
	freeList    *list.List
	flushList   *list.List
	readList    *list.List
	mux         *sync.Mutex
}

var CB *CacheBuffer

func NewCacheBuffer(maxCacheNum uint32) *CacheBuffer {
	cb:= &CacheBuffer{
		maxCacheNum: maxCacheNum,
		freeList:    list.New(),
		flushList:   list.New(),
		readList:    list.New(),
		pagePool: make(map[uint32]map[uint64]*pcache.BuffPage),
	}
	cb.init()
	return cb
}

func (cb *CacheBuffer) GetPage(tsID uint32, pageNo uint64) *pcache.BuffPage{
	// 如果缓存中存在使用缓存
	if pg,exist:=cb.pagePool[tsID][pageNo];exist{
		return pg
	}
	pg:=cb.GetFreePage()
	pg.Date()[:] =spaceManage.TF(tsID).GetPage(pageNo)[:]
	return pg
}

func (cb *CacheBuffer) init()  {
	num:=int(cb.maxCacheNum)
	for i:=0;i<num;i++ {
		cb.freeList.PushBack(pcache.NewBuffPage())
	}
	CB=cb
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
