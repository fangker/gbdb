package frm

import (
	"../../../utils"
	//"../../ctype"
	//"math"
)

const (
	patch     = "../../__test__"
	SUFFIX_DB = "gbdb"
)

const (
	frm_type   = 0
	frm_engine = 0
)

const (
	frm_chunk = 1024
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
	//var  headerLength uint32=34
	var chunkNum uint32
	 header:= make([]byte,frm_chunk)
	chunkNum++
	// frm type
	copy(header[0:2],utils.PutUint16(frm_type))
	// engine type
	copy(header[2:4],utils.PutUint16(frm_engine))
	copy(header[4:6],utils.PutUint16(6))
	copy(header[6:8],utils.PutUint16(6))
	//// key offset
	//copy(header[8:2],utils.PutUint16(1000))
	//copy(header[16:2],utils.PutUint16(cType.CHAR_TYPE_UTF8))
	//// 行类型 default dynamic
	//copy(header[0:2], utils.PutUint16(1))
	//var keyLength, fieldLength uint32
	//keyLength = 6
	//for _, key := range ts.Keys {
	//	// key
	//	keyLength += 6
	//	// keyPart
	//	keyLength += uint32(len(key.kPart) * 8)
	//	// keyName and keyComment
	//	keyLength += uint32(len(key.kPart)*8 + len(key.kPart)*4)
	//}
	//keyPlace := math.Ceil(float64(keyLength)/frm_chunk) * frm_chunk
	//// 计算表字段字节占用
	//fieldLength = uint32(len(ts.Common))*8 + 300
	//for _, field := range ts.Fields {
	//	common_len := uint32(len(field.fComment))
	//	name_len := uint32(len(field.fName))
	//	fieldLength += common_len
	//	fieldLength += name_len
	//	fieldLength += (common_len + name_len) * 2
	//}
	//fieldPlace := math.Ceil(float64(fieldLength)/frm_chunk) * frm_chunk
	//// 设置key
	//keyNum:= len(ts.Keys)
	//var partNum uint32
	//for _,kmum:= range ts.Keys{
	//	partNum+=uint32(len(kmum.kPart))
	//}
	//keyDataPlace:=make([]byte,uint(keyPlace))
	//copy(keyDataPlace[0:1],utils.PutUint16(keyNum));
	//utils.Info("",fieldPlace,keyNum,keyDataPlace,keyPlace,utils.PutUint16(1))

}

