package file

import (
	"github.com/fangker/gbdb/backend/conf"
	"github.com/fangker/gbdb/backend/def/cType"
	"github.com/fangker/gbdb/backend/utils/uassert"
	"reflect"
	"testing"
	"unsafe"
)

func TestFileSys(t *testing.T) {
	fsys := CreateFilSys()
	fspace := fsys.CreateFilSpace("space", 1, conf.GetServerStartConfig().DbDirPath, 1, 1024)
	wd := &cType.PageData{'a', 'b', 'c', 'd'}
	var wdp = unsafe.Pointer(wd)
	rd := &cType.PageData{}
	var rdp = unsafe.Pointer(rd)
	fspace.Write(1, 16*1024, (*byte)(wdp))
	fspace.Read(1, 16*1024, (*byte)(rdp))
	uassert.True(reflect.DeepEqual(wd, rd))
}
