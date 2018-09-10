package im

import (
	"sync"
	"github.com/fangker/gbdb/backend/cache"
	"github.com/fangker/gbdb/backend/dm/page"
)

var (
	cachePool = cache.CP
)

type BPlusTree struct {
	cache.Wrapper
	bootPage    uint32
	lock        sync.Mutex
	cacheBuffer *cache.CachePool
}

func CreateBPlusTree(wrapper cache.Wrapper, rootPage uint32) *BPlusTree {
	// 检测是否存在
	return &BPlusTree{}
}

func LoadTree(wrapper cache.Wrapper, rootPage uint32) *BPlusTree {
	page.NewDictPage(cachePool.GetPage(wrapper,rootPage))
	return &BPlusTree{}
}

func Search(){

}
