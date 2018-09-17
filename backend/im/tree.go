package im

import (
	"sync"
	"github.com/fangker/gbdb/backend/cache"
	"github.com/fangker/gbdb/backend/dm/page"
	"github.com/fangker/gbdb/backend/wrapper"
)

var (
	cachePool = cache.CP
)

type BPlusTree struct {
	wp.Wrapper
	bootPage    uint32
	lock        sync.Mutex
	cacheBuffer *cache.CachePool
}

func CreateBPlusTree(wrapper wp.Wrapper, rootPage uint32) *BPlusTree {
	// 检测是否存在
	return &BPlusTree{}
}

func LoadTree(wrapper wp.Wrapper, rootPage uint32) *BPlusTree {
	page.NewDictPage(cachePool.GetPage(wrapper,rootPage))
	return &BPlusTree{}
}

func Search(){

}
