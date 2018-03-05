package page

// page header
var (
	FIL_PAGE_SPACE          [4]byte //   page所属的表空间的space id
	FIL_PAGE_OFFSET         [4]byte //   page no，一般是在表空间的物理偏移量
	FIL_PAGE_PREV           [4]byte //    前一页的page no (B+tree的叶子节点是通过链表串起来的，有前后关系
	FIL_PAGE_NEXT           [4]byte //      后一页的page no
	FIL_PAGE_LSN            [8]byte //   最后被修改的LSN日志号
	FIL_PAGE_FILE_FLUSH_LSN [8]byte //   该表空间最会一次被更新的LSN号
	FIL_PAGE_TYPE           [2]byte //        page的类型
)
