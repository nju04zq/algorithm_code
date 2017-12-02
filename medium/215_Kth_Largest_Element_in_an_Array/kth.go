package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func split(nums []int, start, end int) int {
	i, j := start+1, start+1
	for ; i <= end; i++ {
		if nums[i] < nums[start] {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	nums[start], nums[j-1] = nums[j-1], nums[start]
	return j - 1
}

func findInternal(nums []int, start, end int, k int) int {
	mid := split(nums, start, end)
	//fmt.Println("start", start, "end", end, "k", k, "mid", mid)
	//fmt.Println(nums)
	if start == end {
		return start
	}
	if mid == start+k {
		return mid
	} else if mid > start+k {
		return findInternal(nums, start, mid-1, k)
	} else {
		return findInternal(nums, mid+1, end, k-mid+start-1)
	}
}

func findKthLargest(nums []int, k int) int {
	k = len(nums) - k
	i := findInternal(nums, 0, len(nums)-1, k)
	return nums[i]
}

func copyNums(nums []int) []int {
	nums1 := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		nums1[i] = nums[i]
	}
	return nums1
}

func MakeRandArray() []int {
	maxLen, maxElement := 10, 20
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len := r.Int()%maxLen + 1
	a := make([]int, len)
	for i := 0; i < len; i++ {
		a[i] = r.Int() % maxElement
	}
	return a
}

func testFind() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	nums := MakeRandArray()
	k := r.Int()%len(nums) + 1
	nums1 := copyNums(nums)
	nums2 := copyNums(nums)
	res := findKthLargest(nums1, k)
	sort.Slice(nums2, func(i, j int) bool {
		if nums2[i] > nums2[j] {
			return true
		} else {
			return false
		}
	})
	ans := nums2[k-1]
	if res != ans {
		panic(fmt.Sprintf("%v, k %d, get %d, expect %d\n", nums, k, res, ans))
	}
}

func main() {
	for i := 0; i < 10000; i++ {
		testFind()
	}
}
