package tbm

import (
	"github.com/fangker/gbdb/backend/im"
	"NYADB2/backend/parser/statement"
	"github.com/fangker/gbdb/backend/tbm/tfm"
	"github.com/fangker/gbdb/backend/utils"
	"github.com/fangker/gbdb/backend/tm"
)

type filedItem []*field

type TableManage struct {
	TableID   uint32
	SpaceID   uint32
	TableName string
	tfm       *tfm.TableFileManage
	tree      *im.BPlusTree
	//index
	//vm
	fields []*field
}

func NewTableManage(tableName string) *TableManage {
	this := &TableManage{TableName: tableName}
	//this.tree = im.CreateBPlusTree(this.TableID, rootPage)
	//tfm.CreateIndex();
	return this
}

func LoadTableManage(tableName string, root uint32) *TableManage {
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

func (this *TableManage) LoadTfm(tfm *tfm.TableFileManage) {
	this.tfm = tfm
}

func (this *TableManage) Insert(trx *tm.Transaction, st *statement.Insert) {

}

func (this *TableManage) Update() {

}

func (this *TableManage) Delete() {

}
func (this *TableManage) Tree() {

}


func (this *TableManage) parseEntity(ist *statement.Insert) []*field {
	var fields []*field;
	for _, f := range this.fields {
		index := utils.IndexOfStringArray(ist.Fields, f.name)
		if index > -1 {
			fields = append(fields, &field{name: f.name, value: ist.Values[index], fType: f.fType, Length: f.Length, Precision: f.Precision})
		} else {
			fields = append(fields, &field{name: f.name, value: nil, fType: f.fType, Length: f.Length, Precision: f.Precision})
		}
	}
	return fields;
}

// 载入元组
func (this *TableManage) LoadTuple(create *statement.Create) {
	this.TableName = create.TableName
	for _, v := range create.Fields {
		f := &field{name: v.Name, fType: v.FType, Length: v.Length, Precision: v.Precision}
		this.fields = append(this.fields, f)
	}
}
