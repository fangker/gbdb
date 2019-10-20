package gbdb

import (
	"testing"
	"fmt"
	"unsafe"
)

func test1() {
	//定义一个切片类型的指针
	//给指针初始化
	mm := new([66]byte)
	a := (*[3]byte)(unsafe.Pointer(&mm))
	*a = [3]byte{'s'}
	fmt.Println(*a, 11111)
	u := unsafe.Pointer(mm)
	p := (*[3]byte)(u)
	fmt.Println(mm, &mm, &u)
	*p = [3]byte{'a', 'b', 'c'}
	*p = [3]byte{'f', 'f', 'f'}
	fmt.Println(mm, *p)
}

func Test(T *testing.T) {
	test1()
}
