package page

// Inode page

import (
	"github.com/fangker/gbdb/backend/cache/buffPage"
	"github.com/fangker/gbdb/backend/utils"
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
	INN *FirstNode
	BF  *pcache.BuffPage
	wp  cache.Wrapper
}

func NewINodePage(bf *pcache.BuffPage, wrapper cache.Wrapper) *INodePage {
	page := &INodePage{
		FH:  &FilHeader{data: bf.GetData()},
		INN: &FirstNode{data: bf.GetData(), _offset: INODEPAGE_INN_OFFSET},
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
	fsp := NewFSPage(cachePool.GetPage(inp.wp, 0))

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

func (inp *INodePage) getLastUsedPageFromInodeList(il *FistBaseNode)  {
	var tmpPage uint32
	page,_:=il.GetLast()
	cachePool.GetPage(inp.wp, page);
	il.GetLast()

}

func (inp *INodePage) init() {
	// 初始化此页所有entry加入freeInodeList
	fsp := NewFSPage(cachePool.GetPage(inp.wp, 0))
	if (fsp.FSH.freeInodeList.GetLen() == 0) {
		feoa := inp.getFreeInodeEntryInThisInodePage();
		for _, v := range feoa {
			fsp.FSH.freeInodeList.SetFirst(inp.BF.PageNo(), v)
			fsp.FSH.freeInodeList.SetLast(inp.BF.PageNo(), v)
		}
	}

}
