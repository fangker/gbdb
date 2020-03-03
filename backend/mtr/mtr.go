package mtr

import (
	. "github.com/fangker/gbdb/backend/def/cType"
	"bytes"
)

const (
	MTR_NONE_LOG = 0
	MTR_ALL_LOG  = 1
)
const (
	MTR_MEMO_PAGE_X_LOCK MtrMemoLock = 0
	MTR_MEMO_PAGE_S_LOCK MtrMemoLock = 1
	MTR_MEMO_BUF_FIX     MtrMemoLock = 3
	MTR_MEMO_S_LOCK      MtrMemoLock = 11
	MTR_MEMO_L_LOCK      MtrMemoLock = 12
)

var MTR_MEMO_PAGE_X_LOCK_1 = 2

type MtrMemoLock uint

type mo struct {
	mode MtrMemoLock
	obj  MtrObjLocker
}
type Mtr struct {
	TrxID        XID
	memo         [] mo
	log          bytes.Buffer
	nLogRecs     uint32
	logMode      int
	startLsn     LSN
	endLsn       LSN
	modification bool
}

func MtrStart() *Mtr {
	mtr := &Mtr{}
	mtr.logMode = MTR_NONE_LOG;
	return mtr;
}

func MtrCommit(mtr *Mtr) bool {
	mtr.modification = true
	mtr.logMode = MTR_ALL_LOG;
	return true
}

func (mtr *Mtr) AddToMemo(lockMode MtrMemoLock, obj MtrObjLocker) *Mtr {
	if (mtr.IsMemoContains(lockMode, obj)) {
		return mtr
	}
	mo := mo{mode: lockMode, obj: obj}
	mtr.memo = append(mtr.memo, mo);
	return mtr;
}

func (mtr *Mtr) IsMemoContains(lockMode MtrMemoLock, t MtrObjLocker) bool {
	for _, a := range mtr.memo {
		if (t == a.obj && lockMode == a.mode) {
			return true
		}
	}
	return false
}

// * MTR  Obj Locker
type MtrObjLocker interface {
	Lock()
	RLock()
}

