package page

import (
	"github.com/fangker/gbdb/backend/dm/constants/cType"
	"github.com/fangker/gbdb/backend/utils"
)

const (
	FLST_LEN_OFFSET   = 0
	FLST_FIRST_OFFSET = 4
	FLST_LAST_OFFSET  = 10
)
const (
	FLST_LEN_SIZE   = 4
	FLST_FIRST_SIZE = 6
	FLST_LAST_SIZE  = 6
)

type FistBaseNode struct {
	_offset     int
	data        *cType.PageData
}

func (fbn *FistBaseNode) GetLen() uint32 {
	return utils.GetUint32(fbn.reOffset(FLST_LEN_OFFSET, FLST_LEN_SIZE))
}

func (fbn *FistBaseNode) GetFirst() (uint32, uint16) {
	page := utils.GetUint32(fbn.reOffset(FLST_FIRST_OFFSET,4))
	offset := utils.GetUint16(fbn.reOffset(FLST_FIRST_OFFSET+4,2))
	return page, offset
}

func (fbn *FistBaseNode) GetLast() (uint32, uint16) {
	pageNo := utils.GetUint32(fbn.reOffset(FLST_LAST_OFFSET,4))
	offset := utils.GetUint16(fbn.reOffset(FLST_LAST_OFFSET+4,2))
	return pageNo, offset
}

func (fbn *FistBaseNode) SetFirst(pageNo uint32, offset uint16) {
	copy(fbn.reOffset(FLST_FIRST_OFFSET,4), utils.PutUint32(pageNo))
	copy(fbn.reOffset(FLST_FIRST_OFFSET+4,2), utils.PutUint32(pageNo))
}

func (fbn *FistBaseNode) SetLast(pageNo uint32, offset uint16) {
	copy(fbn.reOffset(FLST_LAST_OFFSET,4), utils.PutUint32(pageNo))
	copy(fbn.data[FLST_LAST_OFFSET+4:FLST_LAST_OFFSET+6], utils.PutUint16(offset))
}

func (fbn *FistBaseNode) reOffset(start int, end int) []byte {
	return fbn.data[fbn._offset+start:fbn._offset+end]
}
