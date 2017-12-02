package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func maxTriangle(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}
	sort.Ints(nums)
	for i := n - 1; i >= 2; i-- {
		if nums[i] < nums[i-1]+nums[i-2] {
			return nums[i] + nums[i-1] + nums[i-2]
		}
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func canForm(a, b, c int) bool {
	if a+b > c && a+c > b && b+c > a {
		return true
	} else {
		return false
	}
}

func bf(nums []int) int {
	n := len(nums)
	longest := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if canForm(nums[i], nums[j], nums[k]) {
					longest = max(longest, nums[i]+nums[j]+nums[k])
				}
			}
		}
	}
	return longest
}

func MakeRandInt() int {
	maxNum := 40
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int() % maxNum
}

func MakeRandArray() []int {
	maxLen, maxElement := 100, 100
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len := r.Int() % maxLen
	a := make([]int, len)
	for i := 0; i < len; i++ {
		a[i] = r.Int()%maxElement + 1
	}
	return a
}

func testSolution() {
	nums := MakeRandArray()
	ans := bf(nums)
	res := maxTriangle(nums)
	if res != ans {
		panic(fmt.Errorf("%v, get %d, expect %d", nums, res, ans))
	}
}

func main() {
	for i := 0; i < 100000; i++ {
		fmt.Printf("\r%d", i)
		testSolution()
	}
	fmt.Println()
}
