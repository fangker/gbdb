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
}

func NewINodePage(bf *pcache.BuffPage) *INodePage {
	page := &INodePage{
		FH:  &FilHeader{data: bf.GetData()},
		INN: &FirstNode{data: bf.GetData(), _offset: INODEPAGE_INN_OFFSET},
		BF:  bf,
	}
	return page
}

// 创建FreeInode并初始化首页
func (inp *INodePage) SetFreeInode(pageNo uint32,wrapper cache.Wrapper) {
	SetUsedPage(wrapper,pageNo)
	for i := 0; i < 85; i++ {
		offset := INODEPAGE_INN_OFFSET + 12 + 192*i
		if 0 != utils.GetUint32(inp.BF.GetData()[offset:offset+8]) {
			copy(inp.BF.GetData()[offset:offset+8], utils.PutUint32(pageNo))
			return
		}
		continue
	}
	// 超出范围
}

