package main

import "fmt"
import "sort"

func makeCopy(src []int) []int {
	dst := make([]int, len(src))
	copy(dst, src)
	return dst
}

func subsetsInternal(nums, path []int, start, maxLen int, ans [][]int) [][]int {
	if len(path) == maxLen {
		ans = append(ans, makeCopy(path))
		return ans
	}
	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		path = append(path, nums[i])
		ans = subsetsInternal(nums, path, i+1, maxLen, ans)
		path = path[:len(path)-1]
	}
	return ans
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	path := make([]int, 0, len(nums))
	ans := make([][]int, 0)
	for i := 0; i <= len(nums); i++ {
		ans = subsetsInternal(nums, path, 0, i, ans)
	}
	return ans
}

func testSubsets(nums []int) {
	ans := subsetsWithDup(nums)
	fmt.Printf("nums %v, get %v\n", nums, ans)
}

func main() {
	testSubsets([]int{1})
	testSubsets([]int{1, 2})
	testSubsets([]int{1, 2, 3})
	testSubsets([]int{1, 2, 2})
}
