package page

import (
	"github.com/fangker/gbdb/backend/cache/buffPage"
	"github.com/fangker/gbdb/backend/def/cType"
	"github.com/fangker/gbdb/backend/utils"
	"github.com/fangker/gbdb/backend/wrapper"
)

type FSPageTrx struct {
	FH           *FilHeader
	BP           *pcache.BuffPage
	data         *cType.PageData
	cacheWrapper wp.Wrapper
}

func NewFSPageTrx(bp *pcache.BuffPage) *FSPageTrx {
	fsPageTrx := &FSPageTrx{
		data: bp.GetData(),
		FH:   &FilHeader{data: bp.GetData()},
		BP:   bp,
	}
	return fsPageTrx;
}

func(this *FSPageTrx) SysTrxIDStore() uint32{
  return utils.GetUint32(this.data[FIL_HEADER_END:FIL_HEADER_END+8])
}

func(this *FSPageTrx) SetSysTrxIDStore(trxID uint32) {
  copy(this.data[FIL_HEADER_END:FIL_HEADER_END+8],utils.PutUint32(trxID))
}

// TRX_SYS_TRX_ID_STORE
// TRX_SYS_FSEG_HEADER
// TRX_SYS_RSEGS 128 rollback segment

