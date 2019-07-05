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
	defer fsys.Unlock()
	fsys.hSpaces[fs.id] = fs;
	return fsys
}

func (fsys *FileSys) GetSpace(id uint64) *fileSpace {
	fsys.Lock()
	defer fsys.Unlock()
	return fsys.hSpaces[id]
}

func (fsys *FileSys) CreateFilSpace(name string, id uint64, sType int) *fileSpace {
	fsys.Lock()
	defer fsys.Unlock();
	fspace := &fileSpace{name: name, id: id, sType: sType}
	fsys.hSpaces[id] = fspace
	return fspace;
}

func init() {
	IFileSys = CreateFilSys()
}
