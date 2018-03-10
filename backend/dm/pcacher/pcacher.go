package pcache

import ("sync"
"github.com/fangker/gbdb/backend/dm/page"
)

type PCacher interface {

}

type pcache struct {
	mux *sync.Mutex
	//page
	pType uint16
}