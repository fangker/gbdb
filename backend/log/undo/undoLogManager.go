package undo


type UndoLogManager struct {
	TableID   uint32
	TableName string
	ufm   *UndoFileManager
}

func NewUndoLogManager(ufm *UndoFileManager, tableName string) *UndoLogManager {
	this := &UndoLogManager{ufm: ufm, TableName: tableName, TableID: ufm.TableID}
	return this
}