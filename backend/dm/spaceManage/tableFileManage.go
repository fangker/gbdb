package spaceManage

import (
	"os"
	"github.com/fangker/gbdb/backend/dm/constants/cType"
	"github.com/fangker/gbdb/backend/dm/cacheBuffer"
	"github.com/fangker/gbdb/backend/dm/buffPage"
	"github.com/fangker/gbdb/backend/dm/page"
)


type tableFileManage struct {
	cacheBuffer *cacheBuffer.CacheBuffer
	filePath string
	tableID  uint32
	file     *os.File
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

func (sm *tableFileManage)getFreePage() *pcache.BuffPage{
	return sm.cacheBuffer.GetFreePage(sm.file)
}

func (sm *tableFileManage)getPage(pageNo uint64)*pcache.BuffPage{
	return sm.cacheBuffer.GetPage(sm.tableID,pageNo,sm.file)
}

func (sm *tableFileManage)initSysFile(){
	fsp_bp:=sm.getPage(0)
	fsp_bp.Lock()
	page.NewFSPage(fsp_bp)
	inode_bp:=sm.getPage(1)
	// 创建段描述页
	inode_bp.Lock()
	inode_bp.Dirty()
	inode:=page.NewINodePage(inode_bp)
	inode.FH.SetOffset(1)
	inode_bp.Dirty()

	// 第三个页面创建索引树
	sysIndex_bp:= sm.getPage(3)
	sysIndex_bp.Lock()
	page.NewPage(fsp_bp)
	// sys_tables
	// sys_columns
	// sys_indexes
	// sys_fields
	//sysIndexPage.GetDate()[page.FIL_HEADER_OFFSET:page.FIL_HEADER_OFFSET+8]





}

func (sm *tableFileManage) createSegment(){
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
//// 将表空间扩展至
//func (sm *tableFileManage) extendTo(strat uint32,end uint32){
//	//dv:=strat-end
//
//}
