package file

import (
	"sync"
)

type fileType int

const (
	TypeTuple fileType = 0
	TypeDB    fileType = 1
	TypeRedo  fileType = 2
)

func (ft fileType) FileType() string {
	switch ft {
	case TypeTuple, TypeDB:
		return "db"
	case TypeRedo:
		return "log"
	default:
		return "unknow"
	}
}

var IFileSys *FileSys

type FileSys struct {
	hSpaces map[uint64]*fileSpace
	sync.Mutex
}

func CreateFilSys() *FileSys {
	IFileSys = &FileSys{hSpaces: make(map[uint64]*fileSpace)}
	return IFileSys
}

func (fsys *FileSys) AddSpace(fs *fileSpace) *FileSys {
	fsys.Lock()
	defer func() {
		fsys.Unlock()
	}()
	fsys.hSpaces[fs.id] = fs
	return fsys
}

func (fsys *FileSys) GetSpace(id uint64) *fileSpace {
	fsys.Lock()
	defer func() {
		fsys.Unlock()
	}()
	return fsys.hSpaces[id]
}

func init() {
	IFileSys = CreateFilSys()
}
