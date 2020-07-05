package file

import (
	"fmt"
	"github.com/fangker/gbdb/backend/def/cType"
	"os"
	"unsafe"
)

type filUnit struct {
	filPath string
	size    uint64
	os      *os.File
}

func (fu *filUnit) read(pos uint64, b *byte) {
	fu.os.Seek(int64(pos), 0)
	fu.os.Read((*cType.PageData)(unsafe.Pointer(b))[0:])
}

func (fu *filUnit) write(pos uint64, b *byte) {
	_, err := fu.os.WriteAt((*cType.PageData)(unsafe.Pointer(b))[0:], int64(pos))
	fmt.Println(fu.os.Name(), (*cType.PageData)(unsafe.Pointer(b))[0:])
	if err != nil {
		panic(err)
	}

	fu.os.Sync()
}
