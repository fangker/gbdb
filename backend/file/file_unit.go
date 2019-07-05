package file

import "os"

type filUnit struct {
	filPath string
	size    uint64
	os      *os.File
}

func (fu *filUnit) read(pos uint64, b []byte) {
	fu.os.Seek(int64(pos), 0)
	fu.os.Read(b)
}

func (fu *filUnit) write(pos uint64, b []byte) {
	_, err := fu.os.WriteAt(b, int64(pos));
	if (err != nil) {
		panic(err)
	}

	fu.os.Sync()
}