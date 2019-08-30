package mtr

import (
	"github.com/fangker/gbdb/backend/cache/buffPage"
	. "github.com/fangker/gbdb/backend/def/cType"
)

const (
	MTR_NONE_LOG      = 0
	MTR_ALL_LOG       = 1
	MTR_NONE_UNDO_LOG = 2
)
const (
	X_LOCK = 0
	S_LOCK = 1
)

type mo struct {
	mode uint
	obj  *pcache.BlockPage
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

func MtrStart(this *Mtr) *Mtr {
	this.logMode = MTR_ALL_LOG;
	return this;
}

func (mtr *Mtr) AddToMemo(lockMode uint, obj *pcache.BlockPage) *Mtr {
	if (mtr.isInMemo(obj)) {
		return mtr
	}
	mo := mo{mode: lockMode, obj: obj}
	mtr.memo = append(mtr.memo, mo);
	switch (lockMode) {
	case X_LOCK:
		obj.Lock();
	case S_LOCK:
		obj.RLock();
	}
	return mtr;
}

func (mtr *Mtr) isInMemo(t *pcache.BlockPage) bool {
	for _, a := range mtr.memo {
		if (t == a.obj) {
			return true
		}
	}
	//for _, a := range mtr.memo {
	//	p, s := a.obj.GetPos()
	//	_p, _s := t.GetPos()
	//	if (p == _p && _s == s) {
	//		return true
	//	}
	//}
	return false
}

func MtrWriteUnt8(this *Mtr, pos , val uint8) {

}
