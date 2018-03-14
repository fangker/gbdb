package page

import "github.com/fangker/gbdb/backend/utils"
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
	FIL_PAGE_SPACE_SIZE          = 4  //   page所属的表空间的space id
	FIL_PAGE_OFFSET_SIZE         = 4  //   page no，一般是在表空间的物理偏移量
	FIL_PAGE_PREV_SIZE           = 4  //    前一页的page no (B+tree的叶子节点是通过链表串起来的，有前后关系
	FIL_PAGE_NEXT_SIZE           = 4 //      后一页的page no
	FIL_PAGE_LSN_SIZE            = 8 //   最后被修改的LSN日志号
	FIL_PAGE_FILE_FLUSH_LSN_SIZE = 8 //   该表空间最后一次被更新的LSN号
	FIL_PAGE_TYPE_SIZE           = 2 //        page的类型
)

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

type FilHead struct {
	pType    uint16
	lastLSN  uint64
	space    uint32
	offset   uint32
	prev     uint32
	next     uint32
	flushLSN uint32
}

func (fh *FilHead) parseFilHeader(pd *PageData) *FilHead {
	data := pd
	fh.pType = utils.GetUint16(data[FIL_PAGE_TYPE:FIL_PAGE_TYPE+FIL_PAGE_TYPE_SIZE])
	fh.lastLSN = utils.GetUint64(data[FIL_PAGE_LSN:FIL_PAGE_LSN+FIL_PAGE_LSN_SIZE])
	fh.space = utils.GetUint32(data[FIL_PAGE_SPACE:FIL_PAGE_SPACE+FIL_PAGE_SPACE_SIZE])
	fh.offset = utils.GetUint32(data[FIL_PAGE_OFFSET:FIL_PAGE_OFFSET+FIL_PAGE_OFFSET_SIZE])
	fh.prev = utils.GetUint32(data[FIL_PAGE_PREV:FIL_PAGE_PREV+FIL_PAGE_PREV_SIZE])
	fh.next = utils.GetUint32(data[FIL_PAGE_NEXT:FIL_PAGE_NEXT+FIL_PAGE_NEXT_SIZE])
	fh.flushLSN = utils.GetUint32(data[FIL_PAGE_FILE_FLUSH_LSN:FIL_PAGE_FILE_FLUSH_LSN+FIL_PAGE_FILE_FLUSH_LSN_SIZE])
	return fh
}