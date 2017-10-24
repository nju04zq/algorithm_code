package main

import "fmt"
import (
	"math/rand"
	"sort"
	"time"
)

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); {
		if nums[i] <= 0 || nums[i] > len(nums) {
			i++
			continue
		}
		j := nums[i] - 1
		if j <= i || nums[i] == nums[j] {
			nums[j] = nums[i]
			i++
		} else {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	for i, num := range nums {
		if num != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
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

func firstBF(nums1 []int) int {
	nums := make([]int, len(nums1))
	copy(nums, nums1)
	sort.Ints(nums)
	missing := 1
	for i, num := range nums {
		if num <= 0 {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if num != missing {
			return missing
		} else {
			missing++
		}
	}
	return missing
}

func testFirstMissing() {
	nums := MakeRandArray()
	ans := firstBF(nums)
	res := firstMissingPositive(nums)
	if res != ans {
		panic(fmt.Errorf("nums %v, get %d, should %d\n", nums, res, ans))
	}
}

func main() {
	for i := 0; i < 100000; i++ {
		testFirstMissing()
	}
}
