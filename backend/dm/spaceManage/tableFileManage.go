package spaceManage

import (
	"os"

	"github.com/fangker/gbdb/backend/cache"
	"github.com/fangker/gbdb/backend/dm/buffPage"
	"github.com/fangker/gbdb/backend/dm/constants/cType"
	"github.com/fangker/gbdb/backend/dm/page"
	//"github.com/fangker/gbdb/backend/utils/log"
)

type tableFileManage struct {
	cacheBuffer *cache.CachePool
	filePath    string
	tableID     uint32
	file        *os.File
}

//// 初始化一个文件 设定初始页面构造文件结构
//func (sm *tableFileManage) InitFileStructure() {
//
//}
//
//// 初始化一系统表文件 设定初始页面构造文件结构
//func (sm *tableFileManage) InitSysFileStructure() {
//
//}
//
func (sm *tableFileManage) writeSync(pageNum uint32, data cType.PageData) {
	offset := pageNum * cType.PAGE_SIZE
	sm.file.WriteAt(data[:], int64(offset))
	sm.file.Sync()
}

func (sm *tableFileManage) getFreePage() *pcache.BuffPage {
	return sm.cacheBuffer.GetFreePage(sm.file)
}

func (sm *tableFileManage) getPage(pageNo uint32) *pcache.BuffPage {
	return sm.cacheBuffer.GetPage(wrapper(sm), pageNo)
}

func (sm *tableFileManage) initSysFile() {
	fsp_bp := sm.getPage(0)
	fsp_bp.Lock()
	fsp := page.NewFSPage(fsp_bp)
	fsp.InitSysExtend(wrapper(sm))
	// segment
	//fsp.FSH.
	inode_bp := sm.getPage(1)
	// 创建段描述页
	inode_bp.Lock()
	inode_bp.Dirty()
	inode := page.NewINodePage(inode_bp)
	// 为了建立索引树先初始化一个Inode entity
	sm.getFragmentPage()
	inode.SetFreeInode(1, 1)
	inode.FH.SetOffset(1)
	inode_bp.Dirty()
	// 第三个页面创建索引树
	sysIndex_bp := sm.getPage(2)
	sysIndex_bp.Lock()
	//page.NewPage(fsp_bp)
	// sys_tables
	// sys_columns
	// sys_indexes
	// sys_fields
	//sysIndexPage.GetDate()[page.FIL_HEADER_OFFSET:page.FIL_HEADER_OFFSET+8]

}

func (sm *tableFileManage) createSegment() {
	//fsp:=CB.GetPage()
}

//
//func (sm *tableFileManage) GetPage(offset uint64)cType.PageData{
//	sm.file.Seek(int64(offset*page.PAGE_SIZE),0)
//	buf:=cType.PageData{}[:]
//	sm.file.Read(buf)
//	return utils.GetPageDate(buf)
//}
//

// 将表空间扩展至
func (sm *tableFileManage) FSPExtendFile() {
	fsp_bp := sm.getPage(0)
	fsp := page.NewFSPage(fsp_bp)
	fsp.FSH.SetMaxPage(64)

	// 设定初始化成功
	fsp.FSH.SetLimitPage(64)

}

func (sm *tableFileManage) crateFSPExtend() {

}

func (sm *tableFileManage) space() *page.FSPage {
	return page.NewFSPage(sm.getPage(0))
}

func (sm *tableFileManage) getFragmentPage() {
	pageID, offset := sm.space().FSH.FragFreeList.GetFirst()
	page.GetFragFreePage(wrapper(sm), pageID, offset)
}

func wrapper(sm *tableFileManage) cache.Wrapper {
	return cache.Wrapper{File: sm.file, TableID: sm.tableID}
}
