package algorithm

import (
	"container/list"
	"errors"
)

type CacheNode struct {
	key, value interface{}
}

type LRUCache struct {
	Capacity int
	dList    *list.List
	cacheMap map[interface{}]*list.Element
}

func (c *CacheNode) NewCacheNode(k, v interface{}) *CacheNode {
	return &CacheNode{key: k, value: v}
}

func NewLRUCache(cap int) *LRUCache {
	return &LRUCache{Capacity: cap, dList: list.New(), cacheMap: make(map[interface{}]*list.Element)}
}

func (lru *LRUCache) Size() int {
	return lru.dList.Len()
}

func (lru *LRUCache) Set(k, v interface{}) error {
	if lru.dList == nil {
		return errors.New("LRUCache结构体未初始化")
	}
	if pElement, ok := lru.cacheMap[k]; ok {
		lru.dList.MoveToFront(pElement)
		pElement.Value.(*CacheNode).value = v
		return nil
	}
	newElement := lru.dList.PushFront(&CacheNode{k, v})
	lru.cacheMap[k] = newElement
	if lru.dList.Len() > lru.Capacity {
		//移掉最后一个
		lastElement := lru.dList.Back()
		if lastElement == nil {
			return nil
		}
		cacheNode := lastElement.Value.(*CacheNode)
		delete(lru.cacheMap, cacheNode)
		lru.dList.Remove(lastElement)
	}
	return nil
}

func (lru *LRUCache) Get(k interface{}) (v interface{}, ret bool, err error) {
	if lru.cacheMap == nil {
		return v, false, errors.New("LRUCache结构体未初始化")
	}
	if pElement, ok := lru.cacheMap[k]; ok {
		lru.dList.MoveToFront(pElement)
		return pElement.Value.(*CacheNode).value, true, nil
	}
	return v, false, nil
}

func (lru *LRUCache) Remove(k interface{}) bool {
	if lru.cacheMap == nil {
		return false
	}
	if pElement, ok := lru.cacheMap[k]; ok {
		cacheNode := pElement.Value.(*CacheNode)
		delete(lru.cacheMap, cacheNode.key)
		lru.dList.Remove(pElement)
		return true
	}
	return false

}
