package cache

import "github.com/fangker/gbdb/backend/tm"

type SystemCache struct {
	//sys_tables
	Sys_tables *tm.TableManager
	// sys_fields
	Sys_fields *tm.TableManager
	// sys_columns
	Sys_columns *tm.TableManager
	// sys_indexes
	Sys_index *tm.TableManager
}

func LoadSysCache(tables,fileds,columns,index *tm.TableManager)  *SystemCache{
	return &SystemCache{
		Sys_columns:columns,Sys_fields:fileds,Sys_index:index,Sys_tables:tables
}
}