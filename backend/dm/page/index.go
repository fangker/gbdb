package page

import (
	"github.com/fangker/gbdb/backend/cache/buffPage"
	"github.com/fangker/gbdb/backend/constants/cType"
)

type IndexPage struct {
	FH   *FilHeader
	IH   *IndexHeader
	BP   *pcache.BuffPage
	data *cType.PageData
}

func NewIndexPage(bf *pcache.BuffPage) *IndexPage {
	index := &IndexPage{
		FH: &FilHeader{data: bf.GetData()},
		IH: &IndexHeader{data:bf.GetData(),_offset: FIL_HEADER_END},
	}
	return index
}
