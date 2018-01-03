package main

import "fmt"

type minHeap struct {
	primes []int
	idxs   []int
	buf    []int
	seq    []int
}

func (h *minHeap) init(primes []int) *minHeap {
	h.primes = primes
	h.seq = []int{0, 1}
	n := len(primes)
	h.idxs = make([]int, n)
	h.buf = make([]int, n)
	for i := 0; i < n; i++ {
		h.idxs[i] = 1
	}
	for i := 0; i < n; i++ {
		h.buf[i] = i
	}
	h.initHeap()
	return h
}

func (h *minHeap) lchild(parent int) int {
	return 2*parent + 1
}

func (h *minHeap) rchild(parent int) int {
	return 2*parent + 2
}

func (h *minHeap) initHeap() {
	n := len(h.buf)
	for i := n / 2; i >= 0; i-- {
		h.minHeapify(i)
	}
}

func (h *minHeap) val(i int) int {
	j := h.buf[i]
	//fmt.Println(i, j, h.primes, h.seq, h.idxs)
	return h.primes[j] * h.seq[h.idxs[j]]
}

func (h *minHeap) minHeapify(i int) {
	n := len(h.buf)
	for i < n {
		lchild := h.lchild(i)
		rchild := h.rchild(i)
		minIdx := i
		if lchild < n && h.val(lchild) < h.val(minIdx) {
			minIdx = lchild
		}
		if rchild < n && h.val(rchild) < h.val(minIdx) {
			minIdx = rchild
		}
		if minIdx == i {
			break
		}
		h.buf[i], h.buf[minIdx] = h.buf[minIdx], h.buf[i]
		i = minIdx
	}
}

func (h *minHeap) next() {
	res := h.val(0)
	h.seq = append(h.seq, res)
	for {
		h.idxs[h.buf[0]]++
		h.minHeapify(0)
		if h.val(0) != res {
			break
		}
	}
}

func (h *minHeap) dump() {
	fmt.Println("==========")
	fmt.Println("h.primes", h.primes)
	fmt.Println("h.seq", h.seq)
	fmt.Println("h.idxs", h.idxs)
	fmt.Println("h.buf", h.buf)
	for i := 0; i < len(h.buf); i++ {
		fmt.Printf("%d ", h.val(h.buf[i]))
	}
	fmt.Println()
	fmt.Println("==========")
}

func (h *minHeap) getLast() int {
	return h.seq[len(h.seq)-1]
}

func nthSuperUglyNumber(n int, primes []int) int {
	if n == 1 {
		return 1
	}
	h := new(minHeap).init(primes)
	for i := 2; i <= n; i++ {
		h.next()
	}
	return h.getLast()
}

func testSuper(n int, primes []int) {
	fmt.Printf("n %d, %v, get %d\n", n, primes, nthSuperUglyNumber(n, primes))
}

func main() {
	primes := []int{2, 7, 13, 19}
	for i := 1; i <= 12; i++ {
		testSuper(i, primes)
	}
}
