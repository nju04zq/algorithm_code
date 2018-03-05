package main

import "fmt"
import "sort"
import "math/rand"
import "time"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	start, end := -1, -2
	minNum, maxNum := nums[n-1], nums[0]
	for i := 1; i < n; i++ {
		maxNum = max(maxNum, nums[i])
		minNum = min(minNum, nums[n-i-1])
		if nums[i] != maxNum {
			end = i
		}
		if nums[n-i-1] != minNum {
			start = n - i - 1
		}
	}
	return end - start + 1
}

func bf(nums []int) int {
	a := make([]int, len(nums))
	for i := 0; i < len(a); i++ {
		a[i] = nums[i]
	}
	sort.Ints(a)
	var i int
	for i = 0; i < len(a); i++ {
		if a[i] != nums[i] {
			break
		}
	}
	start := i
	for i = len(a) - 1; i >= 0; i-- {
		if a[i] != nums[i] {
			break
		}
	}
	end := i
	if start >= end {
		return 0
	} else {
		return end - start + 1
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
	a := MakeRandArray()
	res := findUnsortedSubarray(a)
	ans := bf(a)
	if res != ans {
		panic(fmt.Sprintf("%v, get %d, expect %d\n", a, res, ans))
	}
}

func main() {
	for i := 0; i < 10000; i++ {
		testFind()
	}
}
