package spaceManage

import (
	"github.com/fangker/gbdb/backend/cache"
	"github.com/fangker/gbdb/backend/log/undo"
	"github.com/fangker/gbdb/backend/tbm"
	"github.com/fangker/gbdb/backend/utils"
	"github.com/fangker/gbdb/backend/dm/page"
)

type SpaceManage struct {
	cb        *cache.CachePool
	Spaces    map[uint32]*Space
	UndoSpace *UndoSpace
}
var SM *SpaceManage
func NewSpaceManage(cb *cache.CachePool) *SpaceManage {
	 SM=&SpaceManage{cb: cb, Spaces: make(map[uint32]*Space)}
	 // 初始化space下所有cachePool
	 page.AttachCache()
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
