package hw04

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool // Добавить значение в кэш по ключу.
	Get(key Key) (interface{}, bool)     // Получить значение из кэша по ключу.
	Clear()                              // Очистить кэш.
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

func (lru lruCache) Set(key Key, value interface{}) bool {
	listItem, isCached := lru.items[key]
	if isCached {
		listItem.Value = cacheItem{key: key, value: value}
		lru.queue.MoveToFront(listItem)
	} else {
		if lru.queue.Len() == lru.capacity {
			delItem := lru.queue.Back()
			lru.queue.Remove(lru.queue.Back())
			delete(lru.items, delItem.Value.(cacheItem).key)
		}
		newItem := cacheItem{key: key, value: value}
		lru.queue.PushFront(newItem)
		lru.items[key] = lru.queue.Front()
	}
	return isCached
}

func (lru lruCache) Get(key Key) (interface{}, bool) {
	listItem, isCached := lru.items[key]
	if isCached {
		lru.queue.MoveToFront(listItem)
		return listItem.Value.(cacheItem).value, true
	}
	return nil, false
}

func (lru lruCache) Clear() {

}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
