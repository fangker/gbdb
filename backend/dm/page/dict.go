package page

// 数据字典入口页
var (
	DICT_ROW_ID       [8]byte //最近被赋值的row id，递增，用于给未定义主键的表，作为其隐藏的主键键值来构建btree
	DICT_TABLE_ID     [8]byte //当前系统分配的最大事务ID，每创建一个新表，都赋予一个唯一的table id，然后递增
	DICT_INDEX_ID     [8]byte //用于分配索引ID
	DICT_MAX_SPACE_ID [4]byte //用于分配space id
	DICT_TABLES       [4]byte //SYS_TABLES系统表的聚集索引root page
	DICT_TABLE_IDS    [4]byte //SYS_TABLE_IDS索引的root page
	DICT_COLUMNS      [4]byte //SYS_COLUMNS系统表的聚集索引root page
	DICT_INDEXES      [4]byte //SYS_INDEXES系统表的聚集索引root page
	DICT_FIELDS       [4]byte //SYS_FIELDS系统表的聚集索引root page
)
