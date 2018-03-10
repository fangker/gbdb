package cacheBuffer

import (
	"github.com/fangker/gbdb/backend/dm/pcache"
	"container/list"
	"sync"
)

type CacheBuffer struct {
	pagePool    map[uint32]map[uint64]*pcache.PCacher
	maxCacheNum uint32
	freeList    *list.List
	flushList   *list.List
	mux         *sync.Mutex
}

func NewCacheBuffer(maxCacheNum uint32) *CacheBuffer {
	return &CacheBuffer{
		maxCacheNum: maxCacheNum,
		freeList:    list.New(),
		flushList:   list.New(),
		pagePool: make(map[uint32]map[uint64]*pcache.PCacher),
	}
}

func (cb *CacheBuffer) GetPage(tsID uint32, pageNo uint64, isNew bool) {

}

func (cb *CacheBuffer) SetPage(page  *pcache.PCacher) *pcache.PCacher {
	tsID,pageNo = page.GetGPos()
	if tbPool, exist := cb.pagePool[tsID]; exist {
		tbPool[pageNo] = page
	} else {
		pn := make(map[uint64]*pcache.PCacher)
		pn[pageNo]=page
		cb.pagePool[tsID]=pn
	}
	cb.flushList.PushBack(page)
	return page
}
