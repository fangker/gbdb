package cache

import (
	"container/list"
	"os"
	"sync"

	"github.com/fangker/gbdb/backend/dm/buffPage"
	"github.com/fangker/gbdb/backend/dm/constants/cType"
)

type CachePool struct {
	pagePool    map[uint32]map[uint32]*pcache.BuffPage
	maxCacheNum uint32
	freeList    *list.List
	flushList   *list.List
	readList    *list.List
	mux         *sync.Mutex
}

var CB *CachePool

func NewCacheBuffer(maxCacheNum uint32) *CachePool {
	cb := &CachePool{
		maxCacheNum: maxCacheNum,
		freeList:    list.New(),
		flushList:   list.New(),
		readList:    list.New(),
		pagePool:    make(map[uint32]map[uint32]*pcache.BuffPage),
	}
	cb.init()
	CB = cb
	return cb
}

func (cb *CachePool) GetPage(wrap Wrapper, pageNo uint32) *pcache.BuffPage {
	// 如果缓存中存在使用缓存
	tsID := wrap.TableID
	file := wrap.File
	if pg, exist := cb.pagePool[tsID][pageNo]; exist {
		return pg
	}
	pg := cb.GetFreePage(file)
	pg.File.Seek(int64(pageNo)*cType.PAGE_SIZE, 0)
	var data cType.PageData
	pg.File.Read(data[:])
	pg.SetData(data)
	pg.SetTableId(tsID)
	pg.SetPageNo(pageNo)
	if ts, exist := cb.pagePool[tsID]; exist {
		ts[pageNo] = pg
		return pg
	}
	pn := make(map[uint32]*pcache.BuffPage)
	pn[pageNo] = pg
	cb.pagePool[tsID] = pn
	return pg
}

func (cb *CachePool) init() {
	num := int(cb.maxCacheNum)
	for i := 0; i < num; i++ {
		cb.freeList.PushBack(pcache.NewBuffPage(0,0))
	}
	CB = cb
}

func (cb *CachePool) GetFreePage(file *os.File) *pcache.BuffPage {
	listEle := cb.freeList.Front()
	pg := cb.freeList.Front().Value.(*pcache.BuffPage)
	cb.freeList.Remove(listEle)
	pg.File = file
	return pg
}

//func (cb *CachePool) SetPage(page  *pcache.BuffPage) *pcache.BuffPage {
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
