package page

import (
	"github.com/fangker/gbdb/backend/dm/buffPage"
	"github.com/fangker/gbdb/backend/dm/constants/cType"
	"github.com/fangker/gbdb/backend/utils"
	"github.com/fangker/gbdb/backend/cache"
)

const (
	FSPAGE_FH_OFFSET   = 0
	FSPAGE_FSPH_OFFSET = 34
	FSPAGE_XDES_OFFSET = 139
)

type FSPage struct {
	FH  *FilHeader
	FSH *FSPHeader
	BP  *pcache.BuffPage
	data *cType.PageData
}

func NewFSPage(bp *pcache.BuffPage) *FSPage {
	fsPage := &FSPage{
		data: bp.GetData(),
		FH: &FilHeader{data: bp.GetData()},
		BP: bp,
		FSH: newFSPHeader(FSPAGE_FSPH_OFFSET,bp.GetData()),
	}
	fsPage.FH.SetPtype(PAGE_TYPE_FSP)
	fsPage.FH.SetSpace(0)
	fsPage.FH.SetOffset(0)
	return fsPage
}

func (fsp *FSPage) InitSysExtend(cache *cache.CachePool) {
	fsp.FSH.FragFreeList.SetFirst(0,FSPAGE_XDES_OFFSET)
	fsp.FSH.FragFreeList.SetLast(0,FSPAGE_XDES_OFFSET)
	fsp.FSH.FragFreeList.SetLen(1)
	copy(fsp.data[FSPAGE_XDES_OFFSET:FSPAGE_XDES_OFFSET+XDES_ENTRY_SIZE*1],utils.PutUint32(1))
	fsp.setUsedExtendPage(1)
	fsp.setUsedExtendPage(2)
	fsp.setUsedExtendPage(3)
	fsp.setUsedExtendPage(4)
	fsp.setUsedExtendPage(5)
}

func (fsp *FSPage) setUsedExtendPage(p int) {
	var enum =[5]uint8{0,192,48,12,3}
	remain:=enum[(p%64)%4]
	offset:=FSPAGE_XDES_OFFSET+int(p/64)*XDES_ENTRY_SIZE+int((p%64)/8)
	remain=[]byte(fsp.data[offset+24:offset+25])[0]|remain
	copy(fsp.data[offset+24:offset+25],[]byte{remain})
}

func (fsp *FSPage) SetFreeInodFirst(page uint32,offset uint16) {
	fsp.FSH.freeInodeList.SetFirst(page,offset)
}
func (fsp *FSPage) SetFreeInodeLen(len uint32) {
	fsp.FSH.freeInodeList.SetLen(len)
}

func( FSPage) GetFreePage(data *pcache.BuffPage,offset uint16){

}