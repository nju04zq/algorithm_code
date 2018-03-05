package main

import "fmt"
import "math"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

type MinHeap struct {
	nums [][]int
	idxs []int
	buf  []int
	done bool
}

func (h *MinHeap) init(nums [][]int) *MinHeap {
	k := len(nums)
	idxs := make([]int, k)
	buf := make([]int, k)
	for i := 0; i < k; i++ {
		idxs[i], buf[i] = 0, i
	}
	h.nums = nums
	h.idxs = idxs
	h.buf = buf
	for i := k / 2; i >= 0; i-- {
		h.minHeapify(i)
	}
	return h
}

func (h *MinHeap) isDone() bool {
	return h.done
}

func (h *MinHeap) lchild(i int) int {
	return 2*i + 1
}

func (h *MinHeap) rchild(i int) int {
	return 2*i + 2
}

func (h *MinHeap) val(i int) int {
	idx := h.buf[i]
	j := h.idxs[idx]
	return h.nums[idx][j]
}

func (h *MinHeap) minHeapify(i int) {
	n := len(h.buf)
	for i < n {
		least := i
		lchild := h.lchild(i)
		rchild := h.rchild(i)
		if lchild < n && h.val(lchild) < h.val(least) {
			least = lchild
		}
		if rchild < n && h.val(rchild) < h.val(least) {
			least = rchild
		}
		if least == i {
			break
		}
		h.buf[i], h.buf[least] = h.buf[least], h.buf[i]
		i = least
	}
}

func (h *MinHeap) pop() (int, int) {
	if len(h.buf) == 0 {
		return -1, -1
	}
	resIdx, resVal := h.buf[0], h.val(0)
	h.idxs[resIdx]++
	if h.idxs[resIdx] >= len(h.nums[resIdx]) {
		h.done = true
		return resVal, -1
	}
	nextVal := h.val(0)
	h.minHeapify(0)
	return resVal, nextVal
}

func rangeSize(r []int) int {
	return r[1] - r[0] + 1
}

func smaller(r1, r2 []int) bool {
	if len(r2) == 0 {
		return true
	}
	if rangeSize(r1) < rangeSize(r2) {
		return true
	} else if rangeSize(r1) > rangeSize(r2) {
		return false
	} else if r1[0] < r2[0] {
		return true
	} else {
		return false
	}
}

func smallestRange(nums [][]int) []int {
	if len(nums) == 1 {
		return []int{nums[0][0], nums[0][0]}
	}
	curMax := math.MinInt32
	for i := 0; i < len(nums); i++ {
		curMax = max(curMax, nums[i][0])
	}
	heap := new(MinHeap).init(nums)
	minRange := []int{}
	for {
		//fmt.Println(heap.buf, heap.idxs)
		curVal, nextVal := heap.pop()
		r := []int{curVal, curMax}
		if smaller(r, minRange) {
			minRange = r
		}
		//fmt.Printf("curVal %d, nextVal %d, curMax %d, minRange %v\n", curVal, nextVal, curMax, minRange)
		if heap.isDone() {
			break
		}
		curMax = max(curMax, nextVal)
	}
	return minRange
}

func testSmallest(nums [][]int) {
	res := smallestRange(nums)
	fmt.Printf("%v, get %v\n", nums, res)
}

func main() {
	nums := [][]int{
		[]int{4, 10, 15, 24, 26},
		[]int{0, 9, 12, 20},
		[]int{5, 18, 22, 30},
	}
	testSmallest(nums)
}
