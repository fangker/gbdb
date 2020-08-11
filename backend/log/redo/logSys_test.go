package redo

import (
	"fmt"
	"github.com/fangker/gbdb/backend/conf"
	"github.com/fangker/gbdb/backend/file"
	"testing"
)

func TestLogSys(t *testing.T) {
	fmt.Printf("%+v \r", conf.GetConfig())
	fsys := file.CreateFilSys()
	c := NewLogSys(fsys, 3, 16*1024<<10)
	c.WriteToLog([]byte{1})
	fmt.Println(bufferSize-c.LogRemain() == 1)
}
