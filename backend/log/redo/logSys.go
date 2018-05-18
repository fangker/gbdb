package redo

import "io/ioutil"

const (
	REDO_LOG_PATH  = ""
	REDO_LOG_SIZE  = 1024 << 20
	REDO_LOG_GROUP = 2
)

type logSys struct {
	logGroup      *logGroup
	MaxAgeSync    Lsn
	GroupCapacity uint64
}

func NewLogSys(path string) {
	this := &logSys{}
	this.logGroup = &logGroup{nFiles: REDO_LOG_GROUP}
	dir, err := ioutil.ReadDir(REDO_LOG_PATH)
	if err != nil {
		panic(err)
	}
	for _, fi := range dir {
		if fi.IsDir() {
			continue
		}

	}
}
