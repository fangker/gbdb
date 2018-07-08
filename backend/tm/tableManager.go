package tm

import (
	"github.com/fangker/gbdb/backend/cache"
	"github.com/fangker/gbdb/backend/im"
	"github.com/fangker/gbdb/backend/mtr"
	"fmt"
)

type TableManager struct {
	TableID   uint32
	TableName string
	tfm       *TableFileManage
	tree      *im.BPlusTree
}

func NewTableManager(tfm *TableFileManage, tableName string, rootPage uint32) *TableManager {
	this := &TableManager{tfm: tfm, TableName: tableName, TableID: tfm.TableID}
	this.tree = im.CreateBPlusTree(this.TableID, rootPage)
	tfm.CreateIndex();
	return this
}

func (this *TableManager) Tfm() *TableFileManage {
	return this.tfm
}

func (this *TableManager) Wrapper() cache.Wrapper {
	return this.tfm.wrapper()
}

func (this *TableManager) Insert(){
	mtr.NewTransaction()
	fmt.Println(this.tree);
}

func (this *TableManager) Update()  {

}

func (this *TableManager) Delete()  {

}
func (this *TableManager) Tree()  {

}
