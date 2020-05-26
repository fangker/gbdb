package mtr_test

import (
	"fmt"
	"github.com/fangker/gbdb/backend/cache"
	"github.com/fangker/gbdb/backend/cache/buffPage"
	"github.com/fangker/gbdb/backend/conf"
	"github.com/fangker/gbdb/backend/file"
	"github.com/fangker/gbdb/backend/mtr"
	"github.com/fangker/gbdb/backend/utils/uassert"
	"testing"
)

func TestMtr(t *testing.T) {
	cache.NewCacheBuffer(40)
	fSys := file.CreateFilSys()
	fSys.CreateFilSpace("space", 1, conf.GetServerStartConfig().DbDirPath+"/test/", 1, 1024)
	mtr1 := mtr.Start()
	pg := cache.CP.GetPage(1, 1, pcache.BP_X_LOCK, mtr1)
	data := pg.GetData()
	uassert.True(mtr1.IsMemoContains(mtr.MTR_MEMO_PAGE_X_LOCK, pg))
	mLog := mtr.MLogOpen()
	mtr.MLogInitialRecord(&data[1], mLog)
	mtr.MLogClose(mtr1,mLog)
	mtr1.PrintDetail()
	fmt.Println("-->", mLog)
}
