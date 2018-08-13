package main

import (
	"github.com/fangker/gbdb/backend/cache"
	"github.com/fangker/gbdb/backend/spaceManager"
	"github.com/fangker/gbdb/backend/utils"
	"github.com/fangker/gbdb/backend/tm"
	"github.com/fangker/gbdb/backend/log/undo"
	"github.com/fangker/gbdb/backend/cache/system"
	"github.com/fangker/gbdb/backend/tbm"
)

func main() {
	sc := loadDBSys()
	test(sc)
}

func loadDBSys() *sc.SystemCache {
	//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//	//if err != nil {
	//	//	log.Fatal(err)
	//	//}
	// 创建缓冲池子
	cb := cache.NewCacheBuffer(22)
	// 加载字典表的过程
	sm := spaceManage.NewSpaceManage(cb)
	// Undo
	undo_space := spaceManage.NewSpace(1, cb)
	undo_log := undo.NewUndoLogFileManage(utils.ENV_DIR+"/a.undo", 1)
	undo_space=sm.AddUndoSpace(undo.NewUndoLogManager(undo_log, "sys_table"))
	undo_space.InitSysUndoFileStructure()
	// Sys
	sys_tfm := tbm.NewTableFileManage(utils.ENV_DIR+"/a.db", 0)
	sys_space:=sm.AddSpace(0,tbm.NewTableManage(sys_tfm, "sys_table", 0))
	if !sys_space.IsInitialized() {
		sys_space.InitSysFileStructure()
	}
	return sys_space.LoadSysCache()
}
func test(scp *sc.SystemCache) {
	gtm:=tm.NewTransactionManage(sc.SC)
	gtm.NewTransaction();

}
