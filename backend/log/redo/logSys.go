package redo

import (
	"bytes"
	"github.com/fangker/gbdb/backend/conf"
	"github.com/fangker/gbdb/backend/file"
	"path"
	"strconv"
	"sync"
)

const bufferSize = 1024 << 10

type logSys struct {
	logDir        string
	MaxAgeSync    Lsn
	GroupCapacity uint64
	logBuffer     *bytes.Buffer
	sync.Mutex
}

func NewLogSys(fsys *file.FileSys, filNum int, fSize uint) (ls *logSys) {
	logPath := conf.GetConfig().LogDirPath
	ls = &logSys{logDir: logPath, logBuffer: bytes.NewBuffer([]byte{})}
	spaceName := "undo"
	fspace := fsys.CreateFilSpace(spaceName, 1, conf.GetConfig().LogDirPath, file.TypeRedo, uint64(fSize))
	for i := 0; i < filNum; i++ {
		fspace.CreateFilUnit(path.Join(logPath, spaceName+"_"+strconv.Itoa(i)+fspace.FType.FileType()), uint64(fSize))
	}
	return
}
func (ls *logSys) WriteToLog(bytes []byte) {
	ls.logBuffer.Write(bytes)
}

func (ls *logSys) LogRemain() int {
	return bufferSize - ls.logBuffer.Len()
}
