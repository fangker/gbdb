package mtrs

import (
	"bytes"
	"fmt"
	"github.com/fangker/gbdb/backend/cache"
	pcache "github.com/fangker/gbdb/backend/cache/buffPage"
	"github.com/fangker/gbdb/backend/conf"
	"github.com/fangker/gbdb/backend/file"
	"github.com/fangker/gbdb/backend/mtrs/mlog"
	"github.com/fangker/gbdb/backend/mtrs/mtr"
	"github.com/fangker/gbdb/backend/utils/uassert"
	"reflect"
	"testing"
	"unsafe"
)

func TestMtr(t *testing.T) {
	fmt.Println("=======", conf.GetServerStartConfig().DbDirPath)
	cache.NewCacheBuffer(40)
	fSys := file.CreateFilSys()
	fSys.CreateFilSpace("space", 1, conf.GetServerStartConfig().DbDirPath+"/test/", 1, 16*1024)
	mtr1 := mtr.Start()
	pg := cache.CP.GetPage(1, 1, pcache.BP_X_LOCK, mtr1)
	data := pg.GetData()
	uassert.True(mtr1.IsMemoContains(mtr.MTR_MEMO_PAGE_X_LOCK, pg))
	//mLog 测试
	mlog.WriteUint(&data[0], 999, mlog.MLOG_TYPE_BYRE_2, mtr1)
	r := *(*bytes.Buffer)(unsafe.Pointer(reflect.ValueOf(mtr1).Elem().FieldByName("log").UnsafeAddr()))
	uassert.True(reflect.DeepEqual(r.Bytes(), []byte{0, 1, 0, 1, 0, 0, 1, 3, 231}))
	mtr.Commit(mtr1)
}
