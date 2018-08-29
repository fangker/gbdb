package page

import (
	"github.com/fangker/gbdb/backend/constants/cType"
	"github.com/fangker/gbdb/backend/utils"
)

const (
	FLST_LEN_OFFSET   = 0
	FLST_FIRST_OFFSET = 4
	FLST_LAST_OFFSET  = 10
)
const (
	FLST_PREV_OFFSET = 0
	FLST_NEXT_OFFSET = 6
)
const (
	FLST_LEN_SIZE   = 4
	FLST_FIRST_SIZE = 6
	FLST_LAST_SIZE  = 6
)
const (
	FLST_PREV_SIZE = 6
	FLST_NEXT_SIZE = 6
)

type FistBaseNode struct {
	_offset int
	data    *cType.PageData
}

func (fbn *FistBaseNode) GetLen() uint32 {
	return utils.GetUint32(fbn.getData(FLST_LEN_OFFSET, FLST_LEN_SIZE))
}

func (fbn *FistBaseNode) GetFirst() (uint32, uint16) {
	page := utils.GetUint32(fbn.getData(FLST_FIRST_OFFSET, 4))
	offset := utils.GetUint16(fbn.getData(FLST_FIRST_OFFSET+4, 2))
	return page, offset
}

func (fbn *FistBaseNode) GetLast() (uint32, uint16) {
	pageNo := utils.GetUint32(fbn.getData(FLST_LAST_OFFSET, 4))
	offset := utils.GetUint16(fbn.getData(FLST_LAST_OFFSET+4, 2))
	return pageNo, offset
}

func (fbn *FistBaseNode) SetFirst(pageNo uint32, offset uint16) {
	fbn.setData(FLST_FIRST_OFFSET, 4, utils.PutUint32(pageNo))
	fbn.setData(FLST_FIRST_OFFSET+4, 2, utils.PutUint16(offset))
}

func (fbn *FistBaseNode) SetLast(pageNo uint32, offset uint16) {
	fbn.setData(FLST_LAST_OFFSET, 4, utils.PutUint32(pageNo))
	fbn.setData(FLST_LAST_OFFSET+4, 2, utils.PutUint16(offset))
}

func (fbn *FistBaseNode) SetLen(len uint32) {
	//copy(fbn.reOffset(FLST_LEN_OFFSET, FLST_LEN_SIZE), utils.PutUint32(len))
}

func reOffset(_offset int, start int, end int) (int, int) {
	return _offset + start, _offset + start + end
}

func (fbn *FistBaseNode) setData(start int, size int, b []byte) {
	sta, end := reOffset(fbn._offset, start, size)
	copy(fbn.data[sta:end], b)
}

func (fbn *FistBaseNode) getData(start int, size int) []byte {
	sta, end := reOffset(fbn._offset, start, size)
	return fbn.data[sta:end]
}

type FirstNode struct {
	_offset int
	data    *cType.PageData
}

func (fn *FirstNode) GetFirst() (uint32, uint16) {
	page := utils.GetUint32(fn.reOffset(FLST_PREV_OFFSET, 4))
	offset := utils.GetUint16(fn.reOffset(FLST_PREV_OFFSET+4, 2))
	return page, offset
}

func (fn *FirstNode) GetLast() (uint32, uint16) {
	pageNo := utils.GetUint32(fn.reOffset(FLST_LAST_OFFSET, 4))
	offset := utils.GetUint16(fn.reOffset(FLST_LAST_OFFSET+4, 2))
	return pageNo, offset
}

func (fn *FirstNode) SetFirst(pageNo uint32, offset uint16) {
	copy(fn.reOffset(FLST_PREV_OFFSET, 4), utils.PutUint32(pageNo))
	copy(fn.reOffset(FLST_PREV_OFFSET+4, 2), utils.PutUint32(pageNo))
}

func (fn *FirstNode) SetLast(pageNo uint32, offset uint16) {
	copy(fn.reOffset(FLST_LAST_OFFSET, 4), utils.PutUint32(pageNo))
	copy(fn.data[FLST_LAST_OFFSET+4:FLST_LAST_OFFSET+6], utils.PutUint16(offset))
}

func (fn *FirstNode) reOffset(start int, end int) []byte {
	return fn.data[fn._offset+start : fn._offset+end]
}

func (fn *FistBaseNode) getFreePage() {

}
