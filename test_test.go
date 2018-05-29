package gbdb

import (
	"testing"
	"container/list"
	"fmt"
)

type A struct {
	name string
}

func Test(T *testing.T)  {
	ls:=list.New()
	a:=&A{name:"2222"}
	ls.PushFront(a)
	b:=ls.Front().Value.(*A)
	b.name="2333"
	fmt.Println(a,b)
}
