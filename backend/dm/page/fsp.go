package page

import (
	// "github.com/fangker/gbdb/backend/cache"
	"github.com/fangker/gbdb/backend/cache"
	"github.com/fangker/gbdb/backend/cache/buffPage"
	"github.com/fangker/gbdb/backend/constants/cType"
	"github.com/fangker/gbdb/backend/utils"
	"github.com/fangker/gbdb/backend/wrapper"
	"math"
)

const (
	FSPAGE_FH_OFFSET   = 0
	FSPAGE_FSPH_OFFSET = 34
	FSPAGE_XDES_OFFSET = 139
)

var cachePool = cache.CP

type FSPage struct {
	FH   *FilHeader
	FSH  *FSPHeader
	BP   *pcache.BuffPage
	data *cType.PageData
	wp   wp.Wrapper
}

func NewFSPage(bp *pcache.BuffPage) *FSPage {
	fsPage := &FSPage{
		data: bp.GetData(),
		FH:   &FilHeader{data: bp.GetData()},
		BP:   bp,
		FSH:  newFSPHeader(FSPAGE_FSPH_OFFSET, bp.GetData()),
	}
	fsPage.FH.SetPtype(PAGE_TYPE_FSP)
	//fsPage.FH.SetSpace(0)
	//fsPage.FH.SetOffset(0)
	return fsPage
}
func (fsp *FSPage) SetCacheWrapper(c wp.Wrapper) {
	fsp.wp = c
}

func (fsp *FSPage) InitSysExtend() {
	fsp.FSH.FragFreeList.SetFirst(0, FSPAGE_XDES_OFFSET)
	fsp.FSH.FragFreeList.SetLast(0, FSPAGE_XDES_OFFSET)
	fsp.FSH.FragFreeList.SetLen(1)
	// 设置第一个segment被占用
	fsp.FSH.SetLimitPage(64)
	fsp.extendXdesSpace(0)
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
// 设置页已经使用
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


func GetXDESEntry(i int) uint32 {
	fsNo := 256*(math.Ceil(float64(i/256)-1))
	   offset - (i%256)

}
// page 为extend page
func GetFragFreePage(wrap wp.Wrapper, page uint32, offset uint16) uint32 {
	fpge := cache.CP.GetPage(wrap, page)
	fsp_bp := NewFSPage(fpge)
	xdes := parseXdes(fsp_bp.data[offset : offset+XDES_ENTRY_SIZE])
	var freePage int
	for k, v := range xdes.BitMap() {
		if (v == 255) {
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
	return uint32(freePage) + page*256*64;
}

func SetUsedPage(wp wp.Wrapper, p uint32) {
	xdesPageNo := uint32(p/64*256) * 64 * 256
	mod := int(xdesPageNo % (64 * 256))
	if (mod == 0) {
		mod = int(p)
	}
	xdes := NewFSPage(cache.CP.GetPage(wp, xdesPageNo))
	xdes.setUsedExtendPage(mod)
}

func addToFreeInodeList() {

}

func addToFullInodeList() {

}

func getSpaceFsp(wp wp.Wrapper) *FSPage {
	 return NewFSPage(cache.CP.GetPage(wp, 0))
}
// 扩展簇 返回偏移量
func (fsp *FSPage) extendXdesSpace(i int)  {
	// 寻找未使用
	fsp.FSH.LimitPage()

}