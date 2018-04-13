package cache

import (
	"testing"
)


func TestCacheBuffer(t *testing.T) {
	cb:=NewCacheBuffer(22)
	bp:=cb.GetFreePage()
	bp.SetType(1)
	t.Log(bp.Page)
}
