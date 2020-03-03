package srv

import (
	"github.com/fangker/gbdb/backend/cache"
	"github.com/fangker/gbdb/backend/spaceManage"
	"github.com/fangker/gbdb/backend/cache/system"
	"github.com/fangker/gbdb/backend/tbm"
	_ "github.com/fangker/gbdb/backend/utils/ulog"
	"github.com/fangker/gbdb/backend/log/undo"
	"github.com/fangker/gbdb/backend/dm/page"
	"runtime"
	"github.com/fangker/gbdb/backend/conf"
)

func main() {
	runtime.GOMAXPROCS(3)
	loadDBSys()
}

func loadDBSys() (*sc.SystemCache, *spaceManage.SpaceManage) {
	serverStartConfig := conf.GetServerStartConfig()
	// 创建缓冲池子
	cb := cache.NewCacheBuffer(serverStartConfig.BufferPageMemory * 1024 / 16)
	// 加载字典表的过程
	sm := spaceManage.NewSpaceManage(cb)
	// 初始化页全局
	page.Init()

	// Undo
	undo_log := undo.NewUndoLogFileManage(serverStartConfig.DbDirPath+"/a.undo", 1)
	undo_space := sm.AddUndoSpace(1, undo.NewUndoLogManager(undo_log, "sys_table"))
	undo_space.InitSysUndoFileStructure()
	// Sys
	sys_space := sm.AddSpace(0, tbm.NewTableManage("sys_table"))
	isInitSys := sys_space.IsInitialized()
	if isInitSys {
		/*
		初始化Trx页面(4)
		初始化FreeFrag(1)
		初始化首个Inode页面
		初始化Dir页面(8)
		*/
		sys_space.InitSysFileStructure()
	}
	return sys_space.LoadSysCache(!isInitSys), sm
}

//func test(sct *sc.SystemCache, smt *spaceManage.SpaceManage) {
//	// 载入事物管理器
//	tm.NewTransactionManage(sc.SC)
//	// 创建一个新的表
//	trx := tm.TM.TrxStart();
//	sct.Sys_tables.Insert(trx, &statement.Insert{
//		TableName: "Sys_tables",
//		Fields:    []string{"name", "id", "n_cols", "type", "space"},
//		Values:    []string{"students", "4", "3", "1", "3"},
//	})
//}
