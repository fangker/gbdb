package page

// Inode page

import (
	"github.com/fangker/gbdb/backend/cache/buffPage"
	"github.com/fangker/gbdb/backend/utils"
	"github.com/fangker/gbdb/backend/wrapper"
	"github.com/fangker/gbdb/backend/cache"
)

const (
	INODEPAGE_FH_OFFSET  = 0
	INODEPAGE_INN_OFFSET = 34
)

var (
	FSEG_ID              [8]byte  // segment ID
	FSEG_INODE_PAGE_NODE [12]byte // 头链表
	FSEG_NOT_FULL_N_USED [8]byte  // 在链表上已经被使用的page数

	FSEG_FREE                  [16]byte
	FSEG_FULL                  [16]byte
	FSEG_INODE_ENTRY_INIT_MARK [4]byte //标识位初始化
	// 32 array 4byte
)

type INodePage struct {
	FH  *FilHeader
	INL *FirstNode
	BF  *pcache.BuffPage
	wp  wp.Wrapper
}

func NewINodePage(bf *pcache.BuffPage, wrapper wp.Wrapper) *INodePage {
	page := &INodePage{
		FH:  &FilHeader{data: bf.GetData()},
		INL: &FirstNode{data: bf.GetData(), _offset: INODEPAGE_INN_OFFSET},
		BF:  bf,
		wp:  wrapper,
	}
	return page
}

// 设置inode Segment ID
func (inp *INodePage) CreatInode(pageNo uint32) {
	SetUsedPage(inp.wp, pageNo)
	//for i := 0; i < 85; i++ {
	//	offset := INODEPAGE_INN_OFFSET + 12 + 192*i
	//	if 0 != utils.GetUint32(inp.BF.GetData()[offset:offset+8]) {
	//		// index segment
	//		copy(inp.BF.GetData()[offset:offset+8], utils.PutUint32(pageNo))
	//		// leaf segment
	//		copy(inp.BF.GetData()[offset:offset+8], utils.PutUint32(pageNo))
	//		return
	//	}
	//	continue
	//}
	inp.getFreeInode(pageNo)
}

// 获得空闲 inodeEntry
func (inp *INodePage) setFreeInode(i int) {
	//fsp := NewFSPage(cachePool.GetPage(inp.wp, 0))

}

func (inp *INodePage) getFreeInode(pageNo uint32) int {
	for i := 0; i < 85; i++ {
		offset := INODEPAGE_INN_OFFSET + 12 + 192*i
		if 0 != utils.GetUint32(inp.BF.GetData()[offset:offset+8]) {
			return i
		}
		continue
	}
	return 1
}

// 获得本页空闲 inode Entry
func (inp *INodePage) getFreeInodeEntryInThisInodePage() (feoa []uint16) {
	for i := 0; i < 85; i++ {
		offset := INODEPAGE_INN_OFFSET + 12 + 192*i
		if 0 != utils.GetUint32(inp.BF.GetData()[offset:offset+8]) {
			feoa = append(feoa, uint16(offset));
		}
		continue
	}

	return feoa
}

func (inp *INodePage) Init() {
	// 初始化此页所有entry加入freeInodeList
	fsp := NewFSPage(cache.CP.GetPage(inp.wp, 0))
	fsp_free_inode_list := fsp.FSH.freeInodeList
	if (fsp_free_inode_list.GetLen() == 0) {
		fsp_free_inode_list.SetFirst(inp.BF.PageNo(), inp.INL._offset)
		fsp_free_inode_list.SetLast(inp.BF.PageNo(), inp.INL._offset)
		fsp_free_inode_list.SetLen(1)
	} else {
		var iter *FirstNode
		var page uint32
		var i uint32
		// 切到最后一位
		for i = 0; i < fsp_free_inode_list.GetLen(); i++ {
			iter, page, _ = fsp_free_inode_list.GetNext()
		}
		iter.SetLast(inp.BF.PageNo(), inp.INL._offset)
		inp.INL.SetFirst(page, inp.INL._offset)
		inp.INL.SetLast(fsp.BP.PageNo(), fsp_free_inode_list._offset)
		fsp_free_inode_list.SetLast(inp.BF.PageNo(), fsp_free_inode_list._offset)
		fsp_free_inode_list.SetLen(fsp_free_inode_list.GetLen()+1)
	}
}
