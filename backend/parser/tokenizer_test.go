package parser

import (
	"testing"

)

func TestTokenizer(t *testing.T) {
	var a = "update student66 set name='ZYJ' where id = 5"
	Parse(a)
}
