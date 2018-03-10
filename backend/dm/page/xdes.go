package page

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
	XDES_BITMA     [16]byte //表示簇的页使用状态
)
