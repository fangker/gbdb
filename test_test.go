package gbdb

import (
	"testing"
)

type N struct {
	s string
}
type TC struct {
	name string
	age  int
	n N
}

func (TC) Hello() {
}
func (*TC) Hello2() {
}

func Test(T *testing.T)  {
	var t TC
	t.Hello()
}
