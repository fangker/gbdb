package cache

import (
	"testing"
	"fmt"
	"unsafe"
	"github.com/fangker/gbdb/backend/utils/uassert"
)

func TestCacheBuffer(t *testing.T) {
	cb := NewCacheBuffer(4)
	fmt.Println(uintptr(unsafe.Pointer(cb.blockPages[0].Ptr)))
	data := cb.blockPages[0].GetData()
	//fmt.Println(cb.blockPages[0].GetData())
	data[0] = 1;
	block := cb.blockPageAlign(&data[1])
	uassert.True(block.GetData()[0]==1);
}
