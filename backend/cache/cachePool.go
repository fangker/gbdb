package cache

import (
	"container/list"
	"fmt"
	"github.com/fangker/gbdb/backend/cache/buffPage"
	"github.com/fangker/gbdb/backend/cache/cachehelper"
	"github.com/fangker/gbdb/backend/dstr"
	"github.com/fangker/gbdb/backend/file"
	mtr2 "github.com/fangker/gbdb/backend/mtrs/mtr"
	"github.com/fangker/gbdb/backend/utils/uassert"
	"github.com/fangker/gbdb/backend/utils/ulog"
	"sync"
	"unsafe"
)

var CP *Pool
var UnionPageSize uint64 = 16 * 1024

type Pool struct {
	pagePool    map[uint64]map[uint64]*pcache.BlockPage
	maxCacheNum uint64
	freeList    *list.List
	flushList   *list.List
	lock        *sync.RWMutex
	lruList     *dstr.LRUCache
	blockPages  []*pcache.BlockPage
	frameAddr   *byte // 数据页缓存地址
}

func NewCacheBuffer(maxCacheNum uint64) *Pool {
	cb := &Pool{
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
	maddr := make([]byte, maxCacheNum*UnionPageSize)
	cb.frameAddr = (*byte)(unsafe.Pointer(&maddr))
	ulog.Debug("init buffer pool frame addr is ", cb.frameAddr)
	for i := uint64(0); i < maxCacheNum; i++ {
		uptr := unsafe.Pointer(uintptr(unsafe.Pointer(cb.frameAddr)) + uintptr(UnionPageSize*(i)))
		cb.blockPages[i] = pcache.InitBlockPage((* byte)(uptr))
		cb.freeList.PushBack(cb.blockPages[i])
	}
	ulog.Info(fmt.Sprintf("[CacheBuffer] builded ==>  %0.3f mb %d pages ", float32(maxCacheNum*16)/(2<<9), maxCacheNum))
	cachehelper.CpHelper = cb
	CP = cb
	return cb
}

func (cb *Pool) GetPage(spaceId, pageNo uint64, lockType pcache.BpLockType, imtr *mtr2.Mtr) (bp *pcache.BlockPage) {
	cb.lock.Lock()
	defer func() {
		cb.lock.Unlock()
	}()
	// 如果缓存中存在使用缓存
	if cb.poolMapCheck(spaceId, pageNo) {
		bp = cb.pagePool[spaceId][pageNo]
	}
	bp = cb.ReadPageFromFile(spaceId, pageNo)
	var strMemoLockType mtr2.MemoLock
	if pcache.BP_S_LOCK == lockType {
		strMemoLockType = mtr2.MTR_MEMO_PAGE_S_LOCK
		bp.RLock()
	}
	if pcache.BP_X_LOCK == lockType {
		strMemoLockType = mtr2.MTR_MEMO_PAGE_X_LOCK
		bp.Lock()
	}
	var lbp mtr2.ObjLocker
	lbp = bp
	imtr.AddToMemo(strMemoLockType, lbp)
	return
}

// 检查是否存在
func (cb *Pool) poolMapCheck(tpID, pageNo uint64) bool {
	if _, exist := cb.pagePool[tpID]; !exist {
		cb.pagePool[tpID] = make(map[uint64]*pcache.BlockPage)
	}
	if _, exist := cb.pagePool[tpID][pageNo]; exist {
		return true
	}
	return false
}

// add  pageBuffer to bufferPool
func (cb *Pool) ReadPageFromFile(spaceID, pageNo uint64) *pcache.BlockPage {
	bp := (cb.freeList.Remove(cb.freeList.Front())).(*pcache.BlockPage)
	bp.SetSpaceId(spaceID)
	bp.SetPageNo(pageNo)
	t := file.IFileSys.GetSpace(spaceID)
	t.Read(spaceID, pageNo*UnionPageSize, bp.Ptr)
	cb.lruList.Set(uintptr(unsafe.Pointer(&bp)), bp)
	return bp
}

func (cb *Pool) WritePageFromFile(spaceID, pageNo uint64) *pcache.BlockPage {
	bp := (cb.freeList.Remove(cb.freeList.Front())).(*pcache.BlockPage)
	bp.SetSpaceId(spaceID)
	bp.SetSpaceId(pageNo)
	p := unsafe.Pointer(bp.Ptr)
	file.IFileSys.GetSpace(spaceID).Write(spaceID, pageNo*UnionPageSize,(* byte)(p))
	cb.flushList.PushFront(bp)
	return bp
}

func (cb *Pool) PosInBlockAlign(b *byte) *pcache.BlockPage {
	ptr := uintptr(unsafe.Pointer(b))
	cachePoolFrameAddr := uintptr(unsafe.Pointer(cb.frameAddr))
	uassert.True(ptr >= cachePoolFrameAddr && cachePoolFrameAddr+uintptr(UnionPageSize*(cb.maxCacheNum)) >= ptr, "buf not found")
	offset := uint64((ptr - cachePoolFrameAddr) / uintptr(UnionPageSize))
	return cb.blockPages[offset]
}

func (cb *Pool) OffsetInBlockAlign(b *byte) uint64 {
	ptr := uintptr(unsafe.Pointer(b))
	cachePoolFrameAddr := uintptr(unsafe.Pointer(cb.frameAddr))
	uassert.True(ptr >= cachePoolFrameAddr && cachePoolFrameAddr+uintptr(UnionPageSize*(cb.maxCacheNum)) >= ptr, "buf not found")
	offset := uint64((ptr - cachePoolFrameAddr) % uintptr(UnionPageSize))
	return offset
}
