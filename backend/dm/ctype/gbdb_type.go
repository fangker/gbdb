package cType

// Filed type
const (
	FIELD_TYPE_STRING  = iota
	FIELD_TYPE_VARCHAR
	FIELD_TYPE_INT
	FIELD_TYPE_BIT
)

// Character type
const (
	CHAR_TYPE_UTF8 = iota
)

// Table Key Type
const (
	DDL_KEY_TYPE_PRIMARY = iota
	DDL_KEY_TYPE_UNIQUE
	DDL_KEY_TYPE_INDEX
)
