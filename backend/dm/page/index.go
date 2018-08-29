package page

import (
	"github.com/fangker/gbdb/backend/cache/buffPage"
	"github.com/fangker/gbdb/backend/constants/cType"
)


type IndexPage struct {
	FH           *FilHeader
	IH           *IndexHeader
	BP           *pcache.BuffPage
	data         *cType.PageData
}
