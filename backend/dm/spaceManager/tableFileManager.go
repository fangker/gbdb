package spaceManager

import (
	"github.com/fangker/gbdb/backend/cache"
	"github.com/fangker/gbdb/backend/dm/buffPage"
	"github.com/fangker/gbdb/backend/dm/constants/cType"
	"github.com/fangker/gbdb/backend/dm/page"
	"github.com/fangker/gbdb/backend/utils/log"
	"fmt"
)

type tableFileManage struct {
	cacheBuffer *cache.CachePool
	filePath    string
	cache.Wrapper
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
	sm.File.WriteAt(data[:], int64(offset))
	sm.File.Sync()
}

func (sm *tableFileManage) getFlushPage(pageNo uint32) *pcache.BuffPage {
	return sm.cacheBuffer.GetFlushPage(sm.wrapper(), pageNo)
}

func (sm *tableFileManage) initSysFile() {
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
	dict_bp := sm.getFlushPage(8)
	dirct := page.NewDictPage(dict_bp)
	log.Error(sm.getFragmentPage())
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
	sm.cacheBuffer.ForceFlush(sm.wrapper())
	//page.NewPage(fsp_bp)
	//sm.cacheBuffer.

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
	fsp_bp := sm.getFlushPage(0)
	fsp := page.NewFSPage(fsp_bp)
	fsp.FSH.SetMaxPage(64)

	// 设定初始化成功
	fsp.FSH.SetLimitPage(64)

}

func (sm *tableFileManage) IsInitialized() bool {
	if (sm.TableID == 0) {
		dict_bp := sm.cacheBuffer.GetPage(sm.wrapper(), 8)
		var checkInitPage []uint32;
		dict := page.NewDictPage(dict_bp)
		column := dict.HdrColumns()
		table := dict.HdrTables()
		index := dict.HdrIndex()
		field := dict.HdrFields()
		checkInitPage = append(checkInitPage, column, table, index, field)
		for i,v := range checkInitPage {
			if v == 0 {
				fmt.Println(i,v)
				return false
			}
		}
	}else{
		// user table
	}
	return true
}

func (sm *tableFileManage) crateFSPExtend() {

}

func (sm *tableFileManage) space() *page.FSPage {
	return page.NewFSPage(sm.getFlushPage(0))
}

func (sm *tableFileManage) getFragmentPage() uint32 {
	pageID, offset := sm.space().FSH.FragFreeList.GetFirst()
	return page.GetFragFreePage(sm.wrapper(), pageID, offset)
}

func (sm *tableFileManage) wrapper() cache.Wrapper {
	return cache.Wrapper{sm.TableID, sm.File}
}

