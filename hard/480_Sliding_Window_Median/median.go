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
	tbl  map[int]int
	size int
}

func (h *Heap) init(kind int) *Heap {
	h.kind = kind
	h.buf = make([]int, 0)
	h.tbl = make(map[int]int)
	return h
}

func (h *Heap) dump() {
	fmt.Printf("heap vals: %v, size %d, tbl %v\n", h.buf, h.size, h.tbl)
}

func (h *Heap) getSize() int {
	return h.size
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
	if _, ok := h.tbl[num]; ok {
		h.tbl[num]++
	} else {
		h.tbl[num] = 1
	}
	h.size++
}

func (h *Heap) removeInvalidTop() {
	for len(h.buf) > 0 {
		top := h.buf[0]
		if h.tbl[top] > 0 {
			break
		}
		size := len(h.buf)
		h.buf[0] = h.buf[size-1]
		h.buf = h.buf[:size-1]
		h.heapify(0)
	}
}

func (h *Heap) removeTop() int {
	if len(h.buf) == 0 {
		return -1
	}
	size := len(h.buf)
	val := h.buf[0]
	h.buf[0] = h.buf[size-1]
	h.buf = h.buf[:size-1]
	h.heapify(0)
	if h.tbl[val] > 0 {
		h.tbl[val]--
		h.size--
	}
	h.removeInvalidTop()
	return val
}

func (h *Heap) remove(num int) {
	if len(h.buf) == 0 {
		return
	}
	h.tbl[num]--
	h.size--
	h.removeInvalidTop()
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
func Constructor() *MedianFinder {
	f := &MedianFinder{}
	f.minheap = new(Heap).init(MIN_HEAP)
	f.maxheap = new(Heap).init(MAX_HEAP)
	return f
}

func (f *MedianFinder) AddNum(num int) {
	if f.maxheap.getSize() == 0 {
		f.maxheap.add(num)
	} else if f.maxheap.getSize() == f.minheap.getSize() {
		if num <= f.minheap.peek() {
			f.maxheap.add(num)
		} else {
			val := f.minheap.removeTop()
			f.maxheap.add(val)
			f.minheap.add(num)
		}
	} else {
		if num <= f.maxheap.peek() {
			val := f.maxheap.removeTop()
			f.maxheap.add(num)
			f.minheap.add(val)
		} else {
			f.minheap.add(num)
		}
	}
}

func (f *MedianFinder) RemoveNum(num int) {
	if num <= f.maxheap.peek() {
		f.maxheap.remove(num)
	} else {
		f.minheap.remove(num)
	}
	//fmt.Printf("remove %d\n", num)
	//f.Dump()
	if f.maxheap.getSize() == f.minheap.getSize() {
		return
	} else if f.maxheap.getSize() == (f.minheap.getSize() + 1) {
		return
	} else if f.maxheap.getSize() > f.minheap.getSize() {
		a := f.maxheap.removeTop()
		f.minheap.add(a)
	} else {
		a := f.minheap.removeTop()
		f.maxheap.add(a)
		//fmt.Printf("minheap remove %d, maxheap add %d\n", a, a)
		//fmt.Printf("remove %d\n", num)
		//f.Dump()
	}
}

func (f *MedianFinder) FindMedian() float64 {
	if f.maxheap.getSize() == 0 {
		return 0
	} else if f.maxheap.getSize() == f.minheap.getSize() {
		a := f.maxheap.peek()
		b := f.minheap.peek()
		return (float64(a) + float64(b)) / 2
	} else {
		return float64(f.maxheap.peek())
	}
}

func (f *MedianFinder) Dump() {
	fmt.Println("Medians:")
	f.maxheap.dump()
	f.minheap.dump()
}

func medianSlidingWindow(nums []int, k int) []float64 {
	f := Constructor()
	res := make([]float64, 0)
	for i, num := range nums {
		f.AddNum(num)
		if i < k-1 {
			continue
		}
		//fmt.Println("*******************************")
		//fmt.Println("===============================")
		//fmt.Println(nums[i-k+1 : i+1])
		//f.Dump()
		if i >= k {
			f.RemoveNum(nums[i-k])
		}
		//fmt.Println("===============================")
		//fmt.Println(nums[i-k+1 : i+1])
		//f.Dump()
		res = append(res, f.FindMedian())
	}
	return res
}

func bfMedian(nums []int) float64 {
	nums1 := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		nums1[i] = nums[i]
	}
	sort.Ints(nums1)
	n := len(nums1)
	if n%2 != 0 {
		return float64(nums1[n/2])
	} else {
		return (float64(nums1[n/2-1]) + float64(nums1[n/2])) / 2
	}
}

func bf(nums []int, k int) []float64 {
	n := len(nums)
	res := make([]float64, 0)
	for i := 0; i <= n-k; i++ {
		res = append(res, bfMedian(nums[i:i+k]))
	}
	return res
}

func MakeRandInt(i int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int()%i + 1
}

func MakeRandArray() []int {
	maxLen, maxElement := 20, 20
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len := r.Int()%maxLen + 4
	a := make([]int, len)
	for i := 0; i < len; i++ {
		a[i] = r.Int() % maxElement
	}
	return a
}

func testMedian() {
	nums := MakeRandArray()
	k := MakeRandInt(len(nums) / 2)
	res := medianSlidingWindow(nums, k)
	ans := bf(nums, k)
	if len(res) != len(ans) {
		panic(fmt.Errorf("%v, res len %d, ans len %d",
			nums, len(res), len(ans)))
	}
	for i := 0; i < len(res); i++ {
		if res[i] != ans[i] {
			panic(fmt.Errorf("%v, k %d, get %v, expect %v", nums, k, res, ans))
		}
	}
}

func main() {
	for i := 0; i < 10000; i++ {
		fmt.Printf("\r%d", i)
		testMedian()
	}
	fmt.Println()
}
