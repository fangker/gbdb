package wp

import "os"

/*
static wrapper for global file location
*/
type Wrapper struct {
	SpaceID uint32
	TableID uint32
	PageNo  uint32
	File    *os.File
}

func GetWrapper(SpaceID, TableID uint32, pageNo uint32, File *os.File) Wrapper {
	return Wrapper{SpaceID, TableID, pageNo, File}
}

//func (wp *Wrapper) tableID() uint32 {
//	return wp.TableID;
//}
//
//func (wp *Wrapper) file() *os.File {
//	return wp.File;
//}
