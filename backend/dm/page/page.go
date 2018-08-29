package page

import (
	"github.com/fangker/gbdb/backend/cache/buffPage"
)


const (
	PAGE_TYPE_FSP   = 0
	PAGE_TYPE_INODE = 1
	PAGE_TYPE_PAGE  = 2
	PAGE_TYPE_INDEX = 3
)


type Page struct {
	FH   *FilHeader
	BP   *pcache.BuffPage

}

func NewPage(bf *pcache.BuffPage) *Page {
	page := &Page{FH: &FilHeader{data:bf.GetData()}, BP: bf}
	page.FH.ParseFilHeader(bf)
	return page
}
