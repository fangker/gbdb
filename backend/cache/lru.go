package cache

import (
	"container/list"
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
	Capacity int
	dlist    *list.List
	cacheMap map[interface{}]*list.Element
}

func NewLRUCache(cap uint32) (*LRUCache) {
	return &LRUCache{
		Capacity: int(cap),
		dlist:    list.New(),
		cacheMap: make(map[interface{}]*list.Element)}
}

func (lru *LRUCache) Size() (int) {
	return lru.dlist.Len()
}

func (lru *LRUCache) Set(k, v interface{}) {
	if pElement, ok := lru.cacheMap[k]; ok {
		lru.dlist.MoveToFront(pElement)
		pElement.Value.(*CacheNode).Value = v
		return
	}

	newElement := lru.dlist.PushFront(&CacheNode{k, v})
	lru.cacheMap[k] = newElement
	if lru.dlist.Len() > lru.Capacity {
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

	if pElement, ok := lru.cacheMap[k]; ok {
		lru.dlist.MoveToFront(pElement)
		return pElement.Value.(*CacheNode).Value, true, nil
	}
	return v, false, nil
}

func (lru *LRUCache) Remove(k interface{}) (bool) {

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
