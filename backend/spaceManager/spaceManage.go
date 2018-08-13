package spaceManage

import (
	"github.com/fangker/gbdb/backend/cache"
	"github.com/fangker/gbdb/backend/log/undo"
	"github.com/fangker/gbdb/backend/tbm"
)

type SpaceManage struct {
	cb *cache.CachePool
}

func NewSpaceManage(cb *cache.CachePool) *SpaceManage {
	return  &SpaceManage{cb}
}

func (sm *SpaceManage) AddSpace(space uint32,tm *tbm.TableManage) *Space {
	s:= &Space{cb:sm.cb}
	tm.Tfm().CacheBuffer = sm.cb
	s.tf = tm
	return s
}

func (sm *SpaceManage) AddUndoSpace(uf *undo.UndoLogManager) *Space {
	s:= &Space{cb:sm.cb,uf:uf}
	return s;
}
