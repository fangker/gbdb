package systemCache

import ("github.com/fangker/gbdb/backend/cache"
"github.com/fangker/gbdb/backend/tm"
	"github.com/fangker/gbdb/backend/dm/page"
)

type SysTabler interface {
	Tree()
	Wrapper() cache.Wrapper
	Insert()
	Tfm() *tm.TableFileManage
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