package redo

import (
	"fmt"
	"github.com/fangker/gbdb/backend/conf"
	"path"
	"testing"
)

func TestLogSys(t *testing.T) {
	fmt.Printf("%+v", conf.GetConfig())
	NewLogSys(path.Join("", "data"))
}
