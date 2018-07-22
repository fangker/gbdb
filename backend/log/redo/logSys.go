package redo

import (
	"io/ioutil"
	"os"
	"strconv"
	"path"
	"fmt"
)

const (
	REDO_LOG_NAME  = "redo"
	REDO_LOG_PATH  = ""
	REDO_LOG_SIZE  = 1024 << 20
	REDO_LOG_GROUP = 2
)

type logSys struct {
	logDir        string
	logGroup      *logGroup
	MaxAgeSync    Lsn
	GroupCapacity uint64
}

func NewLogSys(fileDir string) {
	this := &logSys{logDir: fileDir}
	this.logGroup = &logGroup{nFiles: REDO_LOG_GROUP}
	dir, err := ioutil.ReadDir(this.logDir)
	if os.IsNotExist(err) {
		err := os.Mkdir(this.logDir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	exist := false
	if len(dir) == REDO_LOG_GROUP {
		for _, fi := range dir {
			if fi.IsDir() {
				continue
			}
			if (fi.Size() != REDO_LOG_SIZE) {
				exist = false
				break
			}
		}
	}
	if (!exist) {
		for i := 0; i < REDO_LOG_GROUP; i++ {
			file, err := os.OpenFile(path.Join(this.logDir, REDO_LOG_NAME+strconv.Itoa(i)), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
			if err != nil {
				panic(err)
			}
			file.WriteAt(make([]byte, REDO_LOG_SIZE), 0)
			file.Sync()
			this.logGroup.file = append(this.logGroup.file, file)
			this.logGroup.getBlock(1);
			fmt.Println()
		}
	}
}
