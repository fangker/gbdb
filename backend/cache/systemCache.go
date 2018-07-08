package cache



type SysTabler interface {
	Tree()
	Wrapper()  Wrapper
	Insert()
	Tfm() TableFileManager

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

type TableFileManager interface {
	sysDir() DictPager
}

type DictPager interface {

}


func LoadSysCache(tables, fileds, columns, index  SysTabler) *SystemCache {
	SC = &SystemCache{
		Sys_columns: columns, Sys_fields: fileds, Sys_index: index, Sys_tables: tables,
	}
	return SC;
}

func (this SystemCache)SysTrxIDStore (TrxID uint32) *SystemCache {
	this.Sys_tables.Tfm().sysDir()
}
