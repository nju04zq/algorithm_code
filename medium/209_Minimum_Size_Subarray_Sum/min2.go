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

func binarySearch(sums []int, key int) int {
	low, high := 0, len(sums)-1
	for low < high {
		mid := low + (high-low)/2
		if sums[mid] < key {
			low = mid + 1
		} else {
			high = mid
		}
	}
	if sums[low] >= key {
		return low
	} else {
		return -1
	}
}

func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	sums := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	minLen := math.MaxInt32
	for i := 0; i < n; i++ {
		j := binarySearch(sums, sums[i]+s)
		if j == -1 {
			break
		}
		minLen = min(minLen, j-i)
	}
	if minLen == math.MaxInt32 {
		return 0
	} else {
		return minLen
	}
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
