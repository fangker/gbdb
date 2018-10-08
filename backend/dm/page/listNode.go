package page

import (
	"github.com/fangker/gbdb/backend/constants/cType"
	"github.com/fangker/gbdb/backend/utils"
	"github.com/fangker/gbdb/backend/wrapper"

)

const (
	FIRST_BASE_NODE_SIZE = 16
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
	_offset uint16
	data    *cType.PageData
	_wp     wp.Wrapper
	page    uint32
}

func (fbn *FistBaseNode) AddToLast(node *FirstNode) {
	if fbn.GetLen() == 0 {
		node.SetFirst(fbn.NPos())
		node.SetLast(fbn.NPos())
		fbn.SetLast(node.NPos())
		fbn.SetFirst(node.NPos())
	} else {
		iter := fbn.GetPrev()
		iter.SetLast(node.NPos())
		fbn.SetFirst(node.NPos())
		node.SetFirst(iter.NPos())
	}
	fbn.SetLen(fbn.GetLen() + 1)

}
func (fbn *FistBaseNode) NPos() Pos {
	return NPos(fbn.page, fbn._offset)
}

func (fbn *FistBaseNode) GetNext() (*FirstNode) {
	pos := fbn.GetFirst()
	node := &FirstNode{_offset: pos.offset, page: pos.page, data: cachePool.GetPage(fbn._wp, pos.page).GetData(), _wp: fbn._wp}
	return node
}

func (fbn *FistBaseNode) GetPrev() (*FirstNode) {
	pos := fbn.GetLast()
	node := &FirstNode{_offset: pos.offset, data: cachePool.GetPage(fbn._wp, pos.page).GetData(), _wp: fbn._wp}
	return node
}

func (fbn *FistBaseNode) GetLen() uint32 {
	return utils.GetUint32(fbn.getData(FLST_LEN_OFFSET, FLST_LEN_SIZE))
}

func (fbn *FistBaseNode) GetFirst() (p Pos) {
	page := utils.GetUint32(fbn.getData(FLST_FIRST_OFFSET, 4))
	offset := utils.GetUint16(fbn.getData(FLST_FIRST_OFFSET+4, 2))
	return NPos(page, offset)
}

func (fbn *FistBaseNode) SetFirst(p Pos) {
	fbn.setData(FLST_FIRST_OFFSET, 4, utils.PutUint32(p.page))
	fbn.setData(FLST_FIRST_OFFSET+4, 2, utils.PutUint16(p.offset))
}

func (fbn *FistBaseNode) GetLast() (p Pos) {
	pageNo := utils.GetUint32(fbn.getData(FLST_LAST_OFFSET, 4))
	offset := utils.GetUint16(fbn.getData(FLST_LAST_OFFSET+4, 2))
	return NPos(pageNo, offset)
}


func (fbn *FistBaseNode) SetLast(p Pos) {
	fbn.setData(FLST_LAST_OFFSET, 4, utils.PutUint32(p.page))
	fbn.setData(FLST_LAST_OFFSET+4, 2, utils.PutUint16(p.offset))
}

func (fbn *FistBaseNode) SetLen(len uint32) {
	fbn.setData(FLST_LEN_OFFSET, 4, utils.PutUint32(len))
}

func reOffset(_offset uint16, start uint16, end uint16) (uint16, uint16) {
	return _offset + start, _offset + start + end
}

func (fbn *FistBaseNode) setData(start uint16, size uint16, b []byte) {
	sta, end := reOffset(fbn._offset, start, size)
	copy(fbn.data[sta:end], b)
}

func (fbn *FistBaseNode) getData(start uint16, size uint16) []byte {
	sta, end := reOffset(fbn._offset, start, size)
	return fbn.data[sta:end]
}

type FirstNode struct {
	page    uint32
	_offset uint16
	data    *cType.PageData
	_wp     wp.Wrapper
}

func (fn *FirstNode) NPos() Pos {
	return NPos(fn.page, fn._offset)
}

func (fn *FirstNode) GetNext() *FirstNode {
	pos := fn.GetFirst()
	node := &FirstNode{_offset: pos.offset, page: pos.page, data: cachePool.GetPage(fn._wp, pos.page).GetData()}
	return node
}

func (fn *FirstNode) GetPrev() *FirstNode {
	pos := fn.GetLast()
	node := &FirstNode{_offset: pos.offset, data: cachePool.GetPage(fn._wp, pos.page).GetData()}
	return node
}

func (fn *FirstNode) GetFirst() (p Pos) {
	page := utils.GetUint32(fn.reOffset(FLST_PREV_OFFSET, 4))
	offset := utils.GetUint16(fn.reOffset(FLST_PREV_OFFSET+4, 2))
	return NPos(page, offset)
}

func (fn *FirstNode) GetLast() (p Pos) {
	pageNo := utils.GetUint32(fn.reOffset(FLST_LAST_OFFSET, 4))
	offset := utils.GetUint16(fn.reOffset(FLST_LAST_OFFSET+4, 2))
	return NPos(pageNo, offset)
}

func (fn *FirstNode) SetFirst(p Pos) {
	copy(fn.reOffset(FLST_PREV_OFFSET, 4), utils.PutUint32(p.page))
	copy(fn.reOffset(FLST_PREV_OFFSET+4, 2), utils.PutUint16(p.offset))
}

func (fn *FirstNode) SetLast(p Pos) {
	copy(fn.reOffset(FLST_LAST_OFFSET, 4), utils.PutUint32(p.page))
	copy(fn.data[FLST_LAST_OFFSET+4:FLST_LAST_OFFSET+6], utils.PutUint16(p.offset))
}

func (fn *FirstNode) reOffset(start uint16, end uint16) []byte {
	return fn.data[fn._offset+start : fn._offset+end+start]
}
