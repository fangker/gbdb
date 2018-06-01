package im

import (
	"github.com/fangker/gbdb/backend/dm/buffPage"
	"sync"
)

type BPlusTree struct {
	tableID  uint32
	bootPage *pcache.BuffPage
	lock     sync.Mutex
}

func Create(tableID uint32,bootPage pcache.BuffPage){

}




