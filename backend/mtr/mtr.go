package mtr

import (
	"github.com/fangker/gbdb/backend/dm/buffPage"
	"github.com/fangker/gbdb/backend/dm/constants/cType"
	"github.com/fangker/gbdb/backend/cache/system"
	"sync/atomic"
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

type Transaction struct {
	usePage  [] *pcache.BuffPage
	log      []byte
	nLogRecs uint32
	logMode  uint8
	startLsn Lsn
	endLsn   Lsn
}

func NewTransaction() *Transaction {
	atomic.AddUint64(&sysTableID, 1)
	return &Transaction{}
}
