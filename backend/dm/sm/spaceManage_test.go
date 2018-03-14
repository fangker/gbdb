package spaceManage

import (
	"testing"
	"github.com/fangker/gbdb/backend/dm/page"
)


func TestSpaceManage(t *testing.T) {
	 sm:=NewSpaceManage("../temData/a.db",0)
	 sm.WriteSync(2,page.PageData{1,1,1,1,1,1})
}
