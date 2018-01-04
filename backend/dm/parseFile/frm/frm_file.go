package frm

import (
	"../../../utils"
	//"../../ctype"
)

const (
	patch     = "../../__test__"
	SUFFIX_DB = "gbdb"
)

const (
	frm_type   = 0
	frm_engine = 0
)

func CreateFrmFile(ts *TableStructure) {
	//file,err := os.OpenFile(patch+"."+SUFFIX_DB,os.O_APPEND|os.O_CREATE|os.O_TRUNC,0777)
	//if err != nil {
	//	defer file.Close()
	//	os.Exit(0)
	//}
	//file.WriteString("")
}

func NewFrmFile(ts *TableStructure) {
	//var header [34]byte
	//// frm type
	//header[0:2] = utils.PutUint16(frm_type)
	//// engine type
	//header[2:2] = utils.PutUint16(frm_engine)
	//header[4:2] = utils.PutUint16(6)
	//header[6:2] = utils.PutUint16(6)
	//// key offset
	//header[8:2] = utils.PutUint16(1000)
	//header[16:2] = utils.PutUint16(cType.CHAR_TYPE_UTF8)
	//// 行类型 default dynamic
	//copy(header[0:2],utils.PutUint16(1))
	//var keyPlaceholder [] byte
	var keyLength, fieldLength uint32
	keyLength = 6
	for _, key := range ts.Keys {
		// key
		keyLength += 6
		// keyPart
		keyLength += uint32(len(key.kPart) * 8)
		// keyName and keyComment
		keyLength += uint32(len(key.kPart)*8 + len(key.kPart)*4)
	}
	// 计算表字段字节占用
	fieldLength = uint32(len(ts.Common))+300
	for _, field := range ts.Fields {
		common_len := uint32(len(field.fComment))
		name_len := uint32(len(field.fName))
		fieldLength += common_len
		fieldLength += name_len
		fieldLength += (common_len +name_len)*2

	}
	utils.Trace("key长度", keyLength)

}
