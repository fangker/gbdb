package spaceManage

import (
	"github.com/fangker/gbdb/backend/cache"
	"github.com/fangker/gbdb/backend/log/undo"
	"github.com/fangker/gbdb/backend/tbm"
)

type SpaceManage struct {
	cb *cache.CachePool
	Spaces map[uint32] *Space
	UndoSpace *UndoSpace
}

func NewSpaceManage(cb *cache.CachePool) *SpaceManage {
	return  &SpaceManage{cb:cb,Spaces:make(map[uint32] *Space)}
}

func (sm *SpaceManage) AddSpace(space uint32,tm *tbm.TableManage) *Space {
	s:= &Space{cb:sm.cb}
	tm.Tfm().CacheBuffer = sm.cb
	s.tbm = tm
	return s
}

func (sm *SpaceManage) AddUndoSpace(space uint32,ubm *undo.UndoLogManager) *UndoSpace {
	s:= &UndoSpace{cb:sm.cb,ubm:ubm,Space:space}
	sm.UndoSpace = s;
	return s;
}
