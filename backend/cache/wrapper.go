package cache

import "os"

type Wrapper struct {
	SpaceID uint32
	TableID uint32
	File    *os.File
}

//func (wp *Wrapper) tableID() uint32 {
//	return wp.TableID;
//}
//
//func (wp *Wrapper) file() *os.File {
//	return wp.File;
//}
