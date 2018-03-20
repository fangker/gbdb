package spaceManage

import (
	"os"
	//"sync"
	"github.com/fangker/gbdb/backend/dm/page"
	"github.com/fangker/gbdb/backend/dm/cacheBuffer"
	"github.com/fangker/gbdb/backend/utils"
)
var CB *cacheBuffer.CacheBuffer  = cacheBuffer.CB

type tableFileManage struct {
	filePath string
	tableID  uint32
	file     *os.File
	//rwLock   *sync.RWMutex
	//mLock    *sync.Mutex
}

// 初始化一个文件 设定初始页面构造文件结构
func (sm *tableFileManage) InitFileStructure() {

}

// 初始化一系统表文件 设定初始页面构造文件结构
func (sm *tableFileManage) InitSysFileStructure() {

}

func (sm *tableFileManage) writeSync(pageNum uint32, data page.PageData) {
	offset := pageNum * page.PAGE_SIZE
	sm.file.WriteAt(data[:], int64(offset))
	sm.file.Sync()
}

func (sm *tableFileManage)initSysFile(){
	fsp:=CB.GetFreePage()
	fsp.WLock()
	fsp.SetType(page.PAGE_TYPE_FSP)
	fsp.Page.FH.SetSpace(0)


	inode:=CB.GetFreePage()
	// 创建段描述页
	inode.SetType(page.PAGE_TYPE_INODE)
	inode.WLock()
	// 第三个页面创建索引树
	sysIndexPage:= CB.GetFreePage()
	sysIndexPage.WLock()
	sysIndexPage.SetType(page.PAGE_TYPE_PAGE)
	sysIndexPage.Page.FH.SetSpace(0)
	sysIndexPage.Page.FH.SetOffset(3)
	// sys_tables
	// sys_columns
	// sys_indexes
	// sys_fields
	//sysIndexPage.GetDate()[page.FIL_HEADER_OFFSET:page.FIL_HEADER_OFFSET+8]





}

func (sm *tableFileManage) createSegment(){
	//fsp:=CB.GetPage()
}

func (sm *tableFileManage) GetPage(offset uint32)page.PageData{
	sm.file.Seek(int64(offset*page.PAGE_SIZE),0)
	buf:=page.PageData{}[:]
	sm.file.Read(buf)
	return utils.GetPageDate(buf)
}

// 将表空间扩展至
func (sm *tableFileManage) extendTo(strat uint32,end uint32){
	//dv:=strat-end

}
