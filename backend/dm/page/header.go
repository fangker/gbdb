package page

import (
	"github.com/fangker/gbdb/backend/utils"
	"github.com/fangker/gbdb/backend/cache/buffPage"
	"github.com/fangker/gbdb/backend/def/cType"
	"github.com/fangker/gbdb/backend/utils/ulog"
)

const (
	FIL_HEADER_END = 34
)

// file header 34bytes
const (
	FIL_PAGE_SPACE_OFFSET          = 0                                                   //   page所属的表空间的space id
	FIL_PAGE_OFFSET_OFFSET         = FIL_PAGE_SPACE_OFFSET + FIL_PAGE_OFFSET_SIZE        //   page no，一般是在表空间的物理偏移量
	FIL_PAGE_PREV_OFFSET           = FIL_PAGE_OFFSET_OFFSET + FIL_PAGE_PREV_SIZE         //    前一页的page no (B+tree的叶子节点是通过链表串起来的，有前后关系
	FIL_PAGE_NEXT_OFFSET           = FIL_PAGE_PREV_OFFSET + FIL_PAGE_NEXT_SIZE           //      后一页的page no
	FIL_PAGE_LSN_OFFSET            = FIL_PAGE_NEXT_OFFSET + FIL_PAGE_LSN_SIZE            //   最后被修改的LSN日志号
	FIL_PAGE_FILE_FLUSH_LSN_OFFSET = FIL_PAGE_LSN_OFFSET + FIL_PAGE_FILE_FLUSH_LSN_SIZE  //   该表空间最后一次被更新的LSN号
	FIL_PAGE_TYPE_OFFSET           = FIL_PAGE_FILE_FLUSH_LSN_OFFSET + FIL_PAGE_TYPE_SIZE //        page的类型
)

const (
	FIL_PAGE_SPACE_SIZE          = 4 //   page所属的表空间的space id
	FIL_PAGE_OFFSET_SIZE         = 4 //   page no，一般是在表空间的物理偏移量
	FIL_PAGE_PREV_SIZE           = 4 //   前一页的page no (B+tree的叶子节点是通过链表串起来的，有前后关系
	FIL_PAGE_NEXT_SIZE           = 4 //   后一页的page no
	FIL_PAGE_LSN_SIZE            = 8 //   最后被修改的LSN日志号
	FIL_PAGE_FILE_FLUSH_LSN_SIZE = 8 //   该表空间最后一次被更新的LSN号
	FIL_PAGE_TYPE_SIZE           = 2 //        page的类型
)

const (
	FSPH_PAGE_PACE = 0
)

const (
	FS_PAGE_SPACE_OFFSET      = 0
	FS_PAGE_MAX_PAGE_OFFSET   = FS_PAGE_MAX_PAGE_SIZE + FS_PAGE_SPACE_OFFSET        // 当前space最大可容纳的page数,文件扩大时才会改变这个值
	FS_PAGE_LIMIT_OFFSET      = FS_PAGE_LIMIT_SIZE + FS_PAGE_MAX_PAGE_OFFSET        // 当前space已经分配初始化的page数,包括空闲的和已经使用的
	FS_PAGE_FRAG_USED_OFFSET  = FS_PAGE_FRAG_USED_SIZE + FS_PAGE_LIMIT_OFFSET       // FSP_FREE_FRAG列表中已经被使用的page数
	FS_FREE_LIST_OFFSET       = FS_FREE_LIST_SIZE + FS_PAGE_FRAG_USED_OFFSET        // space中可用的extent对象列表，extent里面没有一个page被使用
	FS_FRAG_FREE_LIST_OFFSET  = FS_FRAG_FREE_LIST_SIZE + FS_FREE_LIST_OFFSET        // 有可用碎叶page的extent列表，exntent里面有部分page被使用
	FS_FRAG_FULL_LIST_OFFSET  = FS_FRAG_FULL_LIST_SIZE + FS_FRAG_FREE_LIST_OFFSET   // 没有有可用page的extent列表，exntent里面全部page被使用
	FS_SEGMENT_ID_OFFSET      = FS_SEGMENT_ID_SIZE + FS_FRAG_FULL_LIST_OFFSET       // 下一个可利用的segment id
	FS_FULL_INODE_LIST_OFFSET = FS_FULL_INODE_LIST_SIZE + FS_SEGMENT_ID_OFFSET      // space当前完全占满的segment inode页列表
	FS_FREE_INODE_LIST_OFFSET = FS_FREE_INODE_LIST_SIZE + FS_FULL_INODE_LIST_OFFSET // space当前完全占满的segment inode页列表
)

