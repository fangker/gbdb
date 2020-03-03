package mtr_test

import (
	"github.com/fangker/gbdb/backend/cache"
	"testing"
	"github.com/fangker/gbdb/backend/file"
	"github.com/fangker/gbdb/backend/conf"
	"github.com/fangker/gbdb/backend/cache/buffPage"
	"github.com/fangker/gbdb/backend/mtr"
	"fmt"
	"github.com/fangker/gbdb/backend/utils/uassert"
)

func TestMtr(t *testing.T) {
	cache.NewCacheBuffer(40);
	fsys := file.CreateFilSys()
	fsys.CreateFilSpace("space", 1, conf.GetServerStartConfig().DbDirPath+"/test/", 1, 1024)
	mtr_1 := mtr.MtrStart()
	pg := cache.CP.GetPage(1, 1, pcache.BP_X_LOCK, mtr_1)
	data := pg.GetData()
	uassert.True(mtr_1.IsMemoContains(mtr.MTR_MEMO_PAGE_X_LOCK, pg))
	fmt.Println("-->", data)
}
