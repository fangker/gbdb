package spaceManage

import (
	"os"
	"fmt"
	"github.com/fangker/gbdb/backend/dm/cacheBuffer"
)


var SM *SpaceManage

type SpaceManage struct {
	cb  *cacheBuffer.CacheBuffer
	tf  map[uint32] *tableFileManage
}


func NewSpaceManage(cb *cacheBuffer.CacheBuffer)*SpaceManage{
	return &SpaceManage{cb:cb,tf:make(map[uint32] *tableFileManage)}
}

func (sm *SpaceManage)Add(tfm *tableFileManage) *tableFileManage{
	tfm.cacheBuffer=sm.cb
	sm.tf[tfm.tableID]=tfm
	return tfm
}

func NewTableFileManage(filePath string, tableID uint32) *tableFileManage {
	file, err := os.OpenFile(filePath,os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	return &tableFileManage{nil,filePath, tableID, file,}
}

func (sm *SpaceManage) InitSysFileStructure()  {
    sm.tf[0].initSysFile()
}

