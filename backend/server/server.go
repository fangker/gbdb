package main

import (
	"github.com/fangker/gbdb/backend/dm/cacheBuffer"
	"github.com/fangker/gbdb/backend/dm/spaceManage"
)


func main(){
	// 创建缓冲池子
	cacheBuffer.NewCacheBuffer(22)
	// 加载字典表的过程
	sm:=spaceManage.NewSpaceManage()
	tfm:=sm.Add(spaceManage.NewTableFileManage("../temData/a.db",0))

}
