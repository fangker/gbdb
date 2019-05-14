package gbdb

import (
	"testing"
	"fmt"
)

func main() {
	//定义一个切片
	var a []int
	//给a分配一个内存（初始化）
	a = make([]int, 10)
	//给切片赋值
	a[0] = 100
	a[1] = 9
	c := &a[1]
	fmt.Println(a, c + 1)

	//定义一个切片类型的指针
	var p *[]int
	//给指针初始化
	p = new([]int)
	//给切片初始化，分配一个切片内存
	(*p) = make([]int, 10)
	//给切片赋值
	(*p)[0] = 100
	fmt.Println(p)

	p = &a
	(*p)[0] = 1000
	(*p)[1] = 10
	(*p)[2] = 11
	(*p)[3] = 12
	(*p)[4] = 13
	fmt.Println(a)
}

func Test(T *testing.T) {
	main()
}
