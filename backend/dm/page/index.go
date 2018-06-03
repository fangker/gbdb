package page

import (
	"github.com/fangker/gbdb/backend/dm/buffPage"
	"github.com/fangker/gbdb/backend/dm/constants/cType"
)


type IndexPage struct {
	FH           *FilHeader
	IH           *IndexHeader
	BP           *pcache.BuffPage
	data         *cType.PageData
}
