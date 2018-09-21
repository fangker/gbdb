package gbdb

import (
	"testing"
)


func Test(T *testing.T)  {
	s:=&[]byte{'a','b'}
	b:=s[0:1]
	copy(b,[]byte{'d'})
}

func ss( s *[]byte)  {
	s
}