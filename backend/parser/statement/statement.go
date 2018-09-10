package statement

import (
	. "github.com/fangker/gbdb/backend/constants/cType"
)

type Begin struct {
	IsRepeatableRead bool
}

type Commit struct{}
type Abort struct{}

type Drop struct {
	TableName string
}

type Show struct {
}

type Create struct {
	TableName string
	Fields   []Field
	Index    []string
}

type Update struct {
	TableName string
	FieldName string
	Value     string
	Where     *Where
}

type Delete struct {
	TableName string
	Where     *Where
}

type Insert struct {
	TableName string
	Fields    []string
	Values    []string
}

type Read struct {
	TableName string
	Fields    []string
	Where     *Where
}

type Where struct {
	SingleExp1 *SingleExp
	LogicOp    string
	SingleExp2 *SingleExp
}

type SingleExp struct {
	Field string
	CmpOp string
	Value string
}

type Index struct {
	IndexType INDEX_TYPE
	IndexName string
	Field     string
}

type Field struct {
	Name string
	IndexType INDEX_TYPE
	IndexName string
	FType     FIELD_TYPE
	Length    uint32
	Precision  int
}
