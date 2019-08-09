package file

import (
	"testing"
	"github.com/fangker/gbdb/backend/utils/ulog"
	"reflect"
	"github.com/fangker/gbdb/backend/utils/uassert"
)

func TestFileSys(t *testing.T) {
	fsys := CreateFilSys();
	fspace := fsys.CreateFilSpace("./","space", 1, 1)
	var a = &[512]byte{'a', 'b', 'c'}
	var c = &[512]byte{}
	fspace.Write(1, 16*1024, a[0:512])
	fspace.Read(1, 16*1024, c[0:512])
	uassert.True(reflect.DeepEqual(a, c))
	ulog.Caption(c)
	fspace.destroyFiles()
}
