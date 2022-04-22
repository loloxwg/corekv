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
		lru.data[newitem.key] = lru.list.PushFront(&newitem)
		return storeItem{}, false
	}

	evictItem := lru.list.Back()
	item := evictItem.Value.(*storeItem)

	delete(lru.data, item.key)

	eitem, *item = *item, newitem
	//直接取代原来的item，不用拷贝，直接替换
	//原地替换
	//而不是 delete 掉，再加进去
	//这里的 evict 的那个东西，已经被替换成了要放进去的
	lru.data[item.key] = evictItem
	lru.list.MoveToFront(evictItem)
	return eitem, true
}

func (lru *windowLRU) get(v *list.Element) {
	//implement me here!!!
	lru.list.MoveToFront(v)

}
