package page

import (
	"github.com/fangker/gbdb/backend/cache"
	"github.com/fangker/gbdb/backend/cache/buffPage"
	"github.com/fangker/gbdb/backend/constants/cType"
	"github.com/fangker/gbdb/backend/wrapper"
	"math"
	_ "github.com/fangker/gbdb/backend/utils/log"
)

const (
	FSPAGE_FH_OFFSET   = 0
	FSPAGE_FSPH_OFFSET = 34
	FSPAGE_XDES_OFFSET = 138
)

var cachePool *cache.CachePool

func AttachCache() {
	cachePool = cache.CP
}

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
	return fsPage
}

func (fsp *FSPage) Data() *cType.PageData {
	return fsp.data
}

func (fsp *FSPage) SetCacheWrapper(c wp.Wrapper) {
	fsp.wp = c
}

func (fsp *FSPage) InitSysExtend() {
	//TODO: fragfree initialize
	fsp.FSH.FragFreeList.SetFirst(NPos(0, FSPAGE_XDES_OFFSET))
	fsp.FSH.FragFreeList.SetLast(NPos(0, FSPAGE_XDES_OFFSET))
	fsp.FSH.FragFreeList.SetLen(1)

	// 设置第一个segment被占用
	fsp.FSH.SetLimitPage(64)
	// 扩展簇到128page
	fsp.extendXdesSpace()
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
	fsp.FSH.FragFreeList.SetFirst(NPos(0, FSPAGE_XDES_OFFSET))
	fsp.FSH.FragFreeList.SetLast(NPos(0, FSPAGE_XDES_OFFSET))
	fsp.FSH.FragFreeList.SetLen(1)

}

// 设置页已经使用  Ok
func (fsp *FSPage) setUsedExtendPage(p int) {
	i := (p+1)/64 + 1
	site := int(p / 4)
	mod := uint(p % 4)
	offset := FSPAGE_XDES_OFFSET + (site+24)*i
	pos := &fsp.data[offset]
	*pos = *pos | (3 << ((3 - mod) * 2))
}

// 获得XdexEntry
func (fsp *FSPage) GetXdesEntryByPageNo(i uint32) XdesEntry {
	// 0 - 0
	fsNo := uint32(256 * (math.Ceil(float64(i)/256) - 1))
	fsPage := fsp
	if (fsp.BP.PageNo() != fsNo) {
		fsPage = NewFSPage(cachePool.GetPage(fsp.wp, fsNo))
	}
	fsNo = fsPage.BP.PageNo()
	oi := i - 256*fsNo
	offset := uint16(FSPAGE_XDES_OFFSET + oi/64)
	return parseXDES(fsp.wp, NPos(fsNo, offset))
}

// 从簇中获取空白页
func GetFreePageByXdes(wrap wp.Wrapper, page uint32, offset uint16) uint32 {
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
	p := uint32(freePage) + page*256*64;
	SetUsedPage(wrap, p)
	return p
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

// 获取fsp_0
func getSpaceFsp(wp wp.Wrapper) *FSPage {
	return NewFSPage(cache.CP.GetPage(wp, 0))
}

// 扩展簇 返回偏移
func (fsp *FSPage) extendXdesSpace() {
	// 寻找未使用
	limit := math.Ceil(float64(fsp.FSH.LimitPage() / 256))
	extendToPage := fsp.FSH.LimitPage() + 64
	extend := math.Ceil(float64((extendToPage) / 256))
	// todo: 需要初始化页头
	fsp.FSH.SetLimitPage(extendToPage)
	if (limit < extend) {
		// 需初始化新fsp页 移动到最后
		fsp = NewFSPage(cachePool.GetPage(fsp.wp, uint32(math.Ceil(float64((fsp.FSH.LimitPage()+64)/256)-1))))
		// fsp要被加入到frag链表

	}
	// 初始化后放入extend free 链表
	xe := fsp.GetXdesEntryByPageNo(extendToPage)
	fsp0 := getSpaceFsp(fsp.wp)
	fsp0.FSH.FreeList.AddToLast(xe.xdesNode)
}

// 扩展 inode page
func (fsp *FSPage) ExtendInodePage(pageNo uint32) {
	fsp.setUsedExtendPage(int(pageNo))
	fsp.FSH.FreeInodeList.AddToLast(NewINodePage(cachePool.GetPage(fsp.wp, pageNo)).INL)
}

// extend inode(segment) - add xdes
func (fsp *FSPage) AddExtentToSegment(p uint32,ofs uint16){

}