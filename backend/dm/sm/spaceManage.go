package spaceManage

import (
	"os"
	"github.com/fangker/gbdb/backend/dm/page"
	"fmt"
)

type SpaceManage struct {
	filePath string
	tableID  uint32
	file     *os.File
}

func NewSpaceManage(filePath string, tableID uint32) *SpaceManage {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	return &SpaceManage{filePath, tableID, file}
}

// 初始化一个文件 设定初始页面构造文件结构
func (sm *SpaceManage) InitFileStructure() {

}

// 初始化一系统表文件 设定初始页面构造文件结构
func (sm *SpaceManage) InitSysFileStructure() {

}

func (sm *SpaceManage) WriteSync(pageNum uint32, data page.PageData) {
	offset := pageNum * page.PAGE_SIZE
	fmt.Print(offset,data)
	sm.file.Seek(int64(offset), 0)
	sm.file.WriteAt(data[:],int64(offset))
	sm.file.Sync()
}
