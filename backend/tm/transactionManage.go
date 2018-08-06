package tm

import (
	"github.com/fangker/gbdb/backend/dm/constants/cType"
	"github.com/fangker/gbdb/backend/cache/system"
	"github.com/fangker/gbdb/backend/im"
	"github.com/fangker/gbdb/backend/dm/buffPage"
	"sync/atomic"
)

type (
	Lsn cType.Lsn
)

var sysTableID uint64 = 0
var systemCache *sc.SystemCache

// 全局维护事物的相关信息
type TransactionManage struct {
	rwTrxList *im.SortList
}

func NewTransactionManage(scp *sc.SystemCache) *TransactionManage {
	systemCache = scp
	sysTableID = scp.SysTrxIDStore().HdrTableID()
	var this = &TransactionManage{}
	this.rwTrxList = im.NewSortList()
	return this;
}

func (this *TransactionManage) AddToRWTrxList(tr *Transaction) {
	this.rwTrxList.AddTo(tr)
}