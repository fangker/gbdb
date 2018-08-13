package spaceManage

import (
	"github.com/fangker/gbdb/backend/tbm"
	"github.com/fangker/gbdb/backend/cache"
	"github.com/fangker/gbdb/backend/dm/page"
	"github.com/fangker/gbdb/backend/log/undo"
	"github.com/fangker/gbdb/backend/cache/system"
)

var SM *Space

type Space struct {
	cb    *cache.CachePool
	tf    *tbm.TableManage
	uf    *undo.UndoLogManager
	Space uint32
}

func NewSpace(space uint32, cb *cache.CachePool) *Space {
	return &Space{cb: cb, Space: space}
}



func (sm *Space) InitSysFileStructure() {
	sm.tf.Tfm().InitSysFile()
}

func (sm *Space) IsInitialized() bool {
	return sm.tf.Tfm().IsInitialized()
}

func (sm *Space) GetTf() *tbm.TableManage {
	return sm.tf
}

func (sm *Space) LoadSysCache() *sc.SystemCache {
	sm.tf.TableID = 0
	sys := sm.tf
	tfm := sys.Tfm()
	dirct_bp := sm.cb.GetPage(sys.Wrapper(), 8)
	dirct:=page.NewDictPage(dirct_bp);
	newTfm := func() *tbm.TableFileManage {
		return tbm.NewTableFileManage(tfm.FilePath, 0)
	}
	tables := tbm.NewTableManage(newTfm(), "sys_tables",dirct.HdrTables())
	indexes := tbm.NewTableManage(newTfm(), "sys_indexes",dirct.HdrIndex())
	fields := tbm.NewTableManage(newTfm(), "sys_fields",dirct.HdrFields())
	columns := tbm.NewTableManage(newTfm(), "sys_columns",dirct.HdrColumns())
	return sc.LoadSysCache(tables, fields, columns, indexes)
}

func (sm *Space) InitSysUndoFileStructure() bool{
	return sm.uf.Ufm().InitSysUndoFile()
}
