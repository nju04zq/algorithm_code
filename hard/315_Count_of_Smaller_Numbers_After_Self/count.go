package main

import (
	"fmt"
	"math/rand"
	"time"
)

func mergesort(nums, idxs, counts, temp []int, start, end int) {
	if start+1 == end {
		return
	}
	mid := start + (end-start)/2
	mergesort(nums, idxs, counts, temp, start, mid)
	mergesort(nums, idxs, counts, temp, mid, end)
	left, right, k, rightCnt := start, mid, 0, 0
	//fmt.Println(nums, idxs[start:mid], idxs[mid:end], start, end, counts)
	for left < mid || right < end {
		if right >= end ||
			(left < mid && nums[idxs[left]] <= nums[idxs[right]]) {
			temp[k] = idxs[left]
			counts[idxs[left]] += rightCnt
			left++
		} else if left >= mid ||
			(nums[idxs[right]] < nums[idxs[left]]) {
			temp[k] = idxs[right]
			right++
			rightCnt++
		}
		k++
	}
	for i, k := start, 0; i < end; i++ {
		idxs[i] = temp[k]
		k++
	}
	//fmt.Println(start, end, counts, idxs)
}

func countSmaller(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}
	idxs := make([]int, n)
	for i := 0; i < n; i++ {
		idxs[i] = i
	}
	counts := make([]int, n)
	temp := make([]int, n)
	mergesort(nums, idxs, counts, temp, 0, n)
	return counts
}

func MakeRandInt() int {
	maxNum := 40
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int() % maxNum
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

func bf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		cnt := 0
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[i] {
				cnt++
			}
		}
		res[i] = cnt
	}
	return res
}

func testCount() {
	nums := MakeRandArray()
	ans := bf(nums)
	res := countSmaller(nums)
	for i := 0; i < len(ans); i++ {
		if ans[i] != res[i] {
			panic(fmt.Sprintf("%v, get %v, ans %v", nums, res, ans))
		}
	}
}

func main() {
	testCount()
}
