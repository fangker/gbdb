package dataItem

import (
	"github.com/fangker/gbdb/backend/constants/cType"
	"github.com/fangker/gbdb/backend/tbm/fd"
)

/*
    data []byte
	vf   可变字段长度
	ns   null标识
	rh   记录头信息
	xid  cType.XID
	rp   cType.RID
*/

type DataItem struct {
	data  []byte
	vf    []byte
	ns    byte
	rh    [5]byte
	xid   cType.XID
	rp    cType.RID
	field []*fd.Field
}

func ParseToRowData([]*fd.Field){

}
