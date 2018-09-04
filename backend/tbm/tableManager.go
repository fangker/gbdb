package tbm

import (
	"github.com/fangker/gbdb/backend/im"
	"fmt"
	. "github.com/fangker/gbdb/backend/constants/cType"
	"NYADB2/backend/parser/statement"
	"github.com/fangker/gbdb/backend/tbm/tfm"
)

type TableManage struct {
	TableID   uint32
	SpaceID   uint32
	TableName string
	tfm       *tfm.TableFileManage
	tree      *im.BPlusTree
	//index
	//vm
	field []*field
}

func NewTableManage(tableName string) *TableManage {
	this := &TableManage{TableName: tableName}
	//this.tree = im.CreateBPlusTree(this.TableID, rootPage)
	//tfm.CreateIndex();
	return this
}

func LoadTableManage(tableName string,root uint32) *TableManage {
	this := &TableManage{TableName: tableName}
	//this.tree = im.CreateBPlusTree(this.TableID, rootPage)
	//tfm.CreateIndex();
	return this
}

func (this *TableManage) Tfm() *tfm.TableFileManage {
	return this.tfm
}

func (this *TableManage) SetTfm(spaceID, tableID uint32, path string) {
	this.tfm = tfm.NewTableFileManage(spaceID, tableID, path)
}

func (this *TableManage) LoadTfm(tfm *tfm.TableFileManage)  {
	this.tfm = tfm
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
func (this *TableManage) LoadTuple(create *statement.Create) {
	this.TableName = create.TableName
	for _, v := range create.Fields {
		f := &field{name: v.Name, fType: v.FType, Length: v.Length, Precision: v.Precision}
		this.field = append(this.field, f)
	}
}
