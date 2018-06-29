package redo

import (
	"testing"
	"path"
	"fmt"
	"github.com/fangker/gbdb/backend/utils"
)


func TestLogSys(t *testing.T) {
	fmt.Println(utils.DATA_DIR)
	NewLogSys(path.Join("","data"))
}