package tm

import "github.com/fangker/gbdb/backend/cache"

type TableManager struct {
	TableID uint32
	TableName string
	tfm     *TableFileManage
}

func NewTableManager(tfm *TableFileManage,tableName string) *TableManager {
	return &TableManager{tfm: tfm, TableName:tableName,TableID: tfm.TableID}
}

func (this *TableManager) Tfm() *TableFileManage {
	return this.tfm
}

func (this *TableManager) Wrapper() cache.Wrapper {
	return this.tfm.wrapper()
}
