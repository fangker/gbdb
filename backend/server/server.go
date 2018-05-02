package main

import (
	"github.com/fangker/gbdb/backend/cache"
	"github.com/fangker/gbdb/backend/dm/spaceManage"
)

func main(){
	// 创建缓冲池子
	cb:=cache.NewCacheBuffer(22)
	// 加载字典表的过程
	sm:=spaceManage.NewSpaceManage(cb)
	sm.Add(spaceManage.NewTableFileManage("a.db",0))
	sm.InitSysFileStructure()


}
