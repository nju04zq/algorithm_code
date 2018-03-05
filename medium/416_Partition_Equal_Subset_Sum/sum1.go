package main

import (
	"fmt"
	"math/rand"
	"time"
)

func subSum(nums []int, m int) bool {
	n := len(nums)
	dp := make([]int, m+1)
	dp[0] = 1
	for i := 0; i < n; i++ {
		for j := m; j >= 0; j-- {
			if j-nums[i] >= 0 {
				dp[j] += dp[j-nums[i]]
			}
		}
	}
	if dp[m] != 0 {
		return true
	} else {
		return false
	}
}

func canPartition(nums []int) bool {
	n := len(nums)
	if n <= 1 {
		return false
	}
	total := 0
	for i := 0; i < n; i++ {
		total += nums[i]
	}
	if total%2 != 0 {
		return false
	}
	return subSum(nums, total/2)
}

func dfs(nums []int, i int, path []int, sum, k int) bool {
	if sum == k && len(path) > 0 {
		return true
	}
	if i == len(nums) {
		return false
	}
	if dfs(nums, i+1, path, sum, k) {
		return true
	}
	path = append(path, nums[i])
	sum += nums[i]
	if dfs(nums, i+1, path, sum, k) {
		return true
	}
	return false
}

func bf(nums []int, k int) bool {
	path := make([]int, 0)
	return dfs(nums, 0, path, 0, k)
}

func MakeRandArray() []int {
	maxLen, maxElement := 30, 100
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len := r.Int()%maxLen + 4
	a := make([]int, len)
	for i := 0; i < len; i++ {
		a[i] = r.Int() % maxElement
	}
	return a
}

func testSub() {
	a := MakeRandArray()
	sum := 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	k := sum / 2
	res := subSum(a, k)
	ans := bf(a, k)
	if res != ans {
		panic(fmt.Sprintf("%v, %d, get %t, expect %t", a, k, res, ans))
	}
}

func testPartition(nums []int) {
	fmt.Printf("%v, get %t\n", nums, canPartition(nums))
}

func main() {
	testPartition([]int{1, 5, 11, 5})
	a := []int{66, 90, 7, 6, 32, 16, 2, 78, 69, 88, 85, 26, 3, 9, 58, 65, 30, 96, 11, 31, 99, 49, 63, 83, 79, 97, 20, 64, 81, 80, 25, 69, 9, 75, 23, 70, 26, 71, 25, 54, 1, 40, 41, 82, 32, 10, 26, 33, 50, 71, 5, 91, 59, 96, 9, 15, 46, 70, 26, 32, 49, 35, 80, 21, 34, 95, 51, 66, 17, 71, 28, 88, 46, 21, 31, 71, 42, 2, 98, 96, 40, 65, 92, 43, 68, 14, 98, 38, 13, 77, 14, 13, 60, 79, 52, 46, 9, 13, 25, 8}
	testPartition(a)
	//for i := 0; i < 10000; i++ {
	//	fmt.Printf("\r%d", i)
	//	testSub()
	//}
	//fmt.Println()
}
