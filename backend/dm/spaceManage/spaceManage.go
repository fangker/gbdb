package spaceManage

import (
	"os"
)


var SM *SpaceManage

type SpaceManage struct {
	tf  map[uint32] *tableFileManage
}


func NewSpaceManage()*SpaceManage{
	return &SpaceManage{tf:make(map[uint32] *tableFileManage)}
}

func (sm *SpaceManage)Add(tfm *tableFileManage) *tableFileManage{
	sm.tf[tfm.tableID]=tfm
	return tfm
}

func NewTableFileManage(filePath string, tableID uint32) *tableFileManage {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	return &tableFileManage{filePath, tableID, file,}
}

func (sm *SpaceManage) initSysFileStructure()  {
    sm.tf[0].initSysFile()

}

func TF(tf uint32) *tableFileManage{
	return SM.tf[tf]
}
