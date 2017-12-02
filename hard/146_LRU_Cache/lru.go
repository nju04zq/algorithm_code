package main

import "fmt"

type LRUEntry struct {
	key, val   int
	prev, next *LRUEntry
}

type LRUCache struct {
	tbl       map[int]*LRUEntry
	head      *LRUEntry
	size, cap int
}

func Constructor(capacity int) LRUCache {
	cache := &LRUCache{
		tbl:  make(map[int]*LRUEntry),
		head: new(LRUEntry),
		size: 0,
		cap:  capacity,
	}
	cache.head.prev = cache.head
	cache.head.next = cache.head
	return *cache
}

func (cache *LRUCache) add(entry *LRUEntry) {
	head := cache.head
	entry.prev = head
	entry.next = head.next
	head.next.prev = entry
	head.next = entry
}

func (cache *LRUCache) removeFromList(entry *LRUEntry) {
	entry.prev.next = entry.next
	entry.next.prev = entry.prev
}

func (cache *LRUCache) update(entry *LRUEntry) {
	cache.removeFromList(entry)
	cache.add(entry)
}

func (cache *LRUCache) remove(entry *LRUEntry) {
	cache.removeFromList(entry)
	delete(cache.tbl, entry.key)
}

func (cache *LRUCache) Get(key int) int {
	entry, ok := cache.tbl[key]
	if !ok {
		return -1
	}
	cache.update(entry)
	return entry.val
}

func (cache *LRUCache) Put(key int, value int) {
	if entry, ok := cache.tbl[key]; ok {
		entry.val = value
		cache.update(entry)
		return
	}
	if cache.size >= cache.cap {
		cache.remove(cache.head.prev)
	}
	entry := &LRUEntry{key: key, val: value}
	cache.tbl[key] = entry
	cache.add(entry)
	cache.size++
}

func put(cache *LRUCache, key, val int) {
	fmt.Printf("Put %d, %d\n", key, val)
	cache.Put(key, val)
}

func get(cache *LRUCache, key int) {
	fmt.Printf("Get %d, %d\n", key, cache.Get(key))
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	c := Constructor(2)
	cache := &c
	put(cache, 1, 1)
	put(cache, 2, 2)
	get(cache, 1)
	get(cache, 1)
	get(cache, 2)
	put(cache, 3, 3)
	get(cache, 1)
	get(cache, 2)

	fmt.Println("======")

	c = Constructor(2)
	cache = &c
	put(cache, 1, 1)
	put(cache, 2, 2)
	put(cache, 1, 2)
	put(cache, 1, 1)
	get(cache, 2)
	put(cache, 3, 3)
	get(cache, 1)
	get(cache, 2)
}
