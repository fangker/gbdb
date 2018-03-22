package cacheBuffer

import (
	"container/list"
	"sync"
	"github.com/fangker/gbdb/backend/dm/buffPage"
	"os"
	"github.com/fangker/gbdb/backend/dm/constants/cType"
)

type CacheBuffer struct {
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
	CB=cb
	return cb
}

func (cb *CacheBuffer) GetPage(tsID uint32, pageNo uint64,file *os.File) *pcache.BuffPage{
	// 如果缓存中存在使用缓存
	if pg,exist:=cb.pagePool[tsID][pageNo];exist{
		return pg
	}
	pg:=cb.GetFreePage(file)
	pg.File.Seek(int64(pageNo)*cType.PAGE_SIZE,0)
	var data cType.PageData
	pg.File.Read(data[:])
	pg.SetData(data)
	return pg
}

func (cb *CacheBuffer) init()  {
	num:=int(cb.maxCacheNum)
	for i:=0;i<num;i++ {
		cb.freeList.PushBack(pcache.NewBuffPage())
	}
	CB=cb
}

func (cb *CacheBuffer) GetFreePage(file *os.File)  *pcache.BuffPage{
	listEle:= cb.freeList.Front()
	pg:=cb.freeList.Front().Value.(*pcache.BuffPage)
	cb.freeList.Remove(listEle)
	pg.File= file
	return pg
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
