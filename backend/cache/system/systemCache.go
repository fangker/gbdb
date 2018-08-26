package sc

import ("github.com/fangker/gbdb/backend/cache"
	"github.com/fangker/gbdb/backend/dm/page"
	"github.com/fangker/gbdb/backend/tbm"
	. "github.com/fangker/gbdb/backend/dm/constants/cType"
	"NYADB2/backend/parser/statement"
)

type SysTabler interface {
	Tree()
	Wrapper() cache.Wrapper
	Insert(data XID,st *statement.Insert)
	Tfm() *tbm.TableFileManage
	// 用于载入元组
	LoadTuple(create *statement.Create)
}

var SC *SystemCache

type SystemCache struct {
	//sys_tables
	Sys_tables SysTabler
	// sys_fields
	Sys_fields SysTabler
	// sys_columns
	Sys_columns SysTabler
	// sys_indexes
	Sys_index SysTabler
}


func LoadSysCache(tables, fileds, columns, index SysTabler) *SystemCache {
	SC = &SystemCache{
		Sys_columns: columns, Sys_fields: fileds, Sys_index: index, Sys_tables: tables,
	}
	return SC;
}

func (this SystemCache) SetSysTrxIDStore(TrxID uint32) {

}

func (this SystemCache) SysTrxIDStore() page.DictPager {
	return this.Sys_tables.Tfm().SysDir()
}

func (this SystemCache) CreateTable(){

}
func (this SystemCache) LoadSysTuple()  {

}

