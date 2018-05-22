package page

import (
	"github.com/fangker/gbdb/backend/dm/buffPage"
	"github.com/fangker/gbdb/backend/utils"
)

// 数据字典入口页
var (
	DICT_ROW_ID    [8]byte //最近被赋值的row id，递增，用于给未定义主键的表，作为其隐藏的主键键值来构建btree
	DICT_TABLE_ID  [8]byte //当前系统分配的最大事务ID，每创建一个新表，都赋予一个唯一的table id，然后递增
	DICT_INDEX_ID  [8]byte //用于分配索引ID
	DICT_SPACE_ID  [4]byte //用于分配space id
	DICT_TABLE_IDS [4]byte //SYS_TABLE_IDS索引的root page
	DICT_COLUMNS   [4]byte //SYS_COLUMNS系统表的聚集索引root page
	DICT_INDEXES   [4]byte //SYS_INDEXES系统表的聚集索引root page
	DICT_FIELDS    [4]byte //SYS_FIELDS系统表的聚集索引root page
)

const (
	DICT_ROW_ID_OFFSET    = 36
	DICT_TABLE_ID_OFFSET  = DICT_ROW_ID_OFFSET + 8
	DICT_INDEX_ID_OFFSET  = DICT_TABLE_ID_OFFSET + 8
	DICT_SPACE_ID_OFFSET  = DICT_INDEX_ID_OFFSET + 4
	DICT_TABLE_IDS_OFFSET = DICT_SPACE_ID_OFFSET + 4
	DICT_COLUMNS_OFFSET   = DICT_TABLE_IDS_OFFSET + 4
	DICT_INDEXES_OFFSET   = DICT_COLUMNS_OFFSET + 4
	DICT_FIELDS_OFFSET    = DICT_INDEXES_OFFSET + 4
)

type DictPage struct {
	FH *FilHeader
	BF *pcache.BuffPage
}

func NewDictPage(bf *pcache.BuffPage) *DictPage {
	page := &DictPage{
		FH: &FilHeader{data: bf.GetData()},
		BF: bf,
	}
	return page
}

func (this *DictPage) HdrRowID() uint64 {
	return utils.GetUint64(this.BF.GetData()[DICT_ROW_ID_OFFSET : DICT_ROW_ID_OFFSET+8])
}

func (this *DictPage) SetHdrRowID(rowID uint64) {
	copy(this.BF.GetData()[DICT_ROW_ID_OFFSET:DICT_ROW_ID_OFFSET+8], utils.PutUint64(rowID))
}

func (this *DictPage) HdrTableID() uint64 {
	return utils.GetUint64(this.BF.GetData()[DICT_TABLE_ID_OFFSET : DICT_TABLE_ID_OFFSET+8])
}

func (this *DictPage) SetHdrTableID(tableID uint64) {
	copy(this.BF.GetData()[DICT_TABLE_ID_OFFSET:DICT_TABLE_ID_OFFSET+8], utils.PutUint64(tableID))
}

func (this *DictPage) HdrIndexID() uint64 {
	return utils.GetUint64(this.BF.GetData()[DICT_INDEX_ID_OFFSET : DICT_INDEX_ID_OFFSET+8])
}

func (this *DictPage) SetHdrIndexID(indexID uint64) {
	copy(this.BF.GetData()[DICT_INDEX_ID_OFFSET:DICT_INDEX_ID_OFFSET+8], utils.PutUint64(indexID))
}

func (this *DictPage) HdrSpaceID() uint32 {
	return utils.GetUint32(this.BF.GetData()[ DICT_SPACE_ID_OFFSET : DICT_SPACE_ID_OFFSET+4])
}

func (this *DictPage) SetHdrSpaceID(spaceId uint32) {
	copy(this.BF.GetData()[DICT_SPACE_ID_OFFSET:DICT_SPACE_ID_OFFSET+4], utils.PutUint32(spaceId))
}

func (this *DictPage) HdrTables() uint32 {
	return utils.GetUint32(this.BF.GetData()[ DICT_TABLE_IDS_OFFSET : DICT_TABLE_IDS_OFFSET+4])
}

func (this *DictPage) SetHdrTables(tableId uint32) {
	copy(this.BF.GetData()[DICT_TABLE_IDS_OFFSET:DICT_TABLE_IDS_OFFSET+4], utils.PutUint32(tableId))
}

func (this *DictPage) HdrColumns() uint32 {
	return utils.GetUint32(this.BF.GetData()[ DICT_COLUMNS_OFFSET : DICT_COLUMNS_OFFSET+4])
}
func (this *DictPage) SetHdrColumns(column uint32) {
	copy(this.BF.GetData()[DICT_COLUMNS_OFFSET:DICT_COLUMNS_OFFSET+4], utils.PutUint32(column))
}
func (this *DictPage) HdrIndex() uint32 {
	return utils.GetUint32(this.BF.GetData()[ DICT_INDEXES_OFFSET : DICT_INDEXES_OFFSET+4])
}
func (this *DictPage) SetHdrIndex(index uint32) {
	copy(this.BF.GetData()[DICT_INDEXES_OFFSET:DICT_INDEXES_OFFSET+4], utils.PutUint32(index))

}
func (this *DictPage) HdrFields() uint32 {
	return utils.GetUint32(this.BF.GetData()[ DICT_FIELDS_OFFSET : DICT_FIELDS_OFFSET+4])
}

func (this *DictPage) SetHdrFields(field uint32) {
	copy(this.BF.GetData()[DICT_FIELDS_OFFSET:DICT_FIELDS_OFFSET+4], utils.PutUint32(field))

}
