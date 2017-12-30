package frm

import (
	//ctype "../../ctype"
	"errors"
	"strings"
)

type TableStructure struct {
	Name   string
	Keys   []key
	Fields []field
	Char   uint8
}

type key struct {
	KName    string
	fieldNum uint32
	kType    uint8
	kPart    [] keyPart
}

type keyPart struct {
	name   string
	length uint32
	index  uint32
}

type field struct {
	fName    string
	fType    uint8
	fChar    uint8
	fLength  uint32
	fComment string
	fDefault interface{}
}

// 创建表结构
func CreateTableStructure(tableName string) *TableStructure {
	return &TableStructure{Name: tableName}
}

// 增加 Fields
func (this *TableStructure) addFields(fields ...*field) (err error) {
	for _, field := range fields {
		for _, mum := range this.Fields {
			if (strings.Compare(mum.fName, field.fName) == 0) {
				return errors.New("[duplicate failed]" + string(field.fName))
			}
		}
		this.Fields = append(this.Fields, *field)
	}
	return nil
}

// 添加Key
func (this *TableStructure) addKeys(a key) {
	this.Keys = append(this.Keys, a)
}

// 创建Field
func CreateField(name string, character uint8, length uint32, filedType uint8, comment string) *field {
	return &field{
		fName:    name,
		fChar:    character,
		fLength:  length,
		fType:    filedType,
		fComment: comment,
	}
}

// 创建 Key
func (this *TableStructure) CreateKey(name string, kType uint8, keyParts ...keyPart) *key {
	var num uint32
	tKey := &key{KName: name}
	for _, keyPart := range keyParts {
		tKey.fieldNum += 1
		for index, field := range this.Fields {
			if (strings.Compare(keyPart.name, field.fName) == 0) {
				if keyPart.length==0{
					keyPart.length=this.Fields[index].fLength
				}
				keyPart.index = uint32(index)
				tKey.kPart = append(tKey.kPart,keyPart)
				break
			}
		}
	}
	tKey.kType = kType
	tKey.fieldNum = num
	return tKey
}

// 创建 KeyPart
func CreateKeyPart(name string, length uint32) *keyPart {
	return &keyPart{name, length, 0}
}
