package srv

import (
	"testing"
	"fmt"
	"github.com/fangker/gbdb/backend/utils/ulog"
)

func TestServer(T *testing.T) {
	fmt.Println(ulog.AnyViewToString(getSereStartConfig()))
}
