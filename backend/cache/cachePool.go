package cache

import (
	"sync"
	"github.com/fangker/gbdb/backend/cache/buffPage"
	"github.com/fangker/gbdb/backend/constants/cType"
	"container/list"
	"strconv"
	"github.com/fangker/gbdb/backend/wrapper"
)

type CachePool struct {
	pagePool    map[uint32]map[uint32]map[uint32]*pcache.BuffPage
	maxCacheNum uint32
	freeList    *list.List
	flushList   *LRUCache
	lruList     *LRUCache
	mux         *sync.Mutex
}

var CP *CachePool

func NewCacheBuffer(maxCacheNum uint32) *CachePool {
	cb := &CachePool{
		maxCacheNum: maxCacheNum,
		freeList:    list.New(),
		flushList:   NewLRUCache(maxCacheNum),
		lruList:     NewLRUCache(maxCacheNum),
		pagePool:    make(map[uint32]map[uint32]map[uint32]*pcache.BuffPage),
	}
	cb.init()
	CP = cb
	return cb
}

func (cb *CachePool) GetPage(wrap wp.Wrapper, pageNo uint32) *pcache.BuffPage {
	// 如果缓存中存在使用缓存
	tbID := wrap.TableID
	tpID := wrap.SpaceID
	if _, exist := cb.pagePool[tpID]; !exist {
		cb.pagePool[tpID] = make(map[uint32]map[uint32]*pcache.BuffPage)
	}
	if _, exist := cb.pagePool[tpID][tbID]; !exist {
		cb.pagePool[tpID][tbID] = make(map[uint32]*pcache.BuffPage)
	}
	if pg, exist := cb.pagePool[tpID][tbID][pageNo]; exist {
		return pg
	}
	pg := cb.GetFreePage(wrap)
	pg.SetPageNo(pageNo)
	// read data
	var data cType.PageData
	pg.File.Seek(int64(pageNo)*cType.PAGE_SIZE, 0)
	pg.File.Read(data[:])
	pg.SetData(data)
	pn := make(map[uint32]*pcache.BuffPage)
	cb.pagePool[tpID][tbID] = pn
	pn[pageNo] = pg
	return pg
}

func (cb *CachePool) init() {
	num := int(cb.maxCacheNum)
	for i := 0; i < num; i++ {
		cb.freeList.PushBack(pcache.NewBuffPage(wp.Wrapper{}))
	}
	CP = cb
}

// 将缓存等待页面移除加入LRU链表返回bufferPage
func (cb *CachePool) GetFreePage(wp wp.Wrapper) *pcache.BuffPage {
	listEle := cb.freeList.Front()
	pg := cb.freeList.Front().Value.(*pcache.BuffPage)
	cb.freeList.Remove(listEle)
	pg.SetWrapper(wp)
	return pg
}

func (cb *CachePool) GetFlushPage(wrap wp.Wrapper, pageNo uint32) *pcache.BuffPage {
	pg := cb.GetPage(wrap, pageNo)
	pg.SetDirty()
	cb.flushList.Set(strconv.Itoa(int(wrap.TableID))+strconv.Itoa(int(pageNo)), pg)
	return pg;
}

func (cb *CachePool) ForceFlush(wrap wp.Wrapper) {
	for l := cb.flushList.List().Front(); l != nil; l = l.Next() {
		val := l.Value.(*CacheNode);
		pg := val.Value.(*pcache.BuffPage);
		if (pg.Dirty() == true) {
			wrap.File.WriteAt(pg.GetData()[:], int64(pg.PageNo()*cType.PAGE_SIZE))
			wrap.File.Sync()
		}
	}
}
