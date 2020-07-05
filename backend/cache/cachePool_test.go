package cache

import (
	"fmt"
	"testing"
)

func TestCacheBuffer(t *testing.T) {
	cb := NewCacheBuffer(4)
    bp:=cb.ReadPageFromFile(1, 0)
    fmt.Println(bp)
}
