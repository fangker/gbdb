package page

import (
	"github.com/fangker/gbdb/backend/utils"
	"github.com/fangker/gbdb/backend/constants/cType"
	"github.com/fangker/gbdb/backend/wrapper"
)

// fsp file header 104byte
var (
	FSP_SPACE_ID        [4]byte  //表空间ID
	FSP_SIZE            [4]byte  //表空间总page数量
	FSP_FREE_LIMIT      [4]byte  //尚未被初始化最小页面
	FSP_FRAG_N_USED     [4]byte  // 已经被使用页数
	FSP_FREE            [16]byte // 空闲簇链表
	FSP_FREE_FRAG       [16]byte //半满簇链表
	FSP_FULL_FRAG       [16]byte //满簇
	FSP_SEG_ID          [8]byte  // segment Id
	FSP_SEG_INODES_FULL [16]byte // 满Inode节点链表
	FSP_SEG_INODES_FREE [16]byte // 半满Inode链表
)

// xdes Entry 40byte
var (
	XDES_ID        [8]byte
	XDES_FLST_NODE [12]byte // Extent链表双向链表
	XDES_STATE     [4]byte  //状态
	XDES_BITMAP    [16]byte //表示簇的页使用状态
)

const (
	XDES_ID_OFFSET   = 0
	XDES_FLST_OFFSET = XDES_ID_SIZE
)

const (
	XDES_ID_SIZE   = 8
	XDES_FLST_SIZE = 12
)

const (
	XDES_ENTRY_SIZE = 40
)

type xdes struct {
	fn   *FirstNode
	data []byte
}

func parseXdes(d []byte) *xdes {
	xdes := &xdes{}
	xdes.fn = &FirstNode{_offset: 8}
	xdes.data = d
	return xdes
}

func (this *xdes) ID() uint32 {
	return utils.GetUint32(this.data[:8])
}

func (this *xdes) State() uint32 {
	return utils.GetUint32(this.data[8:20])
}

func (this *xdes) BitMap() []byte {
	return this.data[24:]
}

type XdesEntry struct {
	data    *cType.PageData
	_offset uint16
	wp      wp.Wrapper
	pos     Pos
	xdesNode    *FirstNode
}

func parseXDES(wp wp.Wrapper, pos Pos) XdesEntry {
	data := cachePool.GetPage(wp, pos.page).GetData()
	return XdesEntry{
		data:    cachePool.GetPage(wp, pos.page).GetData(),
		_offset: pos.offset,
		wp:      wp,
		pos:     pos,
		xdesNode:    &FirstNode{_wp: wp, _offset: XDES_ID_SIZE + pos.offset, data: data},
	}
}

