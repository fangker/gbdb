package spaceManage

import (
	"github.com/fangker/gbdb/backend/cache"
	"github.com/fangker/gbdb/backend/tbm"
	"github.com/fangker/gbdb/backend/log/undo"
)

type UndoSpace struct {
	cb    *cache.CachePool
	tbm   *tbm.TableManage
	ubm   *undo.UndoLogManager
	Space uint32
}


func (sm *UndoSpace) InitSysUndoFileStructure() bool {
	return sm.ubm.Ufm().InitSysUndoFile()
}
