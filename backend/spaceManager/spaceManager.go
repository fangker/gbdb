package spaceManager

import (
	"os"
	"github.com/fangker/gbdb/backend/tm"
	"github.com/fangker/gbdb/backend/cache"
	"github.com/fangker/gbdb/backend/dm/page"
)

var SM *SpaceManage

type SpaceManage struct {
	cb    *cache.CachePool
	tf    map[uint32]*tm.TableManager
	Space uint32
}

func NewSpaceManage(space uint32, cb *cache.CachePool) *SpaceManage {
	return &SpaceManage{cb: cb, tf: make(map[uint32]*tm.TableManager), Space: space}
}

func (sm *SpaceManage) Add(tm *tm.TableManager) *tm.TableManager {
	tm.Tfm().CacheBuffer = sm.cb
	sm.tf[tm.TableID] = tm
	return tm
}

func (sm *SpaceManage) InitSysFileStructure() {
	sm.tf[0].Tfm().InitSysFile()
}

func (sm *SpaceManage) IsInitialized(i uint32) bool {
	return sm.tf[i].Tfm().IsInitialized()
}

func (sm *SpaceManage) GetTf(i uint32) *tm.TableManager {
	return sm.tf[i]
}

func (sm *SpaceManage) LoadSysCache() *cache.SystemCache {
	sm.tf[0].TableID = 0
	sys := sm.tf[0]
	tfm := sys.Tfm()
	dirct_bp := sm.cb.GetPage(sys.Wrapper(), 8)
	dirct := page.NewDictPage(dirct_bp)
	tables := &tm.TableManager{TableID:0,TableName:"sys_tables",:tfm}
	indexes := &tm.TableManager{0, "sys_indexes", tfm}
	fields := &tm.TableManager{0, "sys_tables", tfm}
	columns := &tm.TableManager{0, "sys_tables", tfm}

	return &cache.SystemCache{}
}
