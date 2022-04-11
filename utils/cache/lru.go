package cache

import "container/list"

type windowLRU struct {
	data map[uint64]*list.Element
	cap  int
	list *list.List
}

type storeItem struct {
	stage    int
	key      uint64
	conflict uint64
	value    interface{}
}

func newWindowLRU(size int, data map[uint64]*list.Element) *windowLRU {
	return &windowLRU{
		data: data,
		cap:  size,
		list: list.New(),
	}
}

func (lru *windowLRU) add(newitem storeItem) (eitem storeItem, evicted bool) {
	//implement me here!!!
	if lru.list.Len() < lru.cap {
		eitem = newitem
		lru.data[eitem.key] = lru.list.PushFront(eitem)
		return
	}
	//要淘汰的item 在list中的位置是链表尾部
	evictItem := lru.list.Back()
	item := evictItem.Value.(*storeItem)

	delete(lru.data, item.key)
	//lru.list.Remove(evictedItem)
	eitem, *item = *item, newitem
	//这里实现的和普通的lru不同
	lru.data[item.key] = lru.list.PushFront(evictItem)
	return eitem, true
}

func (lru *windowLRU) get(v *list.Element) {
	//implement me here!!!
	lru.list.MoveToFront(v)

}
