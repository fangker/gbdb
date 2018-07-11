package spaceManager

import (
	"github.com/fangker/gbdb/backend/tm"
	"github.com/fangker/gbdb/backend/cache"
	"github.com/fangker/gbdb/backend/dm/page"
	"github.com/fangker/gbdb/backend/log/undo"
	"github.com/fangker/gbdb/backend/cache/system"
)

var SM *SpaceManage

type SpaceManage struct {
	cb    *cache.CachePool
	tf    *tm.TableManager
	uf    *undo.UndoLogManager
	Space uint32
}

func NewSpaceManage(space uint32, cb *cache.CachePool) *SpaceManage {
	return &SpaceManage{cb: cb, Space: space}
}

func (sm *SpaceManage) Add(tm *tm.TableManager) *tm.TableManager {
	tm.Tfm().CacheBuffer = sm.cb
	sm.tf = tm
	return tm
}

func (sm *SpaceManage) AddUndoLog(uf *undo.UndoLogManager) * undo.UndoLogManager {
	sm.uf=uf
	return uf
}

func (sm *SpaceManage) InitSysFileStructure() {
	sm.tf.Tfm().InitSysFile()
}

func (sm *SpaceManage) IsInitialized() bool {
	return sm.tf.Tfm().IsInitialized()
}

func (sm *SpaceManage) GetTf() *tm.TableManager {
	return sm.tf
}

func (sm *SpaceManage) LoadSysCache() *systemCache.SystemCache {
	sm.tf.TableID = 0
	sys := sm.tf
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
	return systemCache.LoadSysCache(tables, fields, columns, indexes)
}

func (sm *SpaceManage) InitSysUndoFileStructure() bool{
	return sm.uf.Ufm().InitSysUndoFile()
}
