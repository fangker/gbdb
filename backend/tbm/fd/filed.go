package fd

import (
	. "github.com/fangker/gbdb/backend/constants/cType"
)

type Field struct {
	Name  string
	Value interface{}
	FType FIELD_TYPE
	Length    uint32
	Precision  int

}

//func CreateField(name string, value interface{}, fType FIELD_TYPE) *field {
//	return &field{name, value, fType}
//}
