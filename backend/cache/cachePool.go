package cache

import (
	"sync"
	"github.com/fangker/gbdb/backend/cache/buffPage"
	"github.com/fangker/gbdb/backend/wrapper"
	"unsafe"
	"container/list"
	"github.com/fangker/gbdb/backend/dstr"
	"fmt"
	"github.com/fangker/gbdb/backend/utils/ulog"
)

var UNION_PAGE_SIZE uint64 = 16 * 1024

type CachePool struct {
	pagePool    map[uint64]map[uint64]*pcache.BlockPage
	maxCacheNum uint64
	freeList    *list.List
	flushList   *list.List
	lock        *sync.Mutex
	lruList     *dstr.LRUCache
	blockPages  []*pcache.BlockPage
	frameAddr   *byte // 数据页缓存地址
	mux         sync.RWMutex
}

var CP *CachePool

func NewCacheBuffer(maxCacheNum uint64) *CachePool {
	cb := &CachePool{
		maxCacheNum: maxCacheNum,
		freeList:    list.New(),
		flushList:   list.New(),
		lruList:     dstr.NewLRUCache(maxCacheNum),
		lock:        &sync.Mutex{},
		pagePool:    make(map[uint64]map[uint64]*pcache.BlockPage),
		blockPages:  make([]*pcache.BlockPage, maxCacheNum),
		frameAddr:   nil,
	}
	cb.mux.Lock()
	defer cb.mux.Unlock()
	maddr := make([]byte, maxCacheNum*UNION_PAGE_SIZE);
	cb.frameAddr = (*byte)(unsafe.Pointer(&maddr))
	ulog.Debug("init buffer pool frame addr is ", cb.frameAddr)
	for i := uint64(0); i < maxCacheNum; i++ {
		uptr := uintptr(unsafe.Pointer(cb.frameAddr)) + uintptr(UNION_PAGE_SIZE*(i+1))
		cb.blockPages[i] = pcache.NewBlockPage(uptr)
		cb.freeList.PushBack(cb.blockPages[i])
	}
	for i := uint64(0); i < maxCacheNum; i++ {
		cb.lruList.Set(uintptr(unsafe.Pointer(&cb.blockPages[i])), *cb.blockPages[i]);
	}
	CP = cb
	ulog.Info(fmt.Sprintf("[CacheBuffer] builed now %0.3f mb %d pages ", float32(maxCacheNum*16)/(2<<9), maxCacheNum))
	return cb
}

func (cb *CachePool) GetPage(wrap wp.Wrapper) *pcache.BlockPage {
	cb.mux.Lock()
	defer cb.mux.Unlock()
	// 如果缓存中存在使用缓存
	tpID := wrap.SpaceID
	pageNo := wrap.PageNo
	if (cb.poolMapCheck(tpID, pageNo)) {
		return cb.pagePool[tpID][pageNo];
	}
	return cb.GetFreePage(tpID, pageNo);
}

// 检查是否存在
func (cb *CachePool) poolMapCheck(tpID, pageNo uint64) bool {
	cb.mux.RLock()
	defer cb.mux.RUnlock()
	if _, exist := cb.pagePool[tpID]; !exist {
		cb.pagePool[tpID] = make(map[uint64]*pcache.BlockPage)
	}
	if _, exist := cb.pagePool[tpID][pageNo]; exist {
		return true
	}
	return false;
}

// add  pageBuffer to bufferPool
func (cb *CachePool) GetFreePage(tpID, pageNo uint64) *pcache.BlockPage {
	bp := (cb.freeList.Remove(cb.freeList.Front())).(*pcache.BlockPage)
	bp.SetPageNo(pageNo)
	bp.SetSpaceId(tpID)
	return bp;
}

// 将缓存等待页面移除加入LRU链表返回bufferPage
func (cb *CachePool) addToUrlList(tpID, pageNo uint64) *pcache.BlockPage {
	listEle := cb.freeList.Front()
	pg := cb.freeList.Front().Value.(*pcache.BlockPage)
	cb.freeList.Remove(listEle)
	return pg
}

// 查找匹配block
func (cb *CachePool) blockPageAlign(b *byte) *pcache.BlockPage {
	ptr := uintptr(unsafe.Pointer(b)) - uintptr(UNION_PAGE_SIZE)
	for _, v := range cb.blockPages {
		if ptr <= uintptr(unsafe.Pointer(v.Ptr)) {
			return v
		}
	}
	panic("BlockPage Align Not Found")
}
