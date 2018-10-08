package im

import (
	"sync"
	"github.com/fangker/gbdb/backend/cache"
	"github.com/fangker/gbdb/backend/wrapper"
	"github.com/fangker/gbdb/backend/tbm/tfm"
)

var (
	cachePool = cache.CP
)

type BPlusTree struct {
	wp.Wrapper
	bootPage    uint32
	lock        sync.Mutex
	//cacheBuffer *cache.CachePool
	tfm         *tfm.TableFileManage
}

func CreateBPlusTree(tfm *tfm.TableFileManage, rootPage uint32) *BPlusTree {
	// 检测是否存在
	return &BPlusTree{}
}

func LoadTree(tfm *tfm.TableFileManage, rootPage uint32) *BPlusTree {
	//page.NewDictPage(cachePool.GetPage(tfm.CacheWrapper(), rootPage))
	return &BPlusTree{bootPage: rootPage, tfm: tfm}
}

func (bpt BPlusTree) WP() wp.Wrapper {
	return bpt.tfm.CacheWrapper()
}

func Search() {

}
