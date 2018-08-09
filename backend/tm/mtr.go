package tm

import (
	"github.com/fangker/gbdb/backend/dm/buffPage"
	"sync/atomic"
)




func (this *Transaction) GetDatum() int {
	return int(this.trID)
}

type Transaction struct {
	trID     uint64
	usePage  [] *pcache.BuffPage
	log      []byte
	nLogRecs uint32
	logMode  uint8
	startLsn LSN
	endLsn   LSN
}

func (this *TransactionManage)NewTransaction() *Transaction {
	trID := atomic.AddUint64(&sysTableID, 1)
	return &Transaction{trID: trID}
}
