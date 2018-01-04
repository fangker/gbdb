package frm

import (
	"errors"
	"math"
)

type TableStructure struct {
	Name   string
	Keys   []key
	Fields map[string]field
	Char   uint8
	Common string
}


type key struct {
	kName    string
	fieldNum uint32
	kType    uint8
	kPart    [] keyPart
	kCommon  string
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
	fIndex   uint32
	fComment string
	fDefault interface{}
}

// 创建表结构
func CreateTableStructure(tableName string) *TableStructure {
	return &TableStructure{Name: tableName,Fields:make(map[string]field)}
}

// 增加 Fields
func (this *TableStructure) AddFields(fields ...*field) (err error) {
	for _, field := range fields {
			if _, ok := this.Fields[field.fName]; ok {
				return errors.New("[duplicate failed]" + string(field.fName))
			}
		this.Fields[field.fName] = *field
	}
	return nil
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
func (this *TableStructure) CreateKey(name string, kType uint8, keyParts ...*keyPart) (*key,error) {
	var num uint32
	var isDuplicate bool
	tKey := &key{kName: name}
	for _, keyPart := range keyParts {
		tKey.fieldNum += 1
		keyPart.index  =  this.Fields[keyPart.name].fIndex
		tKey.kPart =  append(tKey.kPart,*keyPart)
	}
	tKey.kType = kType
	tKey.fieldNum = num
	for _, key := range this.Keys{
			min :=int(math.Min(float64(len(key.kPart)),float64(len(tKey.kPart))))
			for i:=0;i<min;i++{
				if(key.kPart[i].index==tKey.kPart[i].index){
					if(i==min){
						isDuplicate=true
						break
					}
					continue
				}else{
					break
				}
			}
	}
	if(isDuplicate){
		return nil,errors.New("[duplicate key in]" + tKey.KName)
	}
	this.Keys = append(this.Keys, *tKey)
	return tKey,nil
}
