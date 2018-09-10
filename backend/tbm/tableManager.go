package tbm

import (
	"github.com/fangker/gbdb/backend/im"
	"github.com/fangker/gbdb/backend/tbm/tfm"
	"github.com/fangker/gbdb/backend/utils"
	"github.com/fangker/gbdb/backend/tm"
	"github.com/fangker/gbdb/backend/cache/system"
	"github.com/fangker/gbdb/backend/tbm/fd"
	"github.com/fangker/gbdb/backend/parser/statement"
	"github.com/fangker/gbdb/backend/utils/log"
)

type TableManage struct {
	TableID   uint32
	SpaceID   uint32
	TableName string
	tfm       *tfm.TableFileManage
	tree      *im.BPlusTree
	//index
	//vm
	fields []*fd.Field
}

func NewTableManage(tableName string) *TableManage {
	this := &TableManage{TableName: tableName}
	//this.tree = im.CreateBPlusTree(this.TableID, rootPage)
	//tfm.CreateIndex();
	return this
}

func LoadTableManage(tableName string, root uint32) *TableManage {
	this := &TableManage{TableName: tableName}
	//this.tree
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
	sc.SC.SysTrxIDStore().HdrRowID()
	t:=this.parseEntity(st)
	log.Trace("assd",log.AnyViewToString(t))
}

func (this *TableManage) Update() {

}

func (this *TableManage) Delete() {

}

func (this *TableManage) Tree() {

}

func (this *TableManage) parseEntity(ist *statement.Insert) []fd.Field {
	var fields []fd.Field;
	for _, f := range this.fields {
		index := utils.IndexOfStringArray(ist.Fields, f.Name)
		if index > -1 {
			fields = append(fields, fd.Field{Name: f.Name, Value: ist.Values[index], FType: f.FType, Length: f.Length, Precision: f.Precision})
		} else {
			fields = append(fields, fd.Field{Name: f.Name, Value: nil, FType: f.FType, Length: f.Length, Precision: f.Precision})
		}
	}
	return fields;
}

// 载入元组
func (this *TableManage) LoadTuple(create *statement.Create) {
	this.TableName = create.TableName
	for _, v := range create.Fields {
		f := &fd.Field{Name: v.Name, FType: v.FType, Length: v.Length, Precision: v.Precision}
		this.fields = append(this.fields, f)
	}
}
