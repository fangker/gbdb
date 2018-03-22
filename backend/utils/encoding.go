package utils

import (
	"bytes"
	"encoding/binary"
	"github.com/fangker/gbdb/backend/dm/constants/cType"
)


func PutUint16(n uint16) []byte {
	x := uint16(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

func PutUint32(n uint32) []byte {
	x := uint32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

func GetUint16(byte []byte) uint16 {
	var x uint16
	binary.Read(bytes.NewBuffer(byte), binary.BigEndian, &x)
	return x
}

func GetUint64(byte []byte) uint64 {
	var x uint64
	binary.Read(bytes.NewBuffer(byte), binary.BigEndian, &x)
	return x
}
func GetUint32(byte []byte) uint32 {
	var x uint32
	binary.Read(bytes.NewBuffer(byte), binary.BigEndian, &x)
	return x
}

func GetPageDate(byte []byte) cType.PageData {
	var x cType.PageData
	binary.Read(bytes.NewBuffer(byte), binary.BigEndian, &x)
	return x
}