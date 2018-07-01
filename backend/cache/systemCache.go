package cache



type SysTabler interface {
	Tree()
	Wrapper() Interface
}

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

func LoadSysCache(tables, fileds, columns, index  SysTabler) *SystemCache {
	return &SystemCache{
		Sys_columns: columns, Sys_fields: fileds, Sys_index: index, Sys_tables: tables,
	}
}
