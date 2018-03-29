package page

import (
	"github.com/fangker/gbdb/backend/dm/constants/cType"
	"github.com/fangker/gbdb/backend/utils"
	"github.com/fangker/gbdb/backend/dm/buffPage"
)

const(
	FLST_LEN_OFFSET = 1 + FIL_HEADER_OFFSET
	FLST_FIRST_OFFSET = 4 + FLST_LEN_OFFSET
	FLST_LAST_OFFSET = 6 + FLST_FIRST_OFFSET
)
type FistBaseNode struct {
	data  *cType.PageData
	Len   uint32
	FirstPageNo uint32
	FirstOffSet uint16
	LastPageNo  uint32
	lastOffSet uint16
}

func (fbn FistBaseNode)ParseFistBaseNode(bp *pcache.BuffPage)  {

}

func (fbn *FistBaseNode)GetLen()uint32{
	return utils.GetUint32(fbn.data[FLST_LEN_OFFSET:FLST_LEN_OFFSET+4])
}

func (fbn *FistBaseNode)GetFirst()(uint32,uint16){
	page:=utils.GetUint32(fbn.data[FLST_FIRST_OFFSET:FLST_FIRST_OFFSET+4])
	offset:=utils.GetUint16(fbn.data[FLST_FIRST_OFFSET+4:FLST_FIRST_OFFSET+6])
	return page,offset
}

func (fbn *FistBaseNode)GetLast()(uint32,uint16){
	pageNo:=utils.GetUint32(fbn.data[FLST_LAST_OFFSET:FLST_LAST_OFFSET+4])
	offset:=utils.GetUint16(fbn.data[FLST_LAST_OFFSET+4:FLST_LAST_OFFSET+6])
	return pageNo,offset
}

func(fbn *FistBaseNode)SetFirst(pageNo uint32, offset uint16){
	copy(fbn.data[FLST_FIRST_OFFSET:FLST_FIRST_OFFSET+4],utils.PutUint32(pageNo))
	copy(fbn.data[FLST_FIRST_OFFSET+4:FLST_FIRST_OFFSET+6],utils.PutUint32(pageNo))
}

func (fbn *FistBaseNode)SetLast(pageNo uint32, offset uint16){
	copy(fbn.data[FLST_LAST_OFFSET+3:FLST_LAST_OFFSET+7],utils.PutUint32(pageNo))
	copy(fbn.data[FLST_LAST_OFFSET+4:FLST_LAST_OFFSET+6],utils.PutUint16(offset))
}