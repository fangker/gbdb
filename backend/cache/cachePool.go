package cache

import (
	"sync"
	"github.com/fangker/gbdb/backend/cache/buffPage"
	"unsafe"
	"container/list"
	"github.com/fangker/gbdb/backend/dstr"
	"fmt"
	"github.com/fangker/gbdb/backend/utils/ulog"
	"github.com/fangker/gbdb/backend/file"
	"github.com/fangker/gbdb/backend/mtr"
	"github.com/fangker/gbdb/backend/cache/cachehelper"
	"github.com/fangker/gbdb/backend/utils/uassert"
)

var CP *CachePool
var UNION_PAGE_SIZE uint64 = 16 * 1024

type CachePool struct {
	pagePool    map[uint64]map[uint64]*pcache.BlockPage
	maxCacheNum uint64
	freeList    *list.List
	flushList   *list.List
	lock        *sync.RWMutex
	lruList     *dstr.LRUCache
	blockPages  []*pcache.BlockPage
	frameAddr   *byte // 数据页缓存地址
}

func NewCacheBuffer(maxCacheNum uint64) *CachePool {
	cb := &CachePool{
		maxCacheNum: maxCacheNum,
		freeList:    list.New(),
		flushList:   list.New(),
		lruList:     dstr.NewLRUCache(maxCacheNum),
		pagePool:    make(map[uint64]map[uint64]*pcache.BlockPage),
		blockPages:  make([]*pcache.BlockPage, maxCacheNum),
		frameAddr:   nil,
		lock:        &sync.RWMutex{},
	}
	cb.lock.Lock()
	defer func() {
		cb.lock.Unlock()
	}()
	maddr := make([]byte, maxCacheNum*UNION_PAGE_SIZE);
	cb.frameAddr = (*byte)(unsafe.Pointer(&maddr))
	ulog.Debug("init buffer pool frame addr is ", cb.frameAddr)
	for i := uint64(0); i < maxCacheNum; i++ {
		uptr := uintptr(unsafe.Pointer(cb.frameAddr)) + uintptr(UNION_PAGE_SIZE*(i))
		cb.blockPages[i] = pcache.InitBlockPage(uptr)
		cb.freeList.PushBack(cb.blockPages[i])
	}
	ulog.Info(fmt.Sprintf("[CacheBuffer] builded ==>  %0.3f mb %d pages ", float32(maxCacheNum*16)/(2<<9), maxCacheNum))
	cachehelper.CpHelper = cb
	CP = cb
	return cb
}

func (cb *CachePool) GetPage(spaceId, pageNo uint64, lockType pcache.BpLockType, imtr *mtr.Mtr) (bp *pcache.BlockPage) {
	cb.lock.Lock()
	defer func() {
		cb.lock.Unlock()
	}()
	// 如果缓存中存在使用缓存
	if (cb.poolMapCheck(spaceId, pageNo)) {
		bp = cb.pagePool[spaceId][pageNo];
	}
	bp = cb.ReadPageFromFile(spaceId, pageNo);
	var strMemoLockType mtr.MtrMemoLock
	if pcache.BP_S_LOCK == lockType {
		strMemoLockType = mtr.MTR_MEMO_PAGE_S_LOCK
		bp.RLock()
	}
	if pcache.BP_X_LOCK == lockType {
		strMemoLockType = mtr.MTR_MEMO_PAGE_X_LOCK
		bp.Lock()
	}
	var lbp mtr.MtrObjLocker
	lbp = bp
	imtr.AddToMemo(strMemoLockType, lbp);
	return
}

// 检查是否存在
func (cb *CachePool) poolMapCheck(tpID, pageNo uint64) bool {
	if _, exist := cb.pagePool[tpID]; !exist {
		cb.pagePool[tpID] = make(map[uint64]*pcache.BlockPage)
	}
	if _, exist := cb.pagePool[tpID][pageNo]; exist {
		return true
	}
	return false;
}

// add  pageBuffer to bufferPool
func (cb *CachePool) ReadPageFromFile(spaceID, pageNo uint64) *pcache.BlockPage {
	bp := (cb.freeList.Remove(cb.freeList.Front())).(*pcache.BlockPage)
	bp.SetSpaceId(spaceID)
	bp.SetPageNo(pageNo)
	t := file.IFileSys.GetSpace(spaceID)
	t.Read(spaceID, pageNo*UNION_PAGE_SIZE, bp.Ptr[:])
	cb.lruList.Set(uintptr(unsafe.Pointer(&bp)), bp);
	return bp;
}

func (cb *CachePool) WritePageFromFile(spaceID, pageNo uint64) *pcache.BlockPage {
	bp := (cb.freeList.Remove(cb.freeList.Front())).(*pcache.BlockPage)
	bp.SetSpaceId(spaceID)
	bp.SetSpaceId(pageNo)
	file.IFileSys.GetSpace(spaceID).Write(spaceID, pageNo*UNION_PAGE_SIZE, bp.Ptr[:])
	cb.flushList.PushFront(bp)
	return bp;
}

func (cb *CachePool) PosInBlockAlign(b *byte) *pcache.BlockPage {
	ptr := uintptr(unsafe.Pointer(b))
	cachePoolFrameAddr := uintptr(unsafe.Pointer(cb.frameAddr))
	uassert.True(ptr >= cachePoolFrameAddr && cachePoolFrameAddr+uintptr(UNION_PAGE_SIZE*(cb.maxCacheNum)) >= ptr, "buf not found")
	offset := uint64((ptr - cachePoolFrameAddr) / uintptr(UNION_PAGE_SIZE))
	return cb.blockPages[offset]
}

func (cb *CachePool) OffsetInBlockAlign(b *byte) uint64 {
	ptr := uintptr(unsafe.Pointer(b))
	cachePoolFrameAddr := uintptr(unsafe.Pointer(cb.frameAddr))
	uassert.True(ptr >= cachePoolFrameAddr && cachePoolFrameAddr+uintptr(UNION_PAGE_SIZE*(cb.maxCacheNum)) >= ptr, "buf not found")
	offset := uint64((ptr - cachePoolFrameAddr) % uintptr(UNION_PAGE_SIZE))
	return offset
}

// 将缓存等待页面移除加入LRU链表返回bufferPage
//func (cb *CachePool) addToUrlList(tpID, pageNo uint64) *pcache.BlockPage {
//	listEle := cb.freeList.Front()
//	pg := cb.freeList.Front().Value.(*pcache.BlockPage)
//	cb.freeList.Remove(listEle)
//	return pg
//}
