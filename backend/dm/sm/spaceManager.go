package spaceManager

import "os"

type SpaceManager struct {
	filePath string
	tableNo  uint32
	file   *os.File
}

func NewSpaceManager(filePath string, tableNo uint32) *SpaceManager {
	file,err := os.OpenFile(filePath,os.O_APPEND|os.O_CREATE|os.O_TRUNC,0777)
	if err != nil {
		panic(err)
	}
	return &SpaceManager{filePath,tableNo,file}
}
// 初始化一个文件 设定初始页面构造文件结构
func (sm *SpaceManager) InitFileStructure(){

}