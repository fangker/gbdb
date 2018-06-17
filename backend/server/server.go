package main

import (
	"github.com/fangker/gbdb/backend/cache"
	"github.com/fangker/gbdb/backend/spaceManager"
	"github.com/fangker/gbdb/backend/utils"
	"github.com/fangker/gbdb/backend/tm"
)

func main() {
	loadDBSys()
}

func loadDBSys() {
	//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//	//if err != nil {
	//	//	log.Fatal(err)
	//	//}
	// 创建缓冲池子
	cb := cache.NewCacheBuffer(22)
	// 加载字典表的过程
	sm := spaceManager.NewSpaceManage(0,cb)
	sys_tfm:=tm.NewTableFileManage(utils.ENV_DIR+"/a.db", 0)
	sm.Add(tm.NewTableManager(sys_tfm,"sys_table"))
	if !sm.IsInitialized(0) {
		sm.InitSysFileStructure()
	}
	sm.LoadSysCache();
}
