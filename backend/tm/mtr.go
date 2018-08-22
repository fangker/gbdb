package tm

import (
	"github.com/fangker/gbdb/backend/dm/buffPage"
	. "github.com/fangker/gbdb/backend/dm/constants/cType"
)

const (
	TR_NONE_LOG      = 0
	TR_ALL_LOG       = 1
	TR_NONE_UNDO_LOG = 2
)

func (this *Transaction) GetDatum() int {
	return int(this.TrxID)
}

type Transaction struct {
	TrxID     XID
	usePage  [] *pcache.BuffPage
	log      []byte
	nLogRecs uint32
	logMode  int
	startLsn LSN
	endLsn   LSN
}

// redo & undp data
func (this *Transaction) SetData(pageNo uint32, offset uint32, bytes []byte) {

}

// redo
func (this *Transaction) WriteLog(pageNo uint32, offset uint32, bytes []byte) {

}

