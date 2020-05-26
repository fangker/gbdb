package tm

import (
	"github.com/fangker/gbdb/backend/im"
	"sync/atomic"
	"unsafe"
	. "github.com/fangker/gbdb/backend/def/cType"
	"github.com/fangker/gbdb/backend/dm/page"
	"github.com/fangker/gbdb/backend/mtrs"
)

type tmSysInfoManager interface {
	SysTrxIDStore() page.DictPager
}

var TM *TransactionManage
// 全局维护事物的相关信息
type TransactionManage struct {
	TrID        XID
	rwTrxList   *im.SortList
	systemCache tmSysInfoManager
}

func NewTransactionManage(scp tmSysInfoManager) *TransactionManage {
	var this = &TransactionManage{}
	this.TrID = XID(scp.SysTrxIDStore().HdrTableID());
	this.rwTrxList = im.NewSortList()
	TM = this;
	return this;
}

func (this *TransactionManage) AddToRWTrxList(tr *mtrs.Transaction) {
	this.rwTrxList.AddTo(tr)
}

func (tm *TransactionManage) generateXID() XID {
	atomic.AddUint64((* uint64)(unsafe.Pointer(&tm.TrID)), 1)
	// TODO: 256累加
	return tm.TrID
}

func (tm *TransactionManage) TrxStart() *mtrs.Transaction {
	trID := tm.generateXID()
	t := &mtrs.Transaction{TrxID: trID}
	tm.AddToRWTrxList(t)
	return t
}

func (tm *TransactionManage) TrxCommit() {

}
