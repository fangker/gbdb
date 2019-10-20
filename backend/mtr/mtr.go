package mtr

import (
	. "github.com/fangker/gbdb/backend/def/cType"
)

const (
	MTR_NONE_LOG      = 0
	MTR_ALL_LOG       = 1
	MTR_NONE_UNDO_LOG = 2
)
const (
	MTR_MEMO_PAGE_X_LOCK MtrMemoLock = 0
	MTR_MEMO_PAGE_S_LOCK MtrMemoLock = 1
)

var MTR_MEMO_PAGE_X_LOCK_1 = 2

type MtrMemoLock uint

type mo struct {
	mode MtrMemoLock
	obj  MtrLockObjer
}
type Mtr struct {
	TrxID        XID
	memo         [] mo
	log          []byte
	nLogRecs     uint32
	logMode      int
	startLsn     LSN
	endLsn       LSN
	modification bool
}

func MtrStart() *Mtr {
	mtr := &Mtr{}
	mtr.logMode = MTR_ALL_LOG;
	return mtr;
}

func (mtr *Mtr) AddToMemo(lockMode MtrMemoLock, obj MtrLockObjer) *Mtr {
	if (mtr.IsMemoContains(lockMode, obj)) {
		return mtr
	}
	mo := mo{mode: lockMode, obj: obj}
	mtr.memo = append(mtr.memo, mo);
	return mtr;
}

func (mtr *Mtr) IsMemoContains(lockMode MtrMemoLock, t MtrLockObjer) bool {
	for _, a := range mtr.memo {
		if (t == a.obj && lockMode == a.mode) {
			return true
		}
	}
	return false
}

// * MTR obj Lock upgrade

type MtrLockObjer interface {
	Lock()
}

//func MtrWriteUnt8(this *Mtr, pos, val uint8) {
//
//}
