package page

import (
	// "github.com/fangker/gbdb/backend/cache"
	"github.com/fangker/gbdb/backend/cache"
	"github.com/fangker/gbdb/backend/cache/buffPage"
	"github.com/fangker/gbdb/backend/constants/cType"
	"github.com/fangker/gbdb/backend/utils"

)

const (
	FSPAGE_FH_OFFSET   = 0
	FSPAGE_FSPH_OFFSET = 34
	FSPAGE_XDES_OFFSET = 139
)

type FSPage struct {
	FH           *FilHeader
	FSH          *FSPHeader
	BP           *pcache.BuffPage
	data         *cType.PageData
	cacheWrapper cache.Wrapper
}

func NewFSPage(bp *pcache.BuffPage) *FSPage {
	fsPage := &FSPage{
		data: bp.GetData(),
		FH:   &FilHeader{data: bp.GetData()},
		BP:   bp,
		FSH:  newFSPHeader(FSPAGE_FSPH_OFFSET, bp.GetData()),
	}
	fsPage.FH.SetPtype(PAGE_TYPE_FSP)
	fsPage.FH.SetSpace(0)
	fsPage.FH.SetOffset(0)
	return fsPage
}
func (fsp *FSPage) SetCacheWrapper(c cache.Wrapper) {
	fsp.cacheWrapper = c
}

func (fsp *FSPage) InitSysExtend() {
	fsp.FSH.FragFreeList.SetFirst(0, FSPAGE_XDES_OFFSET)
	fsp.FSH.FragFreeList.SetLast(0, FSPAGE_XDES_OFFSET)
	fsp.FSH.FragFreeList.SetLen(1)
	// 设置第一个segment被占用
	copy(fsp.data[FSPAGE_XDES_OFFSET:FSPAGE_XDES_OFFSET+XDES_ENTRY_SIZE*1], utils.PutUint32(1))
	fsp.setUsedExtendPage(0)
	fsp.setUsedExtendPage(1)
	fsp.setUsedExtendPage(2)
	fsp.setUsedExtendPage(3)
	fsp.setUsedExtendPage(4)
	fsp.setUsedExtendPage(5)
	fsp.setUsedExtendPage(6)
	fsp.setUsedExtendPage(7)
	fsp.setUsedExtendPage(8)
}

func (fsp *FSPage) InitSysUndoExtend() {
	fsp.FSH.FragFreeList.SetFirst(0, FSPAGE_XDES_OFFSET)
	fsp.FSH.FragFreeList.SetLast(0, FSPAGE_XDES_OFFSET)
	fsp.FSH.FragFreeList.SetLen(1)
}


func (fsp *FSPage) setUsedExtendPage(p int) {
	i := (p+1)/64 + 1
	site := int(p / 4)
	mod := uint(p % 4)
	offset := FSPAGE_XDES_OFFSET + (site+24)*i
	pos := &fsp.data[offset]
	*pos = *pos | (3 << ((3 - mod) * 2))

}

func (fsp *FSPage) SetFreeInodFirst(page uint32, offset uint16) {
	fsp.FSH.freeInodeList.SetFirst(page, offset)
}
func (fsp *FSPage) SetFreeInodeLen(len uint32) {
	fsp.FSH.freeInodeList.SetLen(len)
}

// page 为extend page
func GetFragFreePage(wrap cache.Wrapper, page uint32, offset uint16) uint32 {
	fpge := cache.CP.GetPage(wrap, page)
	fsp_bp := NewFSPage(fpge)
	xdes := parseXdes(fsp_bp.data[offset : offset+XDES_ENTRY_SIZE])
	var freePage int
	for k, v := range xdes.BitMap() {
		if(v==255){
			continue
		}
		for i := uint(3); i >= 0; i = i - 1 {
			if (2<<(i*2)&v)>>(i*2) != 2 {
				freePage = k*4 + int(4-i) - 1
				goto result
			}
		}
	}
result:
	return uint32(freePage) + page*256*64
}

func SetUsedPage(wrap cache.Wrapper, p uint32) {
	xdesPageNo := uint32(p/64*256) * 64 * 256
	mod := int(xdesPageNo % (64 * 256))
	if (mod == 0) {
		mod = int(p)
	}
	xdes := NewFSPage(cache.CP.GetPage(wrap, xdesPageNo))
	xdes.setUsedExtendPage(mod)
}
