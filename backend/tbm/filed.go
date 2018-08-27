package tbm

import (
	. "github.com/fangker/gbdb/backend/dm/constants/cType"
)

type field struct {
	name  string
	value interface{}
	fType FIELD_TYPE
	Length    uint32
	Precision  int

}

//func CreateField(name string, value interface{}, fType FIELD_TYPE) *field {
//	return &field{name, value, fType}
//}
