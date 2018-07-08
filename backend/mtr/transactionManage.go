package mtr

import (
	"github.com/fangker/gbdb/backend/dm/buffPage"
	"github.com/fangker/gbdb/backend/dm/constants/cType"
	"github.com/fangker/gbdb/backend/cache"
)

type (
	Lsn cType.Lsn
)

var SYS_MTR_ID uint64 = 0

func LoadTransaction(){
	cache.SC.
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
	return &Transaction{}
}

