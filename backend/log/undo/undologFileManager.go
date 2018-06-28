package undo

import (
	"github.com/fangker/gbdb/backend/cache"
	"os"
	"github.com/fangker/gbdb/backend/dm/page"
	"github.com/fangker/gbdb/backend/dm/buffPage"
)

type UndoFileManager struct {
	CacheBuffer *cache.CachePool
	FilePath    string
	cache.Wrapper
}

func NewUndoLogFileManage(filePath string, tableID uint32) *UndoFileManager {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	tfm := &UndoFileManager{Wrapper:cache.Wrapper{tableID,file},FilePath:filePath}
	return tfm
}

func (this *UndoFileManager) getFlushPage(pageNo uint32) *pcache.BuffPage {
	return this.CacheBuffer.GetFlushPage(this.wrapper(), pageNo)
}

func (this *UndoFileManager)IsInitialized() bool {
	fsp_bp := this.getFlushPage(0)
	fsp_bp.Lock()
	fsp := page.NewFSPage(fsp_bp)
	fsp.InitSysExtend()

}
func (sm *UndoFileManager) wrapper() cache.Wrapper {
	return cache.Wrapper{sm.TableID, sm.File}
}
