package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const (
	MIN_HEAP = iota
	MAX_HEAP
)

type Heap struct {
	kind int
	buf  []int
}

func (h *Heap) init(kind int) *Heap {
	h.kind = kind
	h.buf = make([]int, 0)
	return h
}

func (h *Heap) size() int {
	return len(h.buf)
}

func (h *Heap) peek() int {
	if len(h.buf) == 0 {
		return -1
	} else {
		return h.buf[0]
	}
}

func (h *Heap) add(num int) {
	h.buf = append(h.buf, num)
	i := len(h.buf) - 1
	for i > 0 {
		parent := h.parent(i)
		if h.priority(parent, i) {
			break
		}
		h.buf[parent], h.buf[i] = h.buf[i], h.buf[parent]
		i = parent
	}
}

func (h *Heap) remove() int {
	if len(h.buf) == 0 {
		return -1
	}
	size := len(h.buf)
	val := h.buf[0]
	h.buf[0] = h.buf[size-1]
	h.buf = h.buf[:size-1]
	h.heapify(0)
	return val
}

func (h *Heap) heapify(i int) {
	if len(h.buf) == 0 {
		return
	}
	n := len(h.buf)
	for {
		lchild := h.lchild(i)
		rchild := h.rchild(i)
		highest := i
		if lchild < n && h.priority(lchild, highest) {
			highest = lchild
		}
		if rchild < n && h.priority(rchild, highest) {
			highest = rchild
		}
		if highest == i {
			break
		} else {
			h.buf[i], h.buf[highest] = h.buf[highest], h.buf[i]
			i = highest
		}
	}
}

func (h *Heap) parent(i int) int {
	return (i - 1) / 2
}

func (h *Heap) lchild(i int) int {
	return 2*i + 1
}

func (h *Heap) rchild(i int) int {
	return 2*i + 2
}

func (h *Heap) priority(i, j int) bool {
	if h.kind == MIN_HEAP {
		if h.buf[i] < h.buf[j] {
			return true
		} else {
			return false
		}
	} else {
		if h.buf[i] > h.buf[j] {
			return true
		} else {
			return false
		}
	}
}

type MedianFinder struct {
	maxheap *Heap
	minheap *Heap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	f := &MedianFinder{}
	f.minheap = new(Heap).init(MIN_HEAP)
	f.maxheap = new(Heap).init(MAX_HEAP)
	return *f
}

func (f *MedianFinder) AddNum(num int) {
	if f.maxheap.size() == 0 {
		f.maxheap.add(num)
	} else if f.maxheap.size() == f.minheap.size() {
		if num <= f.minheap.peek() {
			f.maxheap.add(num)
		} else {
			val := f.minheap.remove()
			f.maxheap.add(val)
			f.minheap.add(num)
		}
	} else {
		if num <= f.maxheap.peek() {
			val := f.maxheap.remove()
			f.maxheap.add(num)
			f.minheap.add(val)
		} else {
			f.minheap.add(num)
		}
	}
}

func (f *MedianFinder) FindMedian() float64 {
	if f.maxheap.size() == 0 {
		return 0
	} else if f.maxheap.size() == f.minheap.size() {
		a := f.maxheap.peek()
		b := f.minheap.peek()
		return (float64(a) + float64(b)) / 2
	} else {
		return float64(f.maxheap.peek())
	}
}

func findMedian(nums []int) float64 {
	nums1 := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		nums1[i] = nums[i]
	}
	sort.Ints(nums1)
	n := len(nums1)
	if n%2 != 0 {
		return float64(nums1[(n-1)/2])
	} else {
		return (float64(nums1[n/2]) + float64(nums1[n/2-1])) / 2
	}
}

func MakeRandArray() []int {
	maxLen, maxElement := 10, 20
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len := r.Int() % maxLen
	a := make([]int, len)
	for i := 0; i < len; i++ {
		a[i] = r.Int() % maxElement
	}
	return a
}

func testFind() {
	find := Constructor()
	f := &find
	nums := MakeRandArray()
	for i := 0; i < len(nums); i++ {
		f.AddNum(nums[i])
		res := f.FindMedian()
		ans := findMedian(nums[:i+1])
		if res != ans {
			panic(fmt.Errorf("Fail on %v, i %d, get %v, ans %v", nums, i, res, ans))
		}
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

func main() {
	testFind()
}
