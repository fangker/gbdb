package tbm

import (
	"github.com/fangker/gbdb/backend/cache"
	"github.com/fangker/gbdb/backend/im"
	"fmt"
	. "github.com/fangker/gbdb/backend/dm/constants/cType"
)

type TableManage struct {
	TableID   uint32
	TableName string
	tfm       *TableFileManage
	tree      *im.BPlusTree
}

func NewTableManage(tfm *TableFileManage, tableName string, rootPage uint32) *TableManage {
	this := &TableManage{tfm: tfm, TableName: tableName, TableID: tfm.TableID}
	this.tree = im.CreateBPlusTree(this.TableID, rootPage)
	tfm.CreateIndex();
	return this
}

func (this *TableManage) Tfm() *TableFileManage {
	return this.tfm
}

func (this *TableManage) Wrapper() cache.Wrapper {
	return this.tfm.wrapper()
}

func (this *TableManage) Insert(xid XID){
	fmt.Println(this.tree);
}

func (this *TableManage) Update()  {

}

func (this *TableManage) Delete()  {

}
func (this *TableManage) Tree()  {

}
