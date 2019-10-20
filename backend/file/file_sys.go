package file

import (
	"sync"
)

var IFileSys *FileSys;

type FileSys struct {
	hSpaces map[uint64]*fileSpace
	sync.Mutex
}

func CreateFilSys() *FileSys {
	IFileSys = &FileSys{hSpaces: make(map[uint64]*fileSpace)};
	return IFileSys
}

func (fsys *FileSys) AddSpace(fs *fileSpace) *FileSys {
	fsys.Lock()
	defer  func() {
		fsys.Unlock()
	}()
	fsys.hSpaces[fs.id] = fs;
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
