package file

import (
	"sync"
	"os"
	"github.com/fangker/gbdb/backend/utils/uassert"
	. "github.com/fangker/gbdb/backend/def/constant"
)

type fileSpace struct {
	name     string
	id       uint64
	size     uint64
	sType    int
	filUnits []*filUnit
	sync.Mutex
}

func (fs *fileSpace) CreateFilUnit(filPath string, size uint64) *filUnit {
	fs.Lock();
	defer fs.Unlock()
	f, err := os.OpenFile(filPath, os.O_CREATE|os.O_RDWR, 0660)
	if err != nil {
		panic(err.Error())
	}
	//info, _ := os.Stat(filPath)
	//fs.size = uint64(info.Size())
	fu := &filUnit{filPath: filPath, os: f, size: size}
	fs.size+=size;
	fs.filUnits = append(fs.filUnits, fu)
	return fu;
}

func (fs *fileSpace) Read(spaceId uint64, offset uint64, b []byte) {
	fileIo(FILE_IO_READ, spaceId, offset, b)
}

func (fs *fileSpace) Write(spaceId uint64, offset uint64, b []byte) {
	fileIo(FILE_IO_WRITE, spaceId, offset, b)
}

const (
	FILE_IO_READ  = iota
	FILE_IO_WRITE
)

func fileIo(action int, spaceId uint64, offset uint64, b []byte) {
	fs := IFileSys.GetSpace(spaceId);
	fs.Lock();
	defer fs.Unlock()
	var unitIndex uint64 = 0;
	// 寻找unit
	for ; unitIndex <= uint64(len(fs.filUnits)); unitIndex++ {
		if (fs.filUnits[unitIndex].size > offset) {
			break
		} else {
			offset -= fs.filUnits[unitIndex].size
		}
	}
	uassert.True(offset%LOG_BLOCK_SIZE == 0)
	uassert.True(uint64(len(b))%LOG_BLOCK_SIZE == 0)
	if action == FILE_IO_WRITE {
		fs.filUnits[unitIndex].write(offset, b)
	}
	if action == FILE_IO_READ {
		fs.filUnits[unitIndex].read(offset, b)
	}
}
