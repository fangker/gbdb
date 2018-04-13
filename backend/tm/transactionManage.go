package tm

import "github.com/fangker/gbdb/backend/cache"

type transactionManage struct {
	usePage []*cache.CachePool
	TID  string
}