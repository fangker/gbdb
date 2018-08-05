package im

import (
	"sort"
	"sync"
)

type ListItem interface {
	GetDatum() int
}

type SortList struct {
	l      []ListItem
	rwLock *sync.RWMutex
}

func (this SortList) add(i int, v ListItem) {
	this.rwLock.Lock()
	defer this.rwLock.Unlock()
	temp := append(make([]ListItem, 0), this.l[i:]...)
	ss := append(this.l[:i], v)
	ss = append(ss, temp...)
}

func (this SortList) searchIndex(value int) int {
	return sort.Search(len(this.l), func(i int) bool { return this.l[i].GetDatum() <= value });
}

func (this SortList) AddTo(v ListItem) {
	if len(this.l) == 0 {
		this.l = append(this.l, v);
	}
	this.add(this.searchIndex(v.GetDatum()), v)
}

func (this SortList) search(v ListItem) ListItem {
	this.rwLock.RLock()
	defer this.rwLock.RUnlock()
	return this.l[this.searchIndex(v.GetDatum())]
}

func (this SortList) remove(i int) {
	this.rwLock.Lock()
	defer this.rwLock.Unlock()
	temp := append(make([]ListItem, 0), this.l[i+1:]...)
	ss := this.l[:i]
	ss = append(ss, temp...)
	this.l = ss
}

func NewSortList() *SortList {
	return &SortList{
		l:      make([]ListItem, 0),
		rwLock: &sync.RWMutex{},
	}
}
