package mtr

import (
	"bytes"
	. "github.com/fangker/gbdb/backend/def/cType"
	"github.com/fangker/gbdb/backend/utils/ulog"
)

const (
	MTR_NONE_LOG = 0
	MTR_ALL_LOG  = 1
)
const (
	MTR_MEMO_PAGE_X_LOCK MemoLock = 0
	MTR_MEMO_PAGE_S_LOCK MemoLock = 1
	MTR_MEMO_BUF_FIX     MemoLock = 3
	MTR_MEMO_S_LOCK      MemoLock = 11
	MTR_MEMO_L_LOCK      MemoLock = 12
)

var MTR_MEMO_PAGE_X_LOCK_1 = 2

type MemoLock uint

type mo struct {
	mode MemoLock
	obj  ObjLocker
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

func (mtr *Mtr)PrintDetail(){
	ulog.Debug(ulog.AnyViewToString(mtr))
}
func Start() *Mtr {
	mtr := &Mtr{}
	mtr.logMode = MTR_NONE_LOG
	return mtr
}

func Commit(mtr *Mtr) bool {
	mtr.modification = true
	mtr.logMode = MTR_ALL_LOG
	return true
}

func (mtr *Mtr) AddToMemo(lockMode MemoLock, obj ObjLocker) *Mtr {
	if mtr.IsMemoContains(lockMode, obj) {
		return mtr
	}
	mo := mo{mode: lockMode, obj: obj}
	mtr.memo = append(mtr.memo, mo)
	return mtr
}

func (mtr *Mtr) IsMemoContains(lockMode MemoLock, t ObjLocker) bool {
	for _, a := range mtr.memo {
		if t == a.obj && lockMode == a.mode {
			return true
		}
	}
	return false
}

// * MTR  Obj Locker
type ObjLocker interface {
	Lock()
	RLock()
}

