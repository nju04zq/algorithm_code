package main

import "fmt"

func makeCopy(src []int) []int {
	dst := make([]int, len(src))
	copy(dst, src)
	return dst
}

func subsetsInternal(nums, path []int, start, maxLen int, ans [][]int) [][]int {
	if len(path) == maxLen {
		ans = append(ans, makeCopy(path))
	}
	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		ans = subsetsInternal(nums, path, i+1, maxLen, ans)
		path = path[:len(path)-1]
	}
	return ans
}

func subsets(nums []int) [][]int {
	path := make([]int, 0, len(nums))
	ans := make([][]int, 0)
	for i := 0; i <= len(nums); i++ {
		ans = subsetsInternal(nums, path, 0, i, ans)
	}
	return ans
}

func testSubsets(nums []int) {
	ans := subsets(nums)
	fmt.Printf("nums %v, get %v\n", nums, ans)
}

func main() {
	testSubsets([]int{1})
	testSubsets([]int{1, 2})
	testSubsets([]int{1, 2, 3})
}
