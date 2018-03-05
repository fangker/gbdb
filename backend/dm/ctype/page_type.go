package cType

// page type

type pageType uint16

const (
	FIL_PAGE_INDEX pageType = iota // 索引页
	FIL_PAGE_UNDO_LOGO // undolog 页
	FIL_PAGE_TYPE_XDES // 簇描述页

)