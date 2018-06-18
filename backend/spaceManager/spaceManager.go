package spaceManager

import (
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
	dirct:=page.NewDictPage(dirct_bp);
	newTfm := func() *tm.TableFileManage {
		return tm.NewTableFileManage(tfm.FilePath, 0)
	}
	tables := tm.NewTableManager(newTfm(), "sys_tables",dirct.HdrTables())
	indexes := tm.NewTableManager(newTfm(), "sys_indexes",dirct.HdrIndex())
	fields := tm.NewTableManager(newTfm(), "sys_fields",dirct.HdrFields())
	columns := tm.NewTableManager(newTfm(), "sys_columns",dirct.HdrColumns())
	return cache.LoadSysCache(tables, fields, columns, indexes)
}
