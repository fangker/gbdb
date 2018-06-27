package undo

import (
	"github.com/fangker/gbdb/backend/cache"
	"os"
)

type UndoFileManager struct {
	CacheBuffer *cache.CachePool
	FilePath    string
	cache.Wrapper
}

func NewUndoLogFileManage(filePath string, tableID uint32) *TableFileManage {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	tfm := &UndoFileManager{Wrapper:cache.Wrapper{tableID,file},FilePath:filePath}
	return tfm
}