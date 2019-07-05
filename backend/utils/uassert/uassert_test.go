package uassert

import (
	"testing"
)

func Test_False(t *testing.T) {
	False(1 != 1)
	Truef(1 != 1, "ERROR %d %s", 1, "reasons")
}
