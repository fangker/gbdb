package frm

import (
	"testing"
	cType "../../ctype"
	log "../../../utils/log"
)


func TestFrm(t *testing.T) {
	sTable:= CreateTableStructure("user")
	var fields [2]*field
	fields[0]=CreateField("name",cType.CHAR_TYPE_UTF8,3222,cType.CHAR_TYPE_UTF8)
	fields[1]=CreateField("age",cType.CHAR_TYPE_UTF8,3222,cType.CHAR_TYPE_UTF8)
	sTable.addFields(fields[:]...)
	sTable.CreateKey("kk",cType.DDL_KEY_TYPE_INDEX,"name","age")
	log.Info("表结构",sTable)
}
