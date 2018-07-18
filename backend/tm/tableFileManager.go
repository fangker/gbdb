package tm

import (
	"github.com/fangker/gbdb/backend/cache"
	"github.com/fangker/gbdb/backend/dm/buffPage"
	"github.com/fangker/gbdb/backend/dm/constants/cType"
	"github.com/fangker/gbdb/backend/dm/page"
	"os"
)


type TableFileManage struct {
	CacheBuffer *cache.CachePool
	FilePath    string
	cache.Wrapper
}

type TableFileManager interface {
	SysDir() *page.DictPage
}

func NewTableFileManage(filePath string, tableID uint32) *TableFileManage {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	tfm := &TableFileManage{Wrapper:cache.Wrapper{tableID,file},FilePath:filePath,CacheBuffer:cache.CB}
	return tfm
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

func (sm *TableFileManage) writeSync(pageNum uint32, data cType.PageData) {
	offset := pageNum * cType.PAGE_SIZE
	sm.File.WriteAt(data[:], int64(offset))
	sm.File.Sync()
}

func (sm *TableFileManage) getFlushPage(pageNo uint32) *pcache.BuffPage {
	return sm.CacheBuffer.GetFlushPage(sm.wrapper(), pageNo)
}

func (sm *TableFileManage) InitSysFile() {
	fsp_bp := sm.getFlushPage(0)
	fsp_bp.Lock()
	fsp := page.NewFSPage(fsp_bp)
	fsp.InitSysExtend()
	// segment
	//fsp.FSH.
	// 为了建立索引树先初始化一个Inode entity
	inode_bp := sm.getFlushPage(1)
	// 创建段描述页
	inode_bp.Lock()
	inode_bp.Dirty()
	inode := page.NewINodePage(inode_bp)
	fsp_trx_bp:= sm.getFlushPage(3)
	fsp_trx:= page.NewFSPageTrx(fsp_trx_bp)
	fsp_trx.SetSysTrxIDStore(0)

	dict_bp := sm.getFlushPage(8)
	dirct := page.NewDictPage(dict_bp)
	// sys_tables
	dirct.SetHdrTables(sm.getFragmentPage())
	inode.SetFreeInode(sm.getFragmentPage(), sm.wrapper())
	// sys_indexes
	dirct.SetHdrIndex(sm.getFragmentPage())
	inode.SetFreeInode(sm.getFragmentPage(), sm.wrapper())
	// sys_fields
	dirct.SetHdrFields(sm.getFragmentPage())
	inode.SetFreeInode(sm.getFragmentPage(), sm.wrapper())
	// sys_columns
	dirct.SetHdrColumns(sm.getFragmentPage())
	inode.SetFreeInode(sm.getFragmentPage(), sm.wrapper())
	inode.FH.SetOffset(1)
	inode_bp.Dirty()
	// 第三个页面创建索引树
	sysIndex_bp := sm.getFlushPage(2)
	sysIndex_bp.Lock()
	sm.CacheBuffer.ForceFlush(sm.wrapper())
}

func (sm *TableFileManage) createSegment() {
	//fsp:=CB.GetPage()
}

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
	fsp.FSH.SetLimitPage(64)

}

func (sm *TableFileManage) IsInitialized() bool {
	if (sm.TableID == 0) {
		dict_bp := sm.CacheBuffer.GetPage(sm.wrapper(), 8)
		var checkInitPage []uint32;
		dict := page.NewDictPage(dict_bp)
		column := dict.HdrColumns()
		table := dict.HdrTables()
		index := dict.HdrIndex()
		field := dict.HdrFields()
		checkInitPage = append(checkInitPage, column, table, index, field)
		for _,v := range checkInitPage {
			if v == 0 {
				return false
			}
		}
	}else{
		// user table
	}
	return true
}

func (sm *TableFileManage) CreateIndex(){

}


func (sm *TableFileManage) crateFSPExtend() {

}

func (sm *TableFileManage) space() *page.FSPage {
	return page.NewFSPage(sm.getFlushPage(0))
}

func (sm *TableFileManage) getFragmentPage() uint32 {
	pageID, offset := sm.space().FSH.FragFreeList.GetFirst()
	return page.GetFragFreePage(sm.wrapper(), pageID, offset)
}

func (sm *TableFileManage) SysDir() *page.DictPage {
	dict_bp := sm.CacheBuffer.GetPage(sm.wrapper(), 8)
	return page.NewDictPage(dict_bp)
}

func (sm *TableFileManage) wrapper() cache.Wrapper {
	return cache.Wrapper{sm.TableID, sm.File}
}

