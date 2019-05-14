package main

import (
	"github.com/fangker/gbdb/backend/cache"
	"github.com/fangker/gbdb/backend/spaceManage"
	"github.com/fangker/gbdb/backend/cache/system"
	"github.com/fangker/gbdb/backend/tbm"
	"github.com/fangker/gbdb/backend/tm"
	"github.com/fangker/gbdb/backend/parser/statement"
	_ "github.com/fangker/gbdb/backend/utils/log"
	"github.com/fangker/gbdb/backend/log/undo"
	"github.com/fangker/gbdb/backend/utils"
	"github.com/fangker/gbdb/backend/dm/page"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	sc, sm := loadDBSys()
	test(sc, sm)
}

func loadDBSys() (*sc.SystemCache, *spaceManage.SpaceManage) {
	//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//	//if err != nil {
	//	//	log.Fatal(err)
	//	//}
	// 创建缓冲池子
	cb := cache.NewCacheBuffer(22)
	// 加载字典表的过程
	sm := spaceManage.NewSpaceManage(cb)
	// 初始化页全局
	page.Init()

	// Undo
	undo_log := undo.NewUndoLogFileManage(utils.ENV_DIR+"/a.undo", 1)
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

func test(sct *sc.SystemCache, smt *spaceManage.SpaceManage) {
	// 载入事物管理器
	tm.NewTransactionManage(sc.SC)
	// 创建一个新的表
	trx := tm.TM.TrxStart();
	sct.Sys_tables.Insert(trx, &statement.Insert{
		TableName: "Sys_tables",
		Fields:    []string{"name", "id", "n_cols", "type", "space"},
		Values:    []string{"students", "4", "3", "1", "3"},
	})
}
