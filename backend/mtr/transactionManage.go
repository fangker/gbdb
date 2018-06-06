package mtr

import (
	"github.com/fangker/gbdb/backend/dm/buffPage"
	"github.com/fangker/gbdb/backend/dm/constants/cType"
)

type (
	Lsn cType.Lsn
)

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
