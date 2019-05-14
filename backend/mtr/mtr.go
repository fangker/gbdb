package mtr

import (
	"github.com/fangker/gbdb/backend/cache/buffPage"
	. "github.com/fangker/gbdb/backend/constants/cType"
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
interface
type mo struct {
	mode uint
	obj  *pcache.BuffPage
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

func (this *Mtr) AddToMemo(lockMode uint, obj *pcache.BuffPage) *Mtr {
	mo := mo{mode: lockMode, obj: obj}
	this.memo = append(this.memo, mo);
	if (X_LOCK == lockMode) {
		mo.obj.RLock()
	}
	return this;
}
