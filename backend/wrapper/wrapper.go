package wp

/*
static wrapper for global file location
*/
type Wrapper struct {
	SpaceID uint64
	PageNo  uint64
}

func GetWrapper(SpaceID, pageNo uint64) Wrapper {
	return Wrapper{SpaceID, pageNo}
}

//func (wp *Wrapper) tableID() uint32 {
//	return wp.TableID;
//}
//
//func (wp *Wrapper) file() *os.File {
//	return wp.File;
//}
