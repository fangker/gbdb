package dstr

import (
	"container/list"
	"sync"
)

// 自己也实现了一个 支持RUL的linked hash map
//  https://github.com/fangker/go-linkedHashMap

type CacheNode struct {
	Key, Value interface{}
}

func (cnode *CacheNode) NewCacheNode(k, v interface{}) *CacheNode {
	return &CacheNode{k, v}
}

type LRUCache struct {
	rwLock sync.RWMutex
	Capacity int
	dlist    *list.List
	cacheMap map[interface{}]*list.Element
}

func NewLRUCache(cap uint64) (*LRUCache) {
	return &LRUCache{
		Capacity: int(cap),
		dlist:    list.New(),
		cacheMap: make(map[interface{}]*list.Element)}
}

func (lru *LRUCache) Size() (int) {
	lru.rwLock.RLock()
	defer lru.rwLock.RUnlock()
	return lru.dlist.Len()
}

func (lru *LRUCache) Set(k, v interface{}) {
	lru.rwLock.RLock()
	defer lru.rwLock.RUnlock()
	if pElement, ok := lru.cacheMap[k]; ok {
		lru.dlist.MoveToFront(pElement)
		pElement.Value.(*CacheNode).Value = v
		return
	}

	newElement := lru.dlist.PushFront(&CacheNode{k, v})
	lru.cacheMap[k] = newElement
	if lru.dlist.Len() > lru.Capacity&&lru.Capacity!=0{
		lastElement := lru.dlist.Back()
		if lastElement == nil {
			return
		}
		cacheNode := lastElement.Value.(*CacheNode)
		delete(lru.cacheMap, cacheNode.Key)
		lru.dlist.Remove(lastElement)
	}
}

func (lru *LRUCache) Get(k interface{}) (v interface{}, ret bool, err error) {
	lru.rwLock.RLock()
	defer lru.rwLock.RUnlock()
	if pElement, ok := lru.cacheMap[k]; ok {
		lru.dlist.MoveToFront(pElement)
		return pElement.Value.(*CacheNode).Value, true, nil
	}
	return v, false, nil
}

func (lru *LRUCache) Remove(k interface{}) (bool) {
	lru.rwLock.Lock()
	defer lru.rwLock.Unlock()
	if lru.cacheMap == nil {
		return false
	}
	if pElement, ok := lru.cacheMap[k]; ok {
		cacheNode := pElement.Value.(*CacheNode)
		delete(lru.cacheMap, cacheNode.Key)
		lru.dlist.Remove(pElement)
		return true
	}
	return false
}

func (lru *LRUCache) List() (*list.List) {
	return lru.dlist
}
