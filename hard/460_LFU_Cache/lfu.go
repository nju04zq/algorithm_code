package main

import "fmt"

type LFUEntry struct {
	key, val, ref int
	prev, next    *LFUEntry
	head          *LFUHead
}

type LFUHead struct {
	ref        int
	head       *LFUEntry
	prev, next *LFUHead
}

type LFUCache struct {
	tbl       map[int]*LFUEntry
	heads     *LFUHead
	size, cap int
}

func Constructor(capacity int) LFUCache {
	cache := &LFUCache{
		tbl:   make(map[int]*LFUEntry),
		heads: new(LFUHead),
		cap:   capacity,
	}
	cache.heads.prev, cache.heads.next = cache.heads, cache.heads
	return *cache

}

func (cache *LFUCache) addEntryToList(head *LFUEntry, entry *LFUEntry) {
	entry.prev = head
	entry.next = head.next
	head.next.prev = entry
	head.next = entry
}

func (cache *LFUCache) removeEntryFromList(entry *LFUEntry) {
	entry.next.prev = entry.prev
	entry.prev.next = entry.next
	entry.head = nil
}

func (cache *LFUCache) addHead(prev *LFUHead, ref int) *LFUHead {
	head := &LFUHead{ref: ref}
	head.prev = prev
	head.next = prev.next
	prev.next.prev = head
	prev.next = head
	head.head = new(LFUEntry)
	head.head.prev = head.head
	head.head.next = head.head
	return head
}

func (cache *LFUCache) removeHead(head *LFUHead) {
	head.next.prev = head.prev
	head.prev.next = head.next
}

func (cache *LFUCache) reInsertEntry(entry *LFUEntry) {
	head0, head1 := entry.head, entry.head
	if head0.next == cache.heads || head0.next.ref != entry.ref {
		head1 = cache.addHead(head0, entry.ref)
	} else {
		head1 = head0.next
	}
	cache.removeEntryFromList(entry)
	if head0.head.next == head0.head {
		cache.removeHead(head0)
	}
	cache.addEntryToList(head1.head, entry)
	entry.head = head1
}

func (cache *LFUCache) insertEntry(entry *LFUEntry) {
	head0, head1 := cache.heads.next, cache.heads.next
	if head0 == cache.heads || head0.ref != 1 {
		head1 = cache.addHead(cache.heads, 1)
	} else {
		head1 = head0
	}
	cache.addEntryToList(head1.head, entry)
	entry.head = head1
}

func (cache *LFUCache) touchEntry(entry *LFUEntry) {
	entry.ref++
	cache.reInsertEntry(entry)
}

func (cache *LFUCache) removeOldEntry() {
	head := cache.heads.next
	entry := head.head.prev
	cache.removeEntryFromList(entry)
	if head.head.next == head.head {
		cache.removeHead(head)
	}
	delete(cache.tbl, entry.key)
	cache.size--
}

func (cache *LFUCache) Get(key int) int {
	entry, ok := cache.tbl[key]
	if !ok {
		return -1
	}
	cache.touchEntry(entry)
	return entry.val
}

func (cache *LFUCache) Put(key int, value int) {
	if cache.cap == 0 {
		return
	}
	if entry, ok := cache.tbl[key]; ok {
		entry.val = value
		cache.touchEntry(entry)
		return
	}
	if cache.size >= cache.cap {
		cache.removeOldEntry()
	}
	entry := &LFUEntry{key: key, val: value, ref: 1}
	cache.insertEntry(entry)
	cache.tbl[key] = entry
	cache.size++
}

func put(cache *LFUCache, key, val int) {
	fmt.Println("======")
	fmt.Printf("Put %d, %d\n", key, val)
	cache.Put(key, val)
	dump(cache)
}

func get(cache *LFUCache, key int) {
	fmt.Println("======")
	fmt.Printf("Get %d, %d\n", key, cache.Get(key))
	dump(cache)
}

func dump(cache *LFUCache) {
	fmt.Println("Cache inside:")
	if cache.heads.next == cache.heads {
		fmt.Println("<Empty>")
		return
	}
	for head := cache.heads.next; head != cache.heads; head = head.next {
		fmt.Printf("%d: ", head.ref)
		for entry := head.head.next; entry != head.head; entry = entry.next {
			fmt.Printf("(%d, %d, %d) ", entry.key, entry.val, entry.ref)
		}
		fmt.Println()
	}
}

func main() {
	c := Constructor(2)
	cache := &c
	put(cache, 1, 1)
	put(cache, 1, 2)
	put(cache, 2, 1)
	put(cache, 2, 2)
	get(cache, 1)
	put(cache, 3, 3)
	// [[2],[1,1],[2,2],[1],[3,3],[2],[3],[4,4],[1],[3],[4]]
	// [null,null,null,1,null,-1,3,null,-1,3,4]
	fmt.Println("***************")
	c = Constructor(2)
	put(cache, 1, 1)
	put(cache, 2, 2)
	get(cache, 1) //1
	put(cache, 3, 3)
	get(cache, 2) //-1
	get(cache, 3) //3
	put(cache, 4, 4)
	get(cache, 1) //-1
	get(cache, 3) //3
	get(cache, 4) //4
	fmt.Println("***************")
	c = Constructor(0)
	put(cache, 0, 0)
	get(cache, 0) //-1
}
