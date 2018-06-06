package spaceManager

import (
	"os"
	"github.com/fangker/gbdb/backend/tm"
	"github.com/fangker/gbdb/backend/cache"
)

var SM *SpaceManage

type SpaceManage struct {
	cb *cache.CachePool
	tf map[uint32]*tm.TableManager
}

func NewSpaceManage(cb *cache.CachePool) *SpaceManage {
	return &SpaceManage{cb: cb, tf: make(map[uint32]*tm.TableManager)}
}

func (sm *SpaceManage) Add(tm *tm.TableManager) *tm.TableManager {
	tm.Tfm().CacheBuffer = sm.cb
	sm.tf[tm.TableID] = tm
	return tm
}

func NewTableFileManage(filePath string, tableID uint32) *tm.TableManager {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	return tm.NewTableManager(nil, filePath,cache.Wrapper{TableID:tableID,File:file})
}

func (sm *SpaceManage) InitSysFileStructure() {
	sm.tf[0].Tfm().InitSysFile()
}

func (sm *SpaceManage) IsInitialized(i uint32) bool{
	return sm.tf[i].Tfm().IsInitialized()
}
