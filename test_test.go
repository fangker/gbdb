package gbdb

import (
	"testing"
	"reflect"
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

func Test(T *testing.T)  {s:=&Student{name:"张某",age:22}

}

func anyViewToString(i interface{}) string {
	rs:=reflect.ValueOf(i)
	if(rs.Kind()==reflect.Slice){
		s:="[ "
		for i:=0;i<rs.Len();i++  {
			s=s+fmt.Sprintf("%+v ",rs.Index(i).Elem())
		}
		s=s+" ]"
		return  s
	}
	if(rs.Kind()==reflect.Struct){
		return fmt.Sprintf("%+v",rs)
	}
	return  fmt.Sprintf("%+v",rs)
}