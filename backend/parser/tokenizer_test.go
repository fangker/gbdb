package parser

import (
	"testing"
)

func TestTokenizer(t *testing.T) {
	var a = "update student66 set name='ZYJ' where id = 5"
	update,err:=Parse(a)
	if(err!=nil){
		t.Log(err)
		t.Fail()
	}
	t.Log(update)
	a= "select name gender  from people where age = 3 "
	read,err:=Parse(a)
		if(err!=nil){
			t.Log(err)
			t.Fail()
		}
		t.Log(read)
}
