package tm

import (
	"github.com/fangker/gbdb/backend/dm/buffPage"
)




func (this *Transaction) GetDatum() int {
	return int(this.trID)
}

type Transaction struct {
	trID     XID
	usePage  [] *pcache.BuffPage
	log      []byte
	nLogRecs uint32
	logMode  uint8
	startLsn LSN
	endLsn   LSN
}

// redo & undp data
func (this *Transaction) SetData(pageNo uint32,offset uint32,bytes []byte) {

}
