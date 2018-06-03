package page

import (
	"github.com/fangker/gbdb/backend/utils"
	"github.com/fangker/gbdb/backend/dm/buffPage"
	"github.com/fangker/gbdb/backend/dm/constants/cType"
)

const (
	FIL_HEADER_OFFSET = 33
)

// file header 34bytes
const (
	FIL_PAGE_SPACE          = 0  //   page所属的表空间的space id
	FIL_PAGE_OFFSET         = 3  //   page no，一般是在表空间的物理偏移量
	FIL_PAGE_PREV           = 7  //    前一页的page no (B+tree的叶子节点是通过链表串起来的，有前后关系
	FIL_PAGE_NEXT           = 11 //      后一页的page no
	FIL_PAGE_LSN            = 15 //   最后被修改的LSN日志号
	FIL_PAGE_FILE_FLUSH_LSN = 23 //   该表空间最后一次被更新的LSN号
	FIL_PAGE_TYPE           = 31 //        page的类型
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
	FS_PAGE_SPACE      = 0
	FS_PAGE_MAX_PAGE   = 3  // 当前space最大可容纳的page数,文件扩大时才会改变这个值
	FS_PAGE_LIMIT      = 7  // 当前space已经分配初始化的page数,包括空闲的和已经使用的
	FS_PAGE_FRAG_USED  = 11 // FSP_FREE_FRAG列表中已经被使用的page数
	FS_FREE_LIST       = 15 // space中可用的extent对象列表，extent里面没有一个page被使用
	FS_FRAG_FREE_LIST  = 31 // 有可用碎叶page的extent列表，exntent里面有部分page被使用
	FS_FRAG_FULL_LIST  = 47 // 没有有可用page的extent列表，exntent里面全部page被使用
	FS_SEGMENT_ID      = 53 // 下一个可利用的segment id
	FS_FULL_INODE_LIST = 61 // space当前完全占满的segment inode页列表
	FS_FREE_INODE_LIST = 77 // space当前完全占满的segment inode页列表
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
	fh.pType = utils.GetUint16(data[FIL_PAGE_TYPE:FIL_PAGE_TYPE+FIL_PAGE_TYPE_SIZE])
	fh.lastLSN = utils.GetUint64(data[FIL_PAGE_LSN:FIL_PAGE_LSN+FIL_PAGE_LSN_SIZE])
	fh.Space = utils.GetUint32(data[FIL_PAGE_SPACE:FIL_PAGE_SPACE+FIL_PAGE_SPACE_SIZE])
	fh.Offset = utils.GetUint32(data[FIL_PAGE_OFFSET:FIL_PAGE_OFFSET+FIL_PAGE_OFFSET_SIZE])
	fh.prev = utils.GetUint32(data[FIL_PAGE_PREV:FIL_PAGE_PREV+FIL_PAGE_PREV_SIZE])
	fh.next = utils.GetUint32(data[FIL_PAGE_NEXT:FIL_PAGE_NEXT+FIL_PAGE_NEXT_SIZE])
	fh.flushLSN = utils.GetUint32(data[FIL_PAGE_FILE_FLUSH_LSN:FIL_PAGE_FILE_FLUSH_LSN+FIL_PAGE_FILE_FLUSH_LSN_SIZE])
	return fh
}

func (fh *FilHeader) SetPtype(pType uint16) {
	copy(fh.data[FIL_PAGE_TYPE:FIL_PAGE_TYPE+FIL_PAGE_TYPE_SIZE], utils.PutUint16(pType))
}

func (fh *FilHeader) SetSpace(space uint32) {
	copy(fh.data[FIL_PAGE_SPACE:FIL_PAGE_SPACE+FIL_PAGE_SPACE_SIZE], utils.PutUint32(space))
}

func (fh *FilHeader) SetOffset(offset uint32) {
	copy(fh.data[FIL_PAGE_OFFSET:FIL_PAGE_OFFSET+FIL_PAGE_OFFSET_SIZE], utils.PutUint32(offset))
}

type FSPHeader struct {
	data          *cType.PageData
	_offset       int
	space         uint32
	maxPage       uint32
	limitPage     uint32
	fragUsed      uint32
	freeList      uint64
	FragFreeList  *FistBaseNode
	FragFullList  *FistBaseNode
	segmentID     uint64
	fullInodeList *FistBaseNode
	freeInodeList *FistBaseNode
}

func newFSPHeader(offset int, data *cType.PageData) *FSPHeader {
	fspHeader := new(FSPHeader)
	fspHeader.data = data
	fspHeader._offset = offset
	fspHeader.FragFreeList = &FistBaseNode{_offset: offset + FS_FRAG_FREE_LIST, data: fspHeader.data}
	fspHeader.FragFullList = &FistBaseNode{_offset: offset + FS_FRAG_FULL_LIST, data: fspHeader.data}
	fspHeader.fullInodeList = &FistBaseNode{_offset: offset + FS_FULL_INODE_LIST, data: fspHeader.data}
	fspHeader.freeInodeList = &FistBaseNode{_offset: offset + FS_FREE_INODE_LIST, data: fspHeader.data}
	return fspHeader
}


func (fsp *FSPHeader)Space() uint32 {
	return utils.GetUint32(fsp.reOffset(FS_PAGE_SPACE,FS_PAGE_SPACE_SIZE))
}

func (fsp *FSPHeader)setSpace (s uint32){
	copy(fsp.reOffset(FS_PAGE_SPACE,FS_PAGE_SPACE_SIZE),utils.PutUint32(s))
}

func (fsp *FSPHeader)SetMaxPage (s uint32){
	fsp.maxPage=s
	copy(fsp.reOffset(FS_PAGE_MAX_PAGE,FS_PAGE_MAX_PAGE_SIZE),utils.PutUint32(s))
}

func (fsp *FSPHeader)SetLimitPage (s uint32){
	fsp.limitPage=s
	copy(fsp.reOffset(FS_PAGE_LIMIT,FS_PAGE_LIMIT_SIZE),utils.PutUint32(s))
}

func (fsp *FSPHeader) reOffset(start int, end int) []byte {
	return fsp.data[fsp._offset+start:fsp._offset+end]
}

// page header 50bytes
var (
	PAGE_DIR_SLOTS         [2]byte  //PageDirectory 个数
	PAGE_HEAP_TOP          [2]byte  // 堆中第一个记录偏移量(未分配空间)
	PAGE_N_HEAP            [2]byte  //堆中记录数
	PAGE_FREE              [2]byte  //指向可复用记录
	PAGE_GARBAGE           [2]byte  // 记录中已经删除字节数
	PAGE_N_RECS            [2]byte  // 该页面中记录个数
	PAGE_MAX_TRX_ID        [8]byte  // 修改当前页最大事务ID(仅在二级索引页中定义)
	PAGE_LEVEL             [2]byte  //当前页在索引树中的位置
	PAGE_INDEX_ID          [8]byte  //当前页所在索引ID
	PAGE_BTR_SEGEMENT_LEAF [10]byte //数据页叶节点
	PAGE_BTR_SEG_TOP       [10]byte // 数据页非页节点
)

type IndexHeader struct{
	data     *cType.PageData
	_offset  int
}

func (id *IndexHeader) reOffset(start int, end int) []byte {
	return id.data[id._offset+start:id._offset+end]
}