package file

import (
	"testing"
	"github.com/fangker/gbdb/backend/utils/ulog"
)

func TestFileSys(t *testing.T) {
	fsys := CreateFilSys();
	fs := fsys.CreateFilSpace("space", 1, 1)
	fs.CreateFilUnit("./space1.db", 1*1024)
	fs.CreateFilUnit("./space2.db", 16*1024)
	var a = &[512]byte{'a', 'b', 'c'}
	var c = &[512]byte{}
	fs.Write(1, 16*1024, a[0:512])
	fs.Read(1, 16*1024, c[0:512])
	ulog.Caption(c)
}
