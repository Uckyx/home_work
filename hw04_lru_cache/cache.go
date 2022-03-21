package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (lc *lruCache) Set(key Key, value interface{}) bool {
	item := &cacheItem{key: key, value: value}
	if listItem, ok := lc.items[key]; ok {
		listItem.Value = item
		lc.queue.MoveToFront(listItem)

		return true
	}

	if len(lc.items) > lc.queue.Len() {
		//lastElement := lc.queue.Back()
		//lc.queue.Back().Value
		//
		//lc.queue.Remove(lc.queue.Back())
		//delete(lc.items[lastElement.Value.key])

	}

	lc.items[key] = lc.queue.PushFront(item)

	return false
}

func (lc *lruCache) Get(key Key) (interface{}, bool) {
	if listItem, ok := lc.items[key]; ok {
		lc.queue.MoveToFront(listItem)

		return listItem.Value, true
	}

	return nil, false
}

func (lc *lruCache) Clear() {
	lc.queue = NewList()
	for item := range lc.items {
		delete(lc.items, item)
	}
}
