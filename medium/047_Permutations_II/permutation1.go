package main

import "fmt"
import "sort"

func makeCopy(src []int) []int {
	dst := make([]int, len(src))
	copy(dst, src)
	return dst
}

func permuteInternal(nums, path, mask []int, ans [][]int) [][]int {
	if len(path) == len(nums) {
		ans = append(ans, makeCopy(path))
		return ans
	}
	prev := -1
	for i, _ := range mask {
		if mask[i] == 1 {
			continue
		}
		if prev != -1 && nums[i] == nums[prev] {
			continue
		}
		prev = i
		mask[i] = 1
		path = append(path, nums[i])
		ans = permuteInternal(nums, path, mask, ans)
		mask[i] = 0
		path = path[:len(path)-1]
	}
	return ans
}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	path := make([]int, 0, len(nums))
	mask := make([]int, len(nums))
	ans := make([][]int, 0)
	ans = permuteInternal(nums, path, mask, ans)
	return ans
}

func testPermute(nums []int) {
	ans := permuteUnique(nums)
	fmt.Printf("nums %v, get %v\n", nums, ans)
}

func main() {
	testPermute([]int{1})
	testPermute([]int{1, 1})
	testPermute([]int{1, 1, 2})
}