//  FSP header  104byte
const (
	FS_PAGE_SPACE_SIZE      = 4
	FS_PAGE_MAX_PAGE_SIZE   = 4  // 当前space最大可容纳的page数,文件扩大时才会改变这个值
	FS_PAGE_LIMIT_SIZE      = 4  // 当前space已经分配初始化的page数,包括空闲的和已经使用的
	FS_PAGE_FRAG_USED_SIZE  = 4  //FSP_FREE_FRAG列表中已经被使用的page数
	FS_FREE_LIST_SIZE       = 16 //    space中可用的extent对象列表，extent里面没有一个page被使用
	FS_FRAG_FREE_LIST_SIZE  = 16 // 有可用碎叶page的extent列表，exntent里面有部分page被使用
	FS_FRAG_FULL_LIST_SIZE  = 16 //  没有有可用page的extent列表，exntent里面全部page被使用
	FS_SEGMENT_ID_SIZE      = 8  // 下一个可利用的segment id
	FS_FULL_INODE_LIST_SIZE = 16 //  space当前完全占满的segment inode页列表
	FS_FREE_INODE_LIST_SIZE = 16 //  space当前完全占满的segment inode页列表
)

type FilHeader struct {
	data     *cType.PageData
	pType    uint16
	lastLSN  uint64
	Space    uint32
	Offset   uint32
	prev     uint32
	next     uint32
	flushLSN uint32
}

func (fh *FilHeader) ParseFilHeader(bp *pcache.BuffPage) *FilHeader {
	data := bp.GetData()
	fh.pType = utils.GetUint16(data[FIL_PAGE_TYPE_OFFSET : FIL_PAGE_TYPE_OFFSET+FIL_PAGE_TYPE_SIZE])
	fh.lastLSN = utils.GetUint64(data[FIL_PAGE_LSN_OFFSET : FIL_PAGE_LSN_OFFSET+FIL_PAGE_LSN_SIZE])
	fh.Space = utils.GetUint32(data[FIL_PAGE_SPACE_OFFSET : FIL_PAGE_SPACE_OFFSET+FIL_PAGE_SPACE_SIZE])
	fh.Offset = utils.GetUint32(data[FIL_PAGE_OFFSET_OFFSET : FIL_PAGE_OFFSET_OFFSET+FIL_PAGE_OFFSET_SIZE])
	fh.prev = utils.GetUint32(data[FIL_PAGE_PREV_OFFSET : FIL_PAGE_PREV_OFFSET+FIL_PAGE_PREV_SIZE])
	fh.next = utils.GetUint32(data[FIL_PAGE_NEXT_OFFSET : FIL_PAGE_NEXT_OFFSET+FIL_PAGE_NEXT_SIZE])
	fh.flushLSN = utils.GetUint32(data[FIL_PAGE_FILE_FLUSH_LSN_OFFSET : FIL_PAGE_FILE_FLUSH_LSN_OFFSET+FIL_PAGE_FILE_FLUSH_LSN_SIZE])
	return fh
}

func (fh *FilHeader) SetPtype(pType uint16) {
	copy(fh.data[FIL_PAGE_TYPE_OFFSET:FIL_PAGE_TYPE_OFFSET+FIL_PAGE_TYPE_SIZE], utils.PutUint16(pType))
}

func (fh *FilHeader) SetSpace(space uint32) {
	copy(fh.data[FIL_PAGE_SPACE_OFFSET:FIL_PAGE_SPACE_OFFSET+FIL_PAGE_SPACE_SIZE], utils.PutUint32(space))
}

func (fh *FilHeader) SetOffset(offset uint32) {
	copy(fh.data[FIL_PAGE_OFFSET_OFFSET:FIL_PAGE_OFFSET_OFFSET+FIL_PAGE_OFFSET_SIZE], utils.PutUint32(offset))
}

type FSPHeader struct {
	data          *cType.PageData
	_offset       uint16
	space         uint32
	maxPage       uint32
	limitPage     uint32
	fragUsed      uint32
	FreeList      *FistBaseNode
	FragFreeList  *FistBaseNode
	FragFullList  *FistBaseNode
	segmentID     uint64
	FullInodeList *FistBaseNode
	FreeInodeList *FistBaseNode
}

