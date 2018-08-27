package cType

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

type LSN uint64
type XID uint64
type INDEX_TYPE int
type FIELD_TYPE int
