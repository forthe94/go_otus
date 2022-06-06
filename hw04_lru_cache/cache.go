package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type cacheItem struct {
	key   Key
	value interface{}
}

func newCacheItem(key Key, value interface{}) *cacheItem {
	newItem := new(cacheItem)
	newItem.key = key
	newItem.value = value
	return newItem
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (cache *lruCache) Set(key Key, value interface{}) bool {
	res := true

	if item, ok := cache.items[key]; !ok {
		newItem := cache.queue.PushFront(newCacheItem(key, value))
		cache.items[key] = newItem
		res = false
		if cache.queue.Len() > cache.capacity {
			toRemove := cache.queue.Back()
			cache.queue.Remove(toRemove)
			delete(cache.items, toRemove.Value.(*cacheItem).key)
		}
	} else {
		item.Value.(*cacheItem).value = value
		cache.queue.MoveToFront(item)
	}
	return res
}

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	if item := cache.items[key]; item != nil {
		cache.queue.MoveToFront(item)
		return item.Value.(*cacheItem).value, true
	}
	return nil, false
}

func (cache *lruCache) Clear() {
	cur := cache.queue.Front()
	for cur != nil {
		delete(cache.items, cur.Value.(*cacheItem).key)
		cache.queue.Remove(cur)
		cur = cache.queue.Front()
	}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
