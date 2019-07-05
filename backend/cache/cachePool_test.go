package cache

import (
	"testing"
	"fmt"
	"unsafe"
)

func TestCacheBuffer(t *testing.T) {
	cb := NewCacheBuffer(4)
	fmt.Println(uintptr(unsafe.Pointer(cb.blockPages[0].Ptr)))
	data := cb.blockPages[0].GetData()
	data[0] = 'a'
	fmt.Println(cb.blockPages[0].GetData())
	fmt.Println(data[0])
	cb.blockPageAlign((*byte)(unsafe.Pointer(data)))
}
