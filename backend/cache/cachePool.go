package cache

import (
	"os"
	"sync"

	"github.com/fangker/gbdb/backend/cache/buffPage"
	"github.com/fangker/gbdb/backend/constants/cType"
	"container/list"
	"strconv"
)
type CachePool struct {
	pagePool    map[uint32]map[uint32]*pcache.BuffPage
	maxCacheNum uint32
	freeList    *list.List
	flushList   *LRUCache
	lruList     *LRUCache
	mux         *sync.Mutex
}

var CB *CachePool

func NewCacheBuffer(maxCacheNum uint32) *CachePool {
	cb := &CachePool{
		maxCacheNum: maxCacheNum,
		freeList:    list.New(),
		flushList:   NewLRUCache(maxCacheNum),
		lruList:     NewLRUCache(maxCacheNum),
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
		cb.freeList.PushBack(pcache.NewBuffPage(0, 0))
	}
	CB = cb
}
// 将缓存等待页面移除加入LRU链表返回bufferPage
func (cb *CachePool) GetFreePage(file *os.File) *pcache.BuffPage {
	listEle := cb.freeList.Front()
	pg := cb.freeList.Front().Value.(*pcache.BuffPage)
	cb.freeList.Remove(listEle)
	pg.File = file
	return pg
}

func (cb *CachePool) GetFlushPage(wrap Wrapper, pageNo uint32) *pcache.BuffPage {
	pg := cb.GetPage(wrap, pageNo)
	pg.SetDirty()
	cb.flushList.Set(strconv.Itoa(int(wrap.TableID))+strconv.Itoa(int(pageNo)), pg)
	return pg;
}

func (cb *CachePool) ForceFlush(wrap Wrapper) {
	for l := cb.flushList.List().Front(); l != nil; l = l.Next() {
		val:=l.Value.(*CacheNode);
		pg:=val.Value.(*pcache.BuffPage);
		if(pg.Dirty()==true){
			wrap.File.WriteAt(pg.GetData()[:],int64(pg.PageNo()*cType.PAGE_SIZE))
			wrap.File.Sync()
		}
	}
}
