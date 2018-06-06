package im

import (
	"github.com/fangker/gbdb/backend/dm/buffPage"
	"sync"
	"github.com/fangker/gbdb/backend/cache"
)

type BPlusTree struct {
	tableID  uint32
	bootPage *pcache.BuffPage
	lock     sync.Mutex
	cacheBuffer *cache.CachePool
}

func Create(tableID uint32,bootPage pcache.BuffPage){

}




