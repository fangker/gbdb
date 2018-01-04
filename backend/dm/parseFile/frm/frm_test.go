package frm

import (
	"testing"
	"../../ctype"
	"../../../utils"
)


func TestFrm(t *testing.T) {
	sTable:= CreateTableStructure("user")
	var fields [2]*field
	fields[0]=CreateField("name",cType.FIELD_TYPE_VARCHAR,3222,cType.CHAR_TYPE_UTF8,"")
	fields[1]=CreateField("age",cType.FIELD_TYPE_INT,3222,cType.CHAR_TYPE_UTF8,"")
	sTable.AddFields(fields[:]...)
	sTable.CreateKey("kk",cType.DDL_KEY_TYPE_INDEX,&keyPart{"age",cType.FIELD_TYPE_VARCHAR,0})
	utils.Info("表结构",sTable)

	NewFrmFile(sTable)
}
