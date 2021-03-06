package file

import (
	"github.com/fangker/gbdb/backend/def/constant"
	"github.com/fangker/gbdb/backend/utils/uassert"
	"math"
	"os"
	"path"
	"strconv"
	"sync"
)

type fileSpace struct {
	name string
	id   uint64
	// 总大小
	size    uint64
	fileDir string
	// 文件类型
	FType    fileType
	filUnits []*filUnit
	// 自增扩展大小
	autoIncSize   uint64
	nextExtendNum int
	sync.Mutex
}

func (fs fileSpace) Size() uint64 {
	return fs.size
}

//func (fs *fileSpace) scanDirWithFilUnit() {
//	files, _ := ioutil.ReadDir(fs.fileDir)
//	var fileSlice []string;
//	for _, v := range files {
//		if ok, _ := regexp.MatchString(`^`+fs.name+`_\d*.db`, v.Name()); ok {
//			fileSlice = append(fileSlice, v.Name())
//		}
//	}
//	// F-TODO:
//	fs.nextExtendNum = 0;
//	for i, v := range fileSlice {
//		fileName := path.Join(fs.fileDir, fs.name+"_"+strconv.Itoa(i)) + ".db";
//		s, _ := os.Stat(fileName)
//		ulog.Caption(fileName)
//		fs.CreateFilUnit(path.Join(fs.fileDir, v), uint64(s.Size()))
//	}
//}

func (fsys *FileSys) CreateFilSpace(name string, id uint64, cdbSpacePath string, sType fileType, autoIncSize uint64) *fileSpace {
	fsys.Lock()
	defer func() {
		fsys.Unlock()
	}()
	err := os.MkdirAll(path.Dir(cdbSpacePath), os.ModePerm)
	if err != nil {
		panic(err)
	}
	fspace := &fileSpace{name: name, id: id, FType: sType, fileDir: cdbSpacePath, autoIncSize: autoIncSize}
	fsys.hSpaces[id] = fspace
	return fspace
}

func (fs *fileSpace) CreateFilUnit(filPath string, size uint64) *filUnit {
	fs.Lock()
	defer func() {
		fs.Unlock()
	}()
	f, err := os.OpenFile(filPath, os.O_CREATE|os.O_RDWR, 0660)
	if err != nil {
		panic(err.Error())
	}
	fu := &filUnit{filPath: filPath, os: f, size: size}
	fs.size += size
	fs.filUnits = append(fs.filUnits, fu)
	fs.nextExtendNum++
	return fu
}

func (fs *fileSpace) Read(spaceId uint64, offset uint64, b *byte) {
	fileIo(FILE_IO_READ, spaceId, offset, b)
}

func (fs *fileSpace) Write(spaceId uint64, offset uint64, data *byte) {
	fileIo(FILE_IO_WRITE, spaceId, offset, data)
}

const (
	FILE_IO_READ = iota
	FILE_IO_WRITE
)

func fileIo(action int, spaceId uint64, offset uint64, data *byte) {
	fs := IFileSys.GetSpace(spaceId)
	// 检查是否需要扩容
	if fs.size < offset {
		// 扩容函数
		fillUnitCount := math.Ceil(float64((offset - fs.size) / fs.autoIncSize))
		for i := 0; i < int(fillUnitCount); i++ {
			filePath := path.Join(fs.fileDir, fs.name+"_"+strconv.Itoa(fs.nextExtendNum)+"."+fs.FType.FileType())
			fs.CreateFilUnit(filePath, fs.autoIncSize)
		}
	}
	var unitIndex uint64 = 0
	// 寻找unit
	for ; unitIndex < uint64(len(fs.filUnits))-1; unitIndex++ {
		if fs.filUnits[unitIndex].size > offset {
			break
		} else {
			offset -= fs.filUnits[unitIndex].size
		}
	}
	uassert.True(offset%constant.LOG_BLOCK_SIZE == 0)
	if action == FILE_IO_WRITE {
		fs.filUnits[unitIndex].write(offset, data)
	}
	if action == FILE_IO_READ {
		fs.filUnits[unitIndex].read(offset, data)
	}
}

func (fs *fileSpace) destroyFiles() {
	for _, v := range fs.filUnits {
		os.Remove(v.filPath)
	}
}
