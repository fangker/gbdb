package tbm

import (
	"github.com/fangker/gbdb/backend/im"
	"fmt"
	. "github.com/fangker/gbdb/backend/dm/constants/cType"
	"NYADB2/backend/parser/statement"
)

type TableManage struct {
	TableID   uint32
	TableName string
	tfm       *TableFileManage
	tree      *im.BPlusTree
	field     []*field
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

//func (this *TableManage) Wrapper() cache.Wrapper {
//	return this.tfm.wrapper()
//}

func (this *TableManage) Insert(xid XID, st *statement.Insert) {
	fmt.Println(this.tree);
}

func (this *TableManage) Update() {

}

func (this *TableManage) Delete() {

}
func (this *TableManage) Tree() {

}

// 载入元组
func (this *TableManage) LoadTuple(t *TableManage, create *statement.Create) {
	this.TableName = create.TableName
	for _, v := range create.Fields {
		f := &field{name: v.Name, fType: v.FType, Length: v.Length, Precision: v.Precision}
		this.field = append(this.field, f)
	}
}
