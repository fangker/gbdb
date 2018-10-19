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

// parse statement

type CreateStmt struct {
	TableName string
	FieldName []string
	FieldType []string
	Index     []string
}

type UpdateStmt struct {
	TableName string
	FieldName string
	Value     interface{}
	Where     *WhereStmt
}

type DeleteStmt struct {
	TableName string
	Where     *WhereStmt
}

type InsertStmt struct {
	TableName string
	Values    []string
}

type SelectStmt struct {
	TableName string
	Fields    []interface{}
	Where     *WhereStmt
}

type WhereStmt struct {
	SingleExp []*SingleExpStmt
}

type SingleExpStmt struct {
	Field string
	CmpOp string
	Value interface{}
	LogicOp    string
}