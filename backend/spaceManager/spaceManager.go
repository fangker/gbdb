package spaceManager

import (
	"github.com/fangker/gbdb/backend/tbm"
	"github.com/fangker/gbdb/backend/cache"
	"github.com/fangker/gbdb/backend/dm/page"
	"github.com/fangker/gbdb/backend/log/undo"
	"github.com/fangker/gbdb/backend/cache/system"
)

var SM *SpaceManage

type SpaceManage struct {
	cb    *cache.CachePool
	tf    *tbm.TableManage
	uf    *undo.UndoLogManager
	Space uint32
}

func NewSpaceManage(space uint32, cb *cache.CachePool) *SpaceManage {
	return &SpaceManage{cb: cb, Space: space}
}

func (sm *SpaceManage) Add(tm *tbm.TableManage) *tbm.TableManage {
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

func (sm *SpaceManage) GetTf() *tbm.TableManage {
	return sm.tf
}

func (sm *SpaceManage) LoadSysCache() *sc.SystemCache {
	sm.tf.TableID = 0
	sys := sm.tf
	tfm := sys.Tfm()
	dirct_bp := sm.cb.GetPage(sys.Wrapper(), 8)
	dirct:=page.NewDictPage(dirct_bp);
	newTfm := func() *tbm.TableFileManage {
		return tbm.NewTableFileManage(tfm.FilePath, 0)
	}
	tables := tbm.NewTableManager(newTfm(), "sys_tables",dirct.HdrTables())
	indexes := tbm.NewTableManager(newTfm(), "sys_indexes",dirct.HdrIndex())
	fields := tbm.NewTableManager(newTfm(), "sys_fields",dirct.HdrFields())
	columns := tbm.NewTableManager(newTfm(), "sys_columns",dirct.HdrColumns())
	return sc.LoadSysCache(tables, fields, columns, indexes)
}

func (sm *SpaceManage) InitSysUndoFileStructure() bool{
	return sm.uf.Ufm().InitSysUndoFile()
}
