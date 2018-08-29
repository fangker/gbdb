package undo

import (
	"github.com/fangker/gbdb/backend/cache"
	"os"
	"github.com/fangker/gbdb/backend/dm/page"
	"github.com/fangker/gbdb/backend/cache/buffPage"
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
	tfm := &UndoFileManager{CacheBuffer:cache.CB,Wrapper:cache.Wrapper{tableID,file},FilePath:filePath}
	return tfm
}

func (this *UndoFileManager) getFlushPage(pageNo uint32) *pcache.BuffPage {
	return this.CacheBuffer.GetFlushPage(this.wrapper(), pageNo)
}

func (this *UndoFileManager)InitSysUndoFile() bool {
	fsp_bp := this.getFlushPage(0)
	fsp_bp.Lock()
	fsp := page.NewFSPage(fsp_bp)
	fsp.InitSysExtend()
	this.CacheBuffer.ForceFlush(this.wrapper())
	return true
}

func (this *UndoFileManager) wrapper() cache.Wrapper {
	return cache.Wrapper{this.TableID, this.File}
}
