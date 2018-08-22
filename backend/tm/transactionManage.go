package tm

import (
	"github.com/fangker/gbdb/backend/cache/system"
	"github.com/fangker/gbdb/backend/im"
	"sync/atomic"
	"unsafe"
	. "github.com/fangker/gbdb/backend/dm/constants/cType"
)



var TM *TransactionManage
// 全局维护事物的相关信息
type TransactionManage struct {
	TrID        XID
	rwTrxList   *im.SortList
	systemCache *sc.SystemCache
}

func NewTransactionManage(scp *sc.SystemCache) *TransactionManage {
	var this = &TransactionManage{}
	this.systemCache = scp
	this.TrID = XID(scp.SysTrxIDStore().HdrTableID());
	this.rwTrxList = im.NewSortList()
	TM = this;
	return this;
}

func (this *TransactionManage) AddToRWTrxList(tr *Transaction) {
	this.rwTrxList.AddTo(tr)
}

func (tm *TransactionManage) generateXID() XID {
	atomic.AddUint64((* uint64)(unsafe.Pointer(&tm.TrID)), 1)
	// TODO: 256累加
	return tm.TrID
}

func (tm *TransactionManage) TrxStart() *Transaction {
	trID := tm.generateXID()
	return &Transaction{TrxID: trID}
}

func (tm *TransactionManage) TrxCommit() {

}
