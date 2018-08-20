package redo

import (
	"os"

	"github.com/fangker/gbdb/backend/dm/constants/cType"
	"math"
	"github.com/fangker/gbdb/backend/utils"
)

type (
	Lsn cType.LSN
)

type logGroup struct {
	nFiles  int
	file    []*os.File
	pointOn int
}

type fileHeader struct {
	fileStartLSN Lsn
	logFileNum   uint32
}

type fileBlock struct {
	blockNum          uint32
	blockData         uint32
	firstRecordOffset uint
}

func (this *logGroup) StartLSN() Lsn {
	data := this.getBlock(1)
	return Lsn(utils.GetUint64(data[0:8]))
}

func (this *logGroup) StartLogFileNum() uint32 {
	data := this.getBlock(1)
	return utils.GetUint32(data[7:11])
}

func (this *logGroup) CheckPoint1() Lsn {
	data := this.getBlock(1)
	return Lsn(utils.GetUint64(data[0:8]))
}

func (this *logGroup) CheckPoint2() Lsn {
	data := this.getBlock(1)
	return Lsn(utils.GetUint32(data[0:8]))
}

func (this *logGroup) SetCheckPoint1(lsn Lsn) {
	blockData := this.getBlock(2)
	copy(blockData[0:8], utils.PutUint64(uint64(lsn)))
	this.setBlock(2, blockData)
}

func (this *logGroup) SetCheckPoint2(lsn Lsn) {
	blockData := this.getBlock(4)
	copy(blockData[0:8], utils.PutUint64(uint64(lsn)))
	this.setBlock(4, blockData)
}

func (this *logGroup) SetStartLSN(lsn Lsn) {
	blockData := this.getBlock(1)
	copy(blockData[0:8], utils.PutUint64(uint64(lsn)))
	this.setBlock(1, blockData)
}

func (this *logGroup) SetStartLogFileNum(num uint32) {
	blockData := this.getBlock(1)
	copy(blockData[7:11], utils.PutUint32(num))
	this.setBlock(1, blockData)
}

func (this *logGroup) getBlock(blockNo uint32) cType.RedoData {
	var blockData cType.RedoData
	fileNo := int(math.Ceil(float64((blockNo * 512) / REDO_LOG_SIZE)))
	file := this.file[fileNo];
	seekOffset := int64(((blockNo - 1) * 512) % REDO_LOG_SIZE)
	file.Seek(seekOffset, 0)
	file.Read(blockData[:])
	return blockData
}

func (this *logGroup) setBlock(blockNo uint32, rd cType.RedoData) {
	fileNo := int(math.Ceil(float64((blockNo * 512) / REDO_LOG_SIZE)))
	file := this.file[fileNo];
	seekOffset := int64(((blockNo - 1) * 512) % REDO_LOG_SIZE)
	file.WriteAt(rd[:], seekOffset)
	file.Sync()
}
