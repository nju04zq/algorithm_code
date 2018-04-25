// Given an array nums and a target value k, find the maximum length of a subarray that sums to k. If there isn't one, return 0 instead.
//
// Example 1:
// Given nums = [1, -1, 5, -2, 3], k = 3,
// return 4. (because the subarray [1, -1, 5, -2] sums to 3 and is the longest)
//
// Example 2:
// Given nums = [-2, -1, 2, 1], k = 1,
// return 2. (because the subarray [-1, 2] sums to 1 and is the longest)
//
// Follow Up:
// Can you do it in O(n) time?

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func maxSubarrayLen(a []int, k int) int {
	n := len(a)
	tbl := make(map[int]int)
	sum, maxLen := 0, 0
	tbl[sum] = 0
	for i := 0; i < n; i++ {
		sum += a[i]
		target := sum - k
		if j, ok := tbl[target]; ok {
			maxLen = max(maxLen, i+1-j)
		}
		if _, ok := tbl[sum]; !ok {
			tbl[sum] = i + 1
		}
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func bf(a []int, k int) int {
	n := len(a)
	maxLen := 0
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += a[j]
			if sum == k {
				maxLen = max(maxLen, j-i+1)
			}
		}
	}
	return maxLen
}

func MakeRandInt() int {
	maxNum := 20
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int()%(maxNum*2) - maxNum
}

func MakeRandArray() []int {
	maxLen, maxElement := 10, 20
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len := r.Int() % maxLen
	a := make([]int, len)
	for i := 0; i < len; i++ {
		a[i] = r.Int()%(maxElement*2) - maxElement
	}
	return a
}

func test() {
	a := MakeRandArray()
	k := MakeRandInt()
	ans := bf(a, k)
	res := maxSubarrayLen(a, k)
	if res != ans {
		panic(fmt.Errorf("Fail on %v, %d, get %d, expect %d", a, k, res, ans))
	}
}

func main() {
	for i := 0; i < 10000; i++ {
		test()
	}
}
