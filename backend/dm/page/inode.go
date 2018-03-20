package page

// Inode page

var (
	FSEG_INODE_PAGE_NODE [12]byte // 头链表
	FSEG_ID  [8]byte // segment ID
	FSEG_NOT_FULL_N_USED [8]byte // 在链表上已经被使用的page数
	FSEG_FREE [16]byte
	FSEG_FULL [16]byte
	FSEG_INODE_ENTRY_INIT_MARK [4]byte //标识位初始化
)


type INodePage struct {
	FH  FilHeader
	data *PageData
}

