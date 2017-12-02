package main

import "fmt"

type dequeEntry struct {
	val        int
	prev, next *dequeEntry
}

type deque struct {
	head dequeEntry
}

func (q *deque) init() *deque {
	q.head.prev = &q.head
	q.head.next = &q.head
	return q
}

func (q *deque) dump(nums []int) {
	if q.empty() {
		fmt.Println("<Empty>")
	}
	for p := q.head.next; p != &q.head; p = p.next {
		fmt.Printf("%d -> ", p.val)
	}
	fmt.Println()
}

func (q *deque) empty() bool {
	if q.head.next == &q.head {
		return true
	} else {
		return false
	}
}

func (q *deque) poll() int {
	if q.empty() {
		return -1
	} else {
		entry := q.head.next
		entry.next.prev = &q.head
		q.head.next = entry.next
		return entry.val
	}
}

func (q *deque) offer(i int) {
	entry := &dequeEntry{val: i}
	entry.prev = q.head.prev
	entry.next = &q.head
	q.head.prev.next = entry
	q.head.prev = entry
}

func (q *deque) pollLast() int {
	if q.empty() {
		return -1
	} else {
		entry := q.head.prev
		entry.prev.next = &q.head
		q.head.prev = entry.prev
		return entry.val
	}
}

func (q *deque) peek() int {
	if q.empty() {
		return -1
	} else {
		return q.head.next.val
	}
}

func (q *deque) peekLast() int {
	if q.empty() {
		return -1
	} else {
		return q.head.prev.val
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	res := make([]int, len(nums)-k+1)
	j := 0
	q := new(deque).init()
	for i := 0; i < len(nums); i++ {
		//fmt.Printf("=====%d=====\n", i)
		//q.dump(nums)
		for !q.empty() && q.peek() < i-k+1 {
			q.poll()
		}
		//q.dump(nums)
		for !q.empty() && nums[q.peekLast()] <= nums[i] {
			q.pollLast()
		}
		q.offer(i)
		//q.dump(nums)
		if i >= k-1 {
			res[j] = nums[q.peek()]
			j++
		}
	}
	return res
}

func testSliding(nums []int, k int) {
	res := maxSlidingWindow(nums, k)
	fmt.Printf("%v, k %d, get %v\n", nums, k, res)
}

func main() {
	testSliding([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
}
