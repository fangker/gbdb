package parser


type Drop struct {
	TableName string
}

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
	Fields    []string
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