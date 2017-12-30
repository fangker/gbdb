package frm

import (
	"testing"
	cType "../../ctype"
	log "../../../utils/log"
)


func TestFrm(t *testing.T) {
	sTable:= CreateTableStructure("user")
	var fields [2]*field
	fields[0]=CreateField("name",cType.FIELD_TYPE_VARCHAR,3222,cType.CHAR_TYPE_UTF8,"")
	fields[1]=CreateField("age",cType.FIELD_TYPE_INT,3222,cType.CHAR_TYPE_UTF8,"")
	sTable.addFields(fields[:]...)
	sTable.CreateKey("kk",cType.DDL_KEY_TYPE_INDEX,&keyPart{"age",cType.FIELD_TYPE_VARCHAR,0})
	log.Info("表结构",sTable)
}
