package mtr

import (
	"github.com/fangker/gbdb/backend/dm/buffPage"
	"github.com/fangker/gbdb/backend/dm/constants/cType"
	"github.com/fangker/gbdb/backend/cache/system"
	"sync/atomic"
	"github.com/fangker/gbdb/backend/im"
)

type (
	Lsn cType.Lsn
)

var sysTableID uint64 = 0
var systemCache *sc.SystemCache

func LoadTransactionManage(scp *sc.SystemCache) {
	systemCache = scp
	sysTableID = scp.SysTrxIDStore().HdrTableID()
}

// 全局维护事物的相关信息
type TransactionManage struct {
	rwTrxList *im.SortList
}

func NewTransactionManage() {
	var this = &TransactionManage{}
	this.rwTrxList = im.NewSortList()
}

func (this *TransactionManage) AddToRWTrxList(tr *Transaction) {
	this.rwTrxList.AddTo(tr)
}

func (this *Transaction) GetDatum() int {
	return int(this.trID)
}

type Transaction struct {
	trID     uint64
	usePage  [] *pcache.BuffPage
	log      []byte
	nLogRecs uint32
	logMode  uint8
	startLsn Lsn
	endLsn   Lsn
}

func NewTransaction() *Transaction {
	trID := atomic.AddUint64(&sysTableID, 1)
	return &Transaction{trID: trID}
}