func newFSPHeader(offset uint16, data *cType.PageData) *FSPHeader {
	fspHeader := new(FSPHeader)
	fspHeader.data = data
	fspHeader._offset = offset
	fspHeader.FreeList = &FistBaseNode{_offset: offset + FS_FREE_LIST_SIZE, data: fspHeader.data}
	fspHeader.FragFreeList = &FistBaseNode{_offset: offset + FS_FRAG_FREE_LIST_OFFSET, data: fspHeader.data}
	fspHeader.FragFullList = &FistBaseNode{_offset: offset + FS_FRAG_FULL_LIST_OFFSET, data: fspHeader.data}
	fspHeader.FullInodeList = &FistBaseNode{_offset: offset + FS_FULL_INODE_LIST_OFFSET, data: fspHeader.data}
	fspHeader.FreeInodeList = &FistBaseNode{_offset: offset + FS_FREE_INODE_LIST_OFFSET, data: fspHeader.data}
	return fspHeader
}

func (fsp *FSPHeader) Space() uint32 {
	return utils.GetUint32(fsp.reOffset(FS_PAGE_SPACE_OFFSET, FS_PAGE_SPACE_SIZE))
}

func (fsp *FSPHeader) LimitPage() uint32 {
	//log.Trace(utils.GetUint32([]byte{0, 0, 0, 64}))
	return utils.GetUint32(fsp.reOffset(FS_PAGE_LIMIT_OFFSET, FS_PAGE_LIMIT_SIZE))
}

func (fsp *FSPHeader) SetSpace(s uint32) {
	fsp.setReOffset(FS_PAGE_SPACE_OFFSET, FS_PAGE_SPACE_SIZE, utils.PutUint32(s))
}

func (fsp *FSPHeader) SetMaxPage(s uint32) {
	fsp.maxPage = s
	copy(fsp.reOffset(FS_PAGE_MAX_PAGE_OFFSET, FS_PAGE_MAX_PAGE_SIZE), utils.PutUint32(s))
}

func (fsp *FSPHeader) SetLimitPage(s uint32) {
	fsp.limitPage = s
	fsp.setReOffset(FS_PAGE_LIMIT_OFFSET, FS_PAGE_LIMIT_SIZE, utils.PutUint32(s))
}

func (fsp *FSPHeader) reOffset(start uint16, end uint16) []byte {
	return fsp.data[fsp._offset+start : fsp._offset+start+end]
}

func (fsp *FSPHeader) setReOffset(start uint16, end uint16, data []byte) {
	copy(fsp.data[fsp._offset+start:fsp._offset+start+end], data)
	//log.Info(fsp.data[fsp._offset+start : fsp._offset+start+end])
}

// page header 50bytes
const (
	PAGE_DIR_SLOTS_SIZE    = 2  //PageDirectory 个数
	PAGE_HEAP_TOP_SIZE     = 2  // 堆中第一个记录偏移量(未分配空间)
	PAGE_N_HEAP_SIZE       = 2  //堆中记录数
	PAGE_FREE_SIZE         = 2  //指向可复用记录
	PAGE_GARBAGE_SIZE      = 2  // 记录中已经删除字节数
	PAGE_LAST_INSERT_SIZE  = 2  // 最近一次插入偏移量
	PAGE_DIRECTION_SIZE    = 2  // 插入方向用于插入优化
	PAGE_N_DIRECTION_SIZE  = 2  // 相同方向插入数量
	PAGE_N_RECS_SIZE       = 2  // 该页面中记录个数
	PAGE_MAX_TRX_ID_SIZE   = 8  // 修改当前页最大事务ID(仅在二级索引页中定义)
	PAGE_LEVEL_SIZE        = 2  //当前页在索引树中的位置
	PAGE_INDEX_ID_SIZE     = 8  //当前页所在索引ID
	PAGE_BTR_SEG_LEAF_SIZE = 10 //数据页叶节点
	PAGE_BTR_SEG_TOP_SIZE  = 10 // 数据页非页节点
)

