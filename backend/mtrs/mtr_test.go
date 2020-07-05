package mtrs

import (
	"fmt"
	"github.com/fangker/gbdb/backend/cache"
	"github.com/fangker/gbdb/backend/cache/buffPage"
	"github.com/fangker/gbdb/backend/conf"
	"github.com/fangker/gbdb/backend/file"
	"github.com/fangker/gbdb/backend/mtrs/mlog"
	"github.com/fangker/gbdb/backend/mtrs/mtr"
	"github.com/fangker/gbdb/backend/utils/uassert"
	"testing"
)

func TestMtr(t *testing.T) {
	fmt.Println("=======", conf.GetServerStartConfig().DbDirPath)
	cache.NewCacheBuffer(40)
	fSys := file.CreateFilSys()
	fSys.CreateFilSpace("space", 1, conf.GetServerStartConfig().DbDirPath+"/test/", 1, 1024)
	mtr1 := mtr.Start()
	pg := cache.CP.GetPage(1, 1, pcache.BP_X_LOCK, mtr1)
	data := pg.GetData()
	uassert.True(mtr1.IsMemoContains(mtr.MTR_MEMO_PAGE_X_LOCK, pg))
	// mLog 测试
	mlog.WriteUint(&data[0], 299, mlog.MLOG_TYPE_BYRE_2)
	var c = []byte{1, 2, 3}
	fmt.Println("======", c)
}
