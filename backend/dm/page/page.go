package page

// file header 34bytes
var (
	FIL_PAGE_SPACE          [4]byte //   page所属的表空间的space id
	FIL_PAGE_OFFSET         [4]byte //   page no，一般是在表空间的物理偏移量
	FIL_PAGE_PREV           [4]byte //    前一页的page no (B+tree的叶子节点是通过链表串起来的，有前后关系
	FIL_PAGE_NEXT           [4]byte //      后一页的page no
	FIL_PAGE_LSN            [8]byte //   最后被修改的LSN日志号
	FIL_PAGE_FILE_FLUSH_LSN [8]byte //   该表空间最会一次被更新的LSN号
	FIL_PAGE_TYPE           [2]byte //        page的类型
)

// page headder
var (
	PAGE_DIR_SLOTS [2]byte  //PageDirectory 个数
	PAGE_HEAP_TOP [2]byte // 堆中第一个记录偏移量(未分配空间)
	PAGE_N_HEAP [2]byte //堆中记录数
	PAGE_FREE [2]byte //指向可复用记录
	PAGE_GARBAGE [2]byte // 记录中已经删除字节数
    PAGE_N_RECS [2]byte // 该页面中记录个数
    PAGE_MAX_TRX_ID [8]byte // 修改当前页最大事务ID(仅在二级索引页中定义)
    PAGE_LEVEL [2]byte //当前页在索引树中的位置
    PAGE_INDEX_ID [8]byte //当前页所在索引ID
    PAGE_BTR_SEGEMENT_LEAF [10]byte //数据页叶节点
    PAGE_BTR_SEG_TOP [10]byte // 数据页非页节点
)