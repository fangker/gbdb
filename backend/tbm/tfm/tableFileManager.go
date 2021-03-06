package tfm

import (
	"github.com/fangker/gbdb/backend/cache"
	"github.com/fangker/gbdb/backend/cache/buffPage"
	"github.com/fangker/gbdb/backend/def/cType"
	"github.com/fangker/gbdb/backend/dm/page"
	"os"
	"github.com/fangker/gbdb/backend/wrapper"
	//"github.com/fangker/gbdb/backend/utils/log"
)

var cachePool *cache.Pool

type TableFileManage struct {
	CacheBuffer  *cache.Pool
	FilePath     string
	cacheWrapper wp.Wrapper
}

type TableFileManager interface {
	SysDir() *page.DictPage
}

func (sm *TableFileManage) CacheWrapper() wp.Wrapper {
	return sm.cacheWrapper
}

func NewTableFileManage(spaceID, tableID uint32, filePath string) *TableFileManage {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	tfm := &TableFileManage{cacheWrapper: wp.GetWrapper(spaceID, tableID, file), FilePath: filePath, CacheBuffer: cache.CP}
	return tfm
}

func (sm *TableFileManage) writeSync(pageNum uint32, data cType.PageData) {
	offset := pageNum * cType.PAGE_SIZE
	sm.cacheWrapper.File.WriteAt(data[:], int64(offset))
	sm.cacheWrapper.File.Sync()
}

//// 初始化一个文件 设定初始页面构造文件结构
//func (sm *TableFileManage) InitFileStructure() {
//
//}
//
//// 初始化一系统表文件 设定初始页面构造文件结构
//func (sm *TableFileManage) InitSysFileStructure() {
//
//}
//

func (sm *TableFileManage) getFlushPage(pageNo uint32) *pcache.BuffPage {
	return sm.CacheBuffer.GetFlushPage(sm.cacheWrapper, pageNo)
}

func (sm *TableFileManage) InitSysFile() {
	fsp_bp := sm.getFlushPage(0)
	fsp_bp.Lock()
	fsp := page.NewFSPage(fsp_bp)
	fsp.InitSysExtend()
	// 设置
	fsp.ExtendInodePage(sm.getFragmentPage())
	// segment
	// 为了建立索引树先初始化一个Inode entity
	fsp_trx_bp := sm.getFlushPage(3)
	fsp_trx := page.NewFSPageTrx(fsp_trx_bp)
	fsp_trx.SetSysTrxIDStore(0)

	dict_bp := sm.getFlushPage(8)

	/*
	初始化目录页并刷新
	*/
	dirct := page.NewDictPage(dict_bp)

	// sys_tables
	dirct.SetHdrTables(sm.getFragmentPage())
	// sys_indexes
	dirct.SetHdrIndex(sm.getFragmentPage())
	// sys_fields
	dirct.SetHdrFields(sm.getFragmentPage())
	// sys_columns
	dirct.SetHdrColumns(sm.getFragmentPage())

	sm.CacheBuffer.ForceFlush(sm.cacheWrapper)
}

func (sm *TableFileManage) createSegment() {
	//fsp:=CB.GetPage()
}

//func (sm *TableFileManage) createTree(rootPage uint32) {
//	inode_bp := sm.getFlushPage(1)
//	// 创建段描述页
//	inode := page.NewINodePage(inode_bp, sm.cacheWrapper)
//	inode.FH.SetOffset(1)
//}

//
//func (sm *TableFileManage) GetPage(offset uint64)cType.PageData{
//	sm.file.Seek(int64(offset*page.PAGE_SIZE),0)
//	buf:=cType.PageData{}[:]
//	sm.file.Read(buf)
//	return utils.GetPageDate(buf)
//}
//

// 将表空间扩展至
func (sm *TableFileManage) FSPExtendFile() {
	fsp_bp := sm.getFlushPage(0)
	fsp := page.NewFSPage(fsp_bp)
	fsp.FSH.SetMaxPage(64)
	// 设定初始化成功

}

func (sm *TableFileManage) IsInitialized() bool {
	if (sm.cacheWrapper.TableID == 0) {
		dict_bp := sm.CacheBuffer.GetPage(sm.cacheWrapper, 8)
		var checkInitPage []uint32;
		dict := page.NewDictPage(dict_bp)
		column := dict.HdrColumns()
		table := dict.HdrTables()
		index := dict.HdrIndex()
		field := dict.HdrFields()
		checkInitPage = append(checkInitPage, column, table, index, field)
		for _, v := range checkInitPage {
			if v == 0 {
				return false
			}
		}
	} else {
		// user table
	}
	return true
}

func (sm *TableFileManage) CreateIndex(pageNo uint32) {
	ip := page.NewIndexPage(sm.CacheBuffer.GetPage(sm.cacheWrapper, pageNo))
	p := sm.space().FSH.FreeInodeList.GetFirst();
	l := page.NewINodePage(
		sm.CacheBuffer.GetPage(sm.cacheWrapper, p.Page()),
	)
	t := page.NewINodePage(
		sm.CacheBuffer.GetPage(sm.cacheWrapper, p.Page()),
	)
	inp_l, off_l := l.GetFreeInode()
	inp_t, off_t := t.GetFreeInode()
	ip.IH.SetLeafSegment(0, inp_l, off_l)
	ip.IH.SetTopSegment(0, inp_t, off_t)
	sm.space().AddExtendToInodeEntry(inp_l, off_l)
}

func (sm *TableFileManage) crateFSPExtend() {

}

func (sm *TableFileManage) space() *page.FSPage {
	return page.NewFSPage(sm.getFlushPage(0))
}

func (sm *TableFileManage) getFragmentPage() uint32 {
	pos := sm.space().FSH.FragFreeList.GetFirst()
	return page.GetFreePageByXdes(sm.cacheWrapper, pos.Page(), pos.Offset())
}

func (sm *TableFileManage) GetPage(pageNo uint32) *pcache.BuffPage {
	return sm.CacheBuffer.GetPage(sm.cacheWrapper, pageNo)
}

func (sm *TableFileManage) SysDir() *page.DictPage {
	dict_bp := sm.CacheBuffer.GetPage(sm.cacheWrapper, 8)
	return page.NewDictPage(dict_bp)
}
