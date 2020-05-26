package spaceManage

import (
	"github.com/fangker/gbdb/backend/cache"
	"github.com/fangker/gbdb/backend/log/undo"
	"github.com/fangker/gbdb/backend/tbm"
	"github.com/fangker/gbdb/backend/utils"
)

type SpaceManage struct {
	cb        *cache.Pool
	Spaces    map[uint32]*Space
	UndoSpace *UndoSpace
}
var SM *SpaceManage
func NewSpaceManage(cb *cache.Pool) *SpaceManage {
	 SM=&SpaceManage{cb: cb, Spaces: make(map[uint32]*Space)}
	return SM
}

func (sm *SpaceManage) AddSpace(space uint32, tm *tbm.TableManage) *Space {
	s := &Space{cb: sm.cb}
	if (space ==0){
	tm.SetTfm(space, 0, utils.ENV_DIR+"/a.db")
	tm.Tfm().CacheBuffer = sm.cb
	s.tbm = tm
	tm.SpaceID = space
	}
	return s
}

func (sm *SpaceManage) AddUndoSpace(space uint32, ubm *undo.UndoLogManager) *UndoSpace {
	s := &UndoSpace{cb: sm.cb, ubm: ubm, Space: space}
	sm.UndoSpace = s;
	return s;
}
