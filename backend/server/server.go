package main

import (
	"github.com/fangker/gbdb/backend/cache"
	"github.com/fangker/gbdb/backend/dm/spaceManage"
	"path/filepath"
	"os"
	"log"
)

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	// 创建缓冲池子
	cb := cache.NewCacheBuffer(22)
	// 加载字典表的过程
	sm := spaceManage.NewSpaceManage(cb)
	sm.Add(spaceManage.NewTableFileManage( dir+"a.db", 0))
	sm.InitSysFileStructure()

}
