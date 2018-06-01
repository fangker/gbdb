package spaceManager

import (
	"os"

	"github.com/fangker/gbdb/backend/cache"
)

var SM *SpaceManage

type SpaceManage struct {
	cb *cache.CachePool
	tf map[uint32]*tableFileManage
}

func NewSpaceManage(cb *cache.CachePool) *SpaceManage {
	return &SpaceManage{cb: cb, tf: make(map[uint32]*tableFileManage)}
}

func (sm *SpaceManage) Add(tfm *tableFileManage) *tableFileManage {
	tfm.cacheBuffer = sm.cb
	sm.tf[tfm.TableID] = tfm
	return tfm
}

func NewTableFileManage(filePath string, tableID uint32) *tableFileManage {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	return &tableFileManage{nil, filePath,cache.Wrapper{TableID:tableID,File:file}}
}

func (sm *SpaceManage) InitSysFileStructure() {
	sm.tf[0].initSysFile()
}

func (sm *SpaceManage) IsInitialized(i uint32) bool{
	return sm.tf[i].IsInitialized()
}
