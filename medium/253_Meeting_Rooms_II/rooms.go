package main

import "fmt"
import "sort"
import "math/rand"
import "time"

// Given an array of meeting time intervals consisting of start and end times [[s1,e1],[s2,e2],...] (si < ei), find the minimum number of conference rooms required.
//
// For example,
// Given [[0, 30],[5, 10],[15, 20]],
// return 2.

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func meetingRooms1(a [][]int) int {
	n := len(a)
	if n == 0 {
		return 0
	}
	begins := make([]int, n)
	ends := make([]int, n)
	for i, interval := range a {
		begins[i] = interval[0]
		ends[i] = interval[1]
	}
	sort.Ints(begins)
	sort.Ints(ends)
	maxN, rooms := 0, 0
	i, j := 0, 0
	for i < n {
		if begins[i] < ends[j] {
			rooms++
			maxN = max(maxN, rooms)
			i++
		} else {
			rooms--
			j++
		}
	}
	return maxN
}

type MinHeap struct {
	buf []int
}

func (h *MinHeap) Init() *MinHeap {
	h.buf = make([]int, 0)
	return h
}

func (h *MinHeap) Size() int {
	return len(h.buf)
}

func (h *MinHeap) parent(child int) int {
	return (child - 1) / 2
}

func (h *MinHeap) lchild(parent int) int {
	return 2*parent + 1
}

func (h *MinHeap) rchild(parent int) int {
	return 2*parent + 2
}

func (h *MinHeap) Add(val int) {
	h.buf = append(h.buf, val)
	i := len(h.buf) - 1
	for i > 0 {
		parent := h.parent(i)
		if h.buf[parent] < h.buf[i] {
			break
		}
		h.buf[parent], h.buf[i] = h.buf[i], h.buf[parent]
		i = parent
	}
}

func (h *MinHeap) Peak() int {
	if len(h.buf) == 0 {
		return -1
	} else {
		return h.buf[0]
	}
}

func (h *MinHeap) Pop() int {
	if len(h.buf) == 0 {
		return -1
	}
	val := h.buf[0]
	h.buf[0] = h.buf[len(h.buf)-1]
	h.buf = h.buf[:len(h.buf)-1]
	i, n := 0, len(h.buf)
	for i < len(h.buf) {
		smallest := i
		lchild := h.lchild(i)
		if lchild < n && h.buf[lchild] < h.buf[smallest] {
			smallest = lchild
		}
		rchild := h.rchild(i)
		if rchild < n && h.buf[rchild] < h.buf[smallest] {
			smallest = rchild
		}
		if smallest == i {
			break
		}
		h.buf[smallest], h.buf[i] = h.buf[i], h.buf[smallest]
		i = smallest
	}
	return val
}

func meetingRooms2(a [][]int) int {
	minHeap := new(MinHeap).Init()
	sort.Slice(a, func(i, j int) bool {
		if a[i][0] < a[j][0] {
			return true
		} else {
			return false
		}
	})
	for _, interval := range a {
		if minHeap.Size() == 0 {
			minHeap.Add(interval[1])
			continue
		}
		end := minHeap.Peak()
		if interval[0] >= end {
			minHeap.Pop()
		}
		minHeap.Add(interval[1])
	}
	return minHeap.Size()
}

func MakeRandInt(minNum, maxNum int) int {
	d := maxNum - minNum
	r := rand.New(randSrc)
	return r.Int()%d + minNum
}

func MakeRandInterval() []int {
	begin := MakeRandInt(1, 100)
	end := MakeRandInt(begin+1, 200)
	return []int{begin, end}
}

func MakeRandIntervals() [][]int {
	size := MakeRandInt(1, 100)
	intervals := make([][]int, size)
	for i, _ := range intervals {
		intervals[i] = MakeRandInterval()
	}
	return intervals
}

func testRooms() {
	intervals := MakeRandIntervals()
	res1 := meetingRooms1(intervals)
	res2 := meetingRooms2(intervals)
	if res1 != res2 {
		panic(fmt.Sprintf("%v, get %d, %d\n", intervals, res1, res2))
	}
}

var randSrc = rand.NewSource(time.Now().UnixNano())

func main() {
	for i := 0; i < 10000; i++ {
		testRooms()
	}
}
