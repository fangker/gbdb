package spaceManage

import (
	"github.com/fangker/gbdb/backend/tbm"
	"github.com/fangker/gbdb/backend/cache"
	"github.com/fangker/gbdb/backend/dm/page"
	"github.com/fangker/gbdb/backend/log/undo"
	"github.com/fangker/gbdb/backend/cache/system"
	"github.com/fangker/gbdb/backend/tbm/tfm"
)

var S *Space

type Space struct {
	cb    *cache.Pool
	tbm   *tbm.TableManage
	ubm   *undo.UndoLogManager
	Space uint32
}

func NewSpace(space uint32, cb *cache.Pool) *Space {
	return &Space{cb: cb, Space: space}
}

func (sm *Space) InitSysFileStructure() {
	sm.tbm.Tfm().InitSysFile()
}

func (sm *Space) IsInitialized() bool {
	return sm.tbm.Tfm().IsInitialized()
}

func (sm *Space) GetTf() *tbm.TableManage {
	return sm.tbm
}

// 载入sys缓存
func (sm *Space) LoadSysCache(init bool) *sc.SystemCache {
	sm.tbm.TableID = 0
	sys := sm.tbm
	stfm := sys.Tfm()
	cacheWrapper:=sm.tbm.Tfm().CacheWrapper()
	dirct_bp := sm.cb.GetPage(cacheWrapper, 8)
	dirct := page.NewDictPage(dirct_bp);
	var tem_id = 0;
	newTfm := func() *tfm.TableFileManage {
		tbm := tfm.NewTableFileManage(0,uint32(tem_id),stfm.FilePath)
		tem_id++
		return tbm;
	}
	if init {
		tfm:= newTfm()
		tbm.CreateTree(tfm,dirct.HdrTables())
		tbm.CreateTree(tfm,dirct.HdrIndex())
		tbm.CreateTree(tfm,dirct.HdrFields())
		tbm.CreateTree(tfm,dirct.HdrColumns())
	}
	tables := tbm.LoadTableManage("sys_tables", newTfm(),dirct.HdrTables())
	indexes := tbm.LoadTableManage("sys_indexes", newTfm(),dirct.HdrIndex())
	fields := tbm.LoadTableManage("sys_fields", newTfm(),dirct.HdrFields())
	columns := tbm.LoadTableManage("sys_columns", newTfm(),dirct.HdrColumns())

	tables.LoadTfm(newTfm())
	SM.AddSpace(0,tables)
	tables.LoadTfm(newTfm())
	SM.AddSpace(0,indexes)
	tables.LoadTfm(newTfm())
	SM.AddSpace(0,fields)
	tables.LoadTfm(newTfm())
	SM.AddSpace(0,columns)

	sct := sc.LoadSysCache(tables, fields, columns, indexes)
	sct.LoadSysTuple()
	return sct;
}
