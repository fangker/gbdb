package cType

// page cType

type pageType uint16

const (
	FIL_PAGE_INDEX     pageType = iota // 索引页
	FIL_PAGE_UNDO_LOGO                 // undolog 页
	FIL_PAGE_TYPE_XDES                 // 簇描述页

)
const (
	PAGE_SIZE = 16384
	REDO_BLOCK_SIZE = 512
)

type PageData [PAGE_SIZE] byte
type RedoData [REDO_BLOCK_SIZE] byte
