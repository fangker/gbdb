package gbdb

import (
	"testing"
	"time"
	"fmt"
)

type N struct {
	s string
}
type Student struct {
	name string
	age  int
	n N
}

func Test(T *testing.T)  {
	begin := time.Now().UnixNano()
	defer func() {
		end := time.Now().UnixNano()
			fmt.Println((end - begin) , "ns")
	}()
	fmt.Println((100000000*(100000000+1))/2)
}

func TestAcc(T *testing.T)  {
	begin := time.Now().UnixNano()
	defer func() {
		end := time.Now().UnixNano()
		fmt.Println((end - begin), "ns")
	}()
	var  a int
	for i:=0;i<100000000;i++{
		a+=a+i
	}
	fmt.Println(a)
}
