package utils

import (
	"bytes"
	"encoding/binary"
	"sort"
)


func PutUint16(n int) []byte {
	x := uint16(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

func PutUint32(n int) []byte {
	x := uint32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

func getUint16(byte []byte) uint16 {
	var x uint16
	binary.Read(bytes.NewBuffer(byte), binary.BigEndian, &x)
	return x
}