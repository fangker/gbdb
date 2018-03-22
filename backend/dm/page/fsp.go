package page

import (
	"github.com/fangker/gbdb/backend/dm/buffPage"
)
type FSPage struct {
	FH       FilHeader
	FSH      FSHeader
	BP       *pcache.BuffPage
}

func NewFSPage(bp *pcache.BuffPage) *FSPage {
	fsPage:= &FSPage{FH:FilHeader{data:bp.GetData()},BP:bp}
	fsPage.FH.ParseFilHeader(bp)
	fsPage.FH.SetPtype(PAGE_TYPE_FSP)
	fsPage.FH.SetSpace(0)
	fsPage.FH.SetOffset(0)
	return  fsPage
}

