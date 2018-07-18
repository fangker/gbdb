package main

import (
	"github.com/fangker/gbdb/backend/cache"
	"github.com/fangker/gbdb/backend/spaceManager"
	"github.com/fangker/gbdb/backend/utils"
	"github.com/fangker/gbdb/backend/tm"
	"github.com/fangker/gbdb/backend/log/undo"
	"github.com/fangker/gbdb/backend/cache/system"
	"fmt"
)

func main() {
	sc := loadDBSys()
	test(sc)
}

func loadDBSys() *systemCache.SystemCache {
	//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//	//if err != nil {
	//	//	log.Fatal(err)
	//	//}
	// 创建缓冲池子
	cb := cache.NewCacheBuffer(22)
	// 加载字典表的过程
	sm := spaceManager.NewSpaceManage(0, cb)
	// Undo
	undo_sm := spaceManager.NewSpaceManage(1, cb)
	undo_log := undo.NewUndoLogFileManage(utils.ENV_DIR+"/a.undo", 1)
	undo_sm.AddUndoLog(undo.NewUndoLogManager(undo_log, "sys_table"))
	undo_sm.InitSysUndoFileStructure()
	// Sys
	sys_tfm := tm.NewTableFileManage(utils.ENV_DIR+"/a.db", 0)
	sm.Add(tm.NewTableManager(sys_tfm, "sys_table", 0))
	if !sm.IsInitialized() {
		sm.InitSysFileStructure()
	}
	return sm.LoadSysCache()
}
func test(sc *systemCache.SystemCache) {
	sc.SysTrxIDStore().HdrTableID()=1
	fmt.Println(sc.SysTrxIDStore().HdrTableID())
}