const (
	PAGE_DIR_SLOYS_OFFSET     = 0
	PAGE_HEAP_TOP_OFFSET      = PAGE_DIR_SLOYS_OFFSET + PAGE_HEAP_TOP_SIZE
	PAGE_N_HEAP_OFFSET        = PAGE_HEAP_TOP_OFFSET + PAGE_N_HEAP_SIZE
	PAGE_FREE_OFFSET          = PAGE_N_HEAP_OFFSET + PAGE_FREE_SIZE
	PAGE_GARBAGE_OFFSET       = PAGE_FREE_OFFSET + PAGE_GARBAGE_SIZE
	PAGE_LAST_INSERT_OFFSET   = PAGE_GARBAGE_OFFSET + PAGE_LAST_INSERT_SIZE
	PAGE_DIRECTION_OFFSET     = PAGE_LAST_INSERT_OFFSET + PAGE_DIRECTION_SIZE
	PAGE_N_DIRECTION_OFFSET   = PAGE_DIRECTION_OFFSET + PAGE_N_DIRECTION_SIZE
	PAGE_N_RECS_OFFSET        = PAGE_N_DIRECTION_OFFSET + PAGE_N_RECS_SIZE
	PAGE_MAX_TRX_ID_OFFSET    = PAGE_N_RECS_OFFSET + PAGE_MAX_TRX_ID_SIZE
	PAGE_LEVEL_OFFSET         = PAGE_MAX_TRX_ID_OFFSET + PAGE_LEVEL_SIZE
	PAGE_INDEX_ID_OFFSET      = PAGE_LEVEL_OFFSET + PAGE_INDEX_ID_SIZE
	PAGE_BTR_SEG_LEAF_OFFSET = PAGE_INDEX_ID_OFFSET + PAGE_BTR_SEG_LEAF_SIZE
	PAGE_BTR_SEG_TOP_OFFSET   = PAGE_BTR_SEG_LEAF_OFFSET + PAGE_BTR_SEG_TOP_SIZE
)

type IndexHeader struct {
	data    *cType.PageData
	_offset uint16
}

func (idx *IndexHeader) LeafSegment() (spaceId uint32, pageNo uint32, offset uint16) {
	return utils.GetUint32(idx.reOffset(PAGE_BTR_SEG_TOP_OFFSET, PAGE_BTR_SEG_TOP_SIZE)),
	utils.GetUint32(idx.reOffset(PAGE_BTR_SEG_LEAF_OFFSET+4, 4)),
	utils.GetUint16(idx.reOffset(PAGE_BTR_SEG_LEAF_OFFSET+6, 2))
}

func (idx *IndexHeader) TopSegment() (spaceId uint32, pageNo uint32, offset uint16) {
	return utils.GetUint32(idx.reOffset(PAGE_BTR_SEG_TOP_OFFSET, PAGE_BTR_SEG_TOP_SIZE)),
		utils.GetUint32(idx.reOffset(PAGE_BTR_SEG_TOP_OFFSET+4, 4)),
		utils.GetUint16(idx.reOffset(PAGE_BTR_SEG_TOP_OFFSET+6, 2))
}

func (idx *IndexHeader) SetTopSegment(spaceId uint32, pageNo uint32, offset uint16) {
	idx.setReOffset(PAGE_BTR_SEG_TOP_OFFSET, PAGE_BTR_SEG_TOP_SIZE, utils.PutUint32(spaceId))
	idx.setReOffset(PAGE_BTR_SEG_TOP_OFFSET+4, 4, utils.PutUint32(pageNo))
	idx.setReOffset(PAGE_BTR_SEG_TOP_OFFSET+6, 2, utils.PutUint16(offset))
}

func (idx *IndexHeader) SetLeafSegment(spaceId uint32, pageNo uint32, offset uint16) {
	idx.setReOffset(PAGE_BTR_SEG_LEAF_OFFSET, PAGE_BTR_SEG_LEAF_SIZE, utils.PutUint32(spaceId))
	idx.setReOffset(PAGE_BTR_SEG_LEAF_OFFSET+4, 4, utils.PutUint32(pageNo))
	idx.setReOffset(PAGE_BTR_SEG_LEAF_OFFSET+6, 2, utils.PutUint16(offset))
}

func (idx *IndexHeader) reOffset(start uint16, end uint16) []byte {
	return idx.data[idx._offset+start : idx._offset+start+end]
}
func (idx *IndexHeader) setReOffset(start uint16, end uint16, data []byte) {
	log.Info(log.AnyViewToString(idx),111111)
	copy(idx.data[idx._offset+start:idx._offset+start+end], data)
}
