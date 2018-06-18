package im

import (
	"sync"
	"github.com/fangker/gbdb/backend/cache"
)

type BPlusTree struct {
	tableID  uint32
	bootPage uint32
	lock     sync.Mutex
	cacheBuffer *cache.CachePool
}

func CreateBPlusTree(tableID uint32,rootPage uint32)*BPlusTree{
	// 检测是否存在
	return &BPlusTree{}
}




