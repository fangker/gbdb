package redo

import (
	"os"

	"github.com/fangker/gbdb/backend/dm/constants/cType"
	"math"
)

type (
	Lsn cType.Lsn
)

type logGroup struct {
	nFiles  int
	file    []*os.File
	pointOn int
}

type fileHeader struct {
	fileStartLSN Lsn
	logFileNum  uint32
}

type fileBlock struct {
	blockNum uint32
	blockData  uint32
	firstRecordOffset uint
}

func (this *logGroup) GetStartLSN() {
	return
}

func (this *logGroup) getBlock(blockNo uint32) cType.RedoDate {
	var blockData cType.RedoDate
	fileNo := int(math.Ceil(float64((blockNo * 512) / REDO_LOG_SIZE)))
	file := this.file[fileNo];
	seekOffset := int64(((blockNo-1) * 512) % REDO_LOG_SIZE)
	file.Seek(seekOffset, 0)
	file.Read(blockData[:])
	return blockData
}
