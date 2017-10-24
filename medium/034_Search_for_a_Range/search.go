package main

import "fmt"
import (
	"math/rand"
	"sort"
	"time"
)

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func searchForLow(nums []int, target int) int {
	idx := -1
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			idx = mid
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return idx
}

func searchForHigh(nums []int, target int) int {
	idx := -1
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			idx = mid
			low = mid + 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return idx
}

func searchRange(nums []int, target int) []int {
	low := searchForLow(nums, target)
	high := searchForHigh(nums, target)
	return []int{low, high}
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

func MakeRandSortedArray() []int {
	a := MakeRandArray()
	sort.Ints(a)
	return a
}

func searchBF(nums []int, target int) []int {
	low, high := -1, -1
	for i, num := range nums {
		if num == target {
			if low == -1 {
				low = i
			}
			high = i
		}
	}
	return []int{low, high}
}

func testSearchRange() {
	nums := MakeRandSortedArray()
	target := MakeRandInt()
	ans := searchBF(nums, target)
	res := searchRange(nums, target)
	if ans[0] != res[0] || ans[1] != res[1] {
		panic(fmt.Errorf("nums %v, target %d, get %v, should %v",
			nums, target, res, ans))
	}
}

func main() {
	for i := 0; i < 100000; i++ {
		testSearchRange()
	}
}
