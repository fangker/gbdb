package uassert

import (
	"testing"
	"fmt"
)

func Test_Uassert(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Print(err);
		}
	}()
	False(1 != 1)
	Truef(1 != 1, "ERROR %d %s", 1, "reasons")
}
