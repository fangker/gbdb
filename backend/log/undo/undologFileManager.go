package undo

import (
	"github.com/fangker/gbdb/backend/cache"
	"os"
	"github.com/fangker/gbdb/backend/cache/buffPage"
	"github.com/fangker/gbdb/backend/wrapper"
	"github.com/fangker/gbdb/backend/dm/page"
)

const UNDO_SPACE = 0;

type UndoFileManager struct {
	CacheBuffer *cache.CachePool
	FilePath    string
	wp.Wrapper
}

func NewUndoLogFileManage(filePath string, tableID uint32) *UndoFileManager {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	tfm := &UndoFileManager{CacheBuffer: cache.CP, Wrapper: wp.Wrapper{UNDO_SPACE, tableID, file}, FilePath: filePath}
	return tfm
}

func (this *UndoFileManager) getFlushPage(pageNo uint32) *pcache.BuffPage {
	return this.CacheBuffer.GetFlushPage(this.wrapper(), pageNo)
}

func (this *UndoFileManager) InitSysUndoFile() bool {
	fsp_bp := this.getFlushPage(0)
	fsp_bp.Lock()
	fsp := page.NewFSPage(fsp_bp)
	fsp.InitSysExtend()
	this.CacheBuffer.ForceFlush(this.wrapper())
	return true
}

func (this *UndoFileManager) wrapper() wp.Wrapper {
	return wp.Wrapper{UNDO_SPACE, this.TableID, this.File}
}
