package constant

// Filed cType
const (
	FIELD_TYPE_UNKNOWN = iota
	FIELD_TYPE_VARCHAR
	FIELD_TYPE_UINT
	FIELD_TYPE_INT
	FIELD_TYPE_BIT
)

// Character cType
const (
	CHAR_TYPE_UTF8 = iota
)

// Index Key Type
const (
	DDL_KEY_TYPE_PRIMARY = iota
	DDL_KEY_TYPE_UNIQUE
	DDL_KEY_TYPE_INDEX
)

// SIZE
const (
	UNIT_PAGE_SIZE uint64 = 16 * 1024
	LOG_BLOCK_SIZE uint64 = 512
)
