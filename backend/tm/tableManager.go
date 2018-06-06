package tm

import "github.com/fangker/gbdb/backend/cache"

type TableManager struct {
	tfm     *TableFileManage
	TableID uint32
}

func NewTableManager(c *cache.CachePool, fp string, cw cache.Wrapper) *TableManager {
	return &TableManager{tfm: &TableFileManage{c, fp, cw}, TableID: cw.TableID}
}

func (this *TableManager) Tfm() *TableFileManage {
	return this.tfm
}
