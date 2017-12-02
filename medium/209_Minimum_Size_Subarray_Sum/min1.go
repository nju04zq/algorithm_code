package main

import "fmt"
import "math"
import "time"
import "math/rand"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func minSubArrayLen(s int, nums []int) int {
	var sum, minLen int
	minLen = math.MaxInt32
	for i, j := 0, 0; i < len(nums); i++ {
		sum += nums[i]
		for sum >= s {
			minLen = min(minLen, i-j+1)
			sum -= nums[j]
			j++
		}
	}
	if minLen == math.MaxInt32 {
		minLen = 0
	}
	return minLen
}

func bf(s int, nums []int) int {
	minLen := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum >= s {
				minLen = min(minLen, j-i+1)
				break
			}
		}
	}
	if minLen == math.MaxInt32 {
		minLen = 0
	}
	return minLen
}

func MakeRandInt() int {
	maxNum := 40
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int()%maxNum + 1
}

func MakeRandArray() []int {
	maxLen, maxElement := 10, 20
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len := r.Int() % maxLen
	a := make([]int, len)
	for i := 0; i < len; i++ {
		a[i] = r.Int()%maxElement + 1
	}
	return a
}

func testMinSub() {
	s := MakeRandInt()
	nums := MakeRandArray()
	res := minSubArrayLen(s, nums)
	ans := bf(s, nums)
	if res != ans {
		panic(fmt.Sprintf("%v, %d, get %d, expect %d\n", nums, s, res, ans))
	}
}

func main() {
	for i := 0; i < 10000; i++ {
		testMinSub()
	}
}
