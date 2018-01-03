package main

import (
	"fmt"
	"math/rand"
	"time"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	tail := make([]int, n)
	size := 0
	for _, num := range nums {
		low, high := 0, size
		for low < high {
			mid := low + (high-low)/2
			if tail[mid] < num {
				low = mid + 1
			} else {
				high = mid
			}
		}
		if low == size {
			tail[size] = num
			size++
		} else {
			tail[low] = num
		}
	}
	return size
}

func dfs(nums []int, start, head int) int {
	if start == len(nums) {
		return 0
	}
	r1 := dfs(nums, start+1, head)
	r2 := 0
	if head == -1 || nums[head] < nums[start] {
		r2 = dfs(nums, start+1, start) + 1
	}
	return max(r1, r2)
}

func bf(nums []int) int {
	return dfs(nums, 0, -1)
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

func testLIS() {
	nums := MakeRandArray()
	res := lengthOfLIS(nums)
	ans := bf(nums)
	if res != ans {
		panic(fmt.Errorf("%v, get %d, expect %d", nums, res, ans))
	}
}

func main() {
	for i := 0; i < 10000; i++ {
		fmt.Printf("\r%d", i)
		testLIS()
	}
	fmt.Println()
}
