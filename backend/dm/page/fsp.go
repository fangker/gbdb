package page

import (
	"github.com/fangker/gbdb/backend/dm/buffPage"
)

const (
	FSPAGE_FH_OFFSET   = 0
	FSPAGE_FSPH_OFFSET = 34
)

type FSPage struct {
	FH  *FilHeader
	FSH *FSPHeader
	BP  *pcache.BuffPage
}

func NewFSPage(bp *pcache.BuffPage) *FSPage {
	fsPage := &FSPage{
		FH: &FilHeader{data: bp.GetData()},
		BP: bp,
		FSH: newFSPHeader(FSPAGE_FSPH_OFFSET,bp.GetData()),
	}
	fsPage.FH.SetPtype(PAGE_TYPE_FSP)
	fsPage.FH.SetSpace(0)
	fsPage.FH.SetOffset(0)
	return fsPage
}
