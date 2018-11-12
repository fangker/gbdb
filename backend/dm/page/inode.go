package page

// Inode page

import (
	"github.com/fangker/gbdb/backend/cache/buffPage"
	"github.com/fangker/gbdb/backend/utils"
	"github.com/fangker/gbdb/backend/wrapper"
	"github.com/fangker/gbdb/backend/constants/cType"
)

const (
	INODEPAGE_FH_OFFSET  = 0
	INODEPAGE_INN_OFFSET = 34
)

const (
	INODEPAGE_INN_SIZE = 12
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

func NewINodePage(bf *pcache.BuffPage) *INodePage {
	page := &INodePage{
		FH:  &FilHeader{data: bf.GetData()},
		INL: &FirstNode{data: bf.GetData(), _offset: INODEPAGE_INN_OFFSET},
		BF:  bf,
		wp:  bf.Wp(),
	}
	return page
}

//  在当前inode页 设置inode Segment ID
func (inp *INodePage) CreateInodeEntry() {
	page, offset := inp.GetFreeInode()
	segmentID := getSpaceFsp(inp.wp).FSH.segmentID + 1
	inp.InitInodeEntry(segmentID, page, offset)
}

// 获得空闲 inodeEntry
func (inp *INodePage) InitInodeEntry(segmentID uint64, page uint32, offset uint16) {
	var ie InodeEntry
	ie.data = inp.BF.GetData()
	ie.getFreeList().GetFirst()
}

func (inp *INodePage) GetFreeInode() (page uint32, offset uint16) {
	for i := 0; i < 85; i++ {
		offset := INODEPAGE_INN_OFFSET + INODEPAGE_INN_SIZE + 192*i
		if 0 == utils.GetUint32(inp.BF.GetData()[offset:offset+8]) {
			return inp.BF.PageNo(), uint16(offset)
		}
		continue
	}
	// TODO:
	return 9999, 9999
}

// 获得本页空闲 inode Entry
func (inp *INodePage) GetFreeInodeEntryInThisInodePage() (feoa []uint16) {
	for i := 0; i < 85; i++ {
		offset := INODEPAGE_INN_OFFSET + 12 + 192*i
		if 0 != utils.GetUint32(inp.BF.GetData()[offset:offset+8]) {
			feoa = append(feoa, uint16(offset));
		}
		continue
	}
	return feoa
}

//func (inp *INodePage) Init() {
//	// 初始化此页所有entry加入freeInodeList
//	fsp := NewFSPage(cache.CP.GetPage(inp.wp, 0))
//	fsp.FSH.freeInodeList.AddToLast(inp.INL)
//}

//func (inp *INodePage) IsThisInode(page uint32, offset uint16) {
//
//}

type InodeEntry struct {
	data    *cType.PageData
	_offset uint16
}

func (ie InodeEntry) getFreeList() *FistBaseNode {
	return &FistBaseNode{_offset: ie._offset, data: ie.data}
}

func (ie InodeEntry) getNotFullList() *FistBaseNode {
	return &FistBaseNode{_offset: ie._offset + FIRST_BASE_NODE_SIZE, data: ie.data}
}

func (ie InodeEntry) getFullList() *FistBaseNode {
	return &FistBaseNode{_offset: ie._offset + FIRST_BASE_NODE_SIZE*2, data: ie.data}
}
