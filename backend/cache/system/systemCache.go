package sc

import (
	"github.com/fangker/gbdb/backend/dm/page"
	. "github.com/fangker/gbdb/backend/constants/cType"
	"NYADB2/backend/parser/statement"
	"github.com/fangker/gbdb/backend/tbm"
)

type SysTabler interface {
	Tree()
	Insert(data XID, st *statement.Insert)
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

func (this SystemCache) CreateTable() {

}

func (this SystemCache) LoadSysTuple() {
	// 载入系统元组
	this.Sys_tables.LoadTuple(&statement.Create{TableName: "SYS_TABLES",
		Fields: []statement.Field{
			{Name: "name", FType: FIELD_TYPE_VARCHAR, Length: 64, Precision: 0},
			{Name: "id", FType: FIELD_TYPE_UINT, Length: 8, Precision: 0},
			{Name: "n_cols", FType: FIELD_TYPE_UINT, Length: 1, Precision: 0},
			{Name: "type", FType: FIELD_TYPE_UINT, Length: 1, Precision: 0},
			{Name: "SPACE", FType: FIELD_TYPE_VARCHAR, Length: 4, Precision: 0},
		}});
	this.Sys_fields.LoadTuple(&statement.Create{TableName: "SYS_FIELDS",
		Fields: []statement.Field{
			{Name: "index_id", FType: FIELD_TYPE_UINT, Length: 64, Precision: 0},
			{Name: "pos", FType: FIELD_TYPE_UINT, Length: 8, Precision: 0},
			{Name: "col_name", FType: FIELD_TYPE_UINT, Length: 64, Precision: 0},
		}})
	this.Sys_columns.LoadTuple(&statement.Create{TableName: "SYS_COLUMNS",
		Fields: []statement.Field{
			{Name: "table_id", FType: FIELD_TYPE_VARCHAR, Length: 8, Precision: 0},
			{Name: "pos", FType: FIELD_TYPE_UINT, Length: 1, Precision: 0},
			{Name: "name", FType: FIELD_TYPE_UINT, Length: 64, Precision: 0},
			{Name: "mtype", FType: FIELD_TYPE_UINT, Length: 1, Precision: 0},
			{Name: "prtype", FType: FIELD_TYPE_UINT, Length: 1, Precision: 0},
			{Name: "len", FType: FIELD_TYPE_UINT, Length: 8, Precision: 0},
		}})
	this.Sys_index.LoadTuple(&statement.Create{TableName: "SYS_INDEXES",
		Fields: []statement.Field{
			{Name: "table_id", FType: FIELD_TYPE_VARCHAR, Length: 8, Precision: 0},
			{Name: "id", FType: FIELD_TYPE_UINT, Length: 8, Precision: 0},
			{Name: "name", FType: FIELD_TYPE_UINT, Length: 64, Precision: 0},
			{Name: "n_fields", FType: FIELD_TYPE_UINT, Length: 1, Precision: 0},
			{Name: "type", FType: FIELD_TYPE_UINT, Length: 1, Precision: 0},
			{Name: "len", FType: FIELD_TYPE_UINT, Length: 8, Precision: 0},
		}})
}

//SYS_TABLES
//用来存储所有InnoDB为存储引擎的表
//NAME：表示一个表的表名
//ID：表示这个表的ID号
//N_COLS：表示这个表的列的个数，建表指定的列数。
//TYPE：表示这个表的存储类型，包括记录的格式，压缩等信息。
//SPACE：表示这个表所在表空间ID号。这个表对应的主键列为NAME，同时还有一个在ID号上的唯一索引。
//
//SYS_COLUMNS
//用来存储InnoDB中定义的所有表中所有列的信息，每一列对应这个表的一条记录。
//TABLE_ID：表示这个列所属的表的ID号
//POS：表示这个列在表中是第几列。
//NAME：表示这个列名。
//MTYPE：表示这个列的主数据类型。
//PRTYPE：表示这个列的一些精确数据类型，它是一个组合值，包括NULL标志，是否有符号数的标志，是否是二进制字符串的标志及表示这个列是真的varchar
//LEN：表示这个列数据的精度，目前没有用到。
//
//SYS_INDEXES
//用来存储InnoDB中所有表的索引信息，每一条记录对应一个索引
//TABLE_ID：表示这个索引所属的表的ID号。
//ID：表示这个索引的索引ID号
//NAME：表示这个索引的索引名
//N_FIELDS：表示这个索引包括的列个数。
//TYPE：表示这个索引的类型，包括聚簇索引，唯一索引，等
//SPACE：表示这个索引数据所在的表空间ID号
//PAGE_NO：表示这个索引对应的B+树根页面。
//
//
//SYS_FIEDS
//用来存储所有索引中定义的索引列，每一条记录对应一个索引列。
//INDEX_ID：这个列所在的索引
//POS：这个列在某个索引中是第几个索引列
//COL_NAME：这个索引列的列名。
