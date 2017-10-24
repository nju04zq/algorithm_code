package main

import "fmt"
import "sort"

func combinationInternal(nums []int, start int, path []int, target int, ans [][]int) [][]int {
	for i := start; i < len(nums); i++ {
		num := nums[i]
		if num > target {
			break
		} else if num == target {
			res := make([]int, len(path))
			copy(res, path)
			res = append(res, num)
			ans = append(ans, res)
			break
		}
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		path = append(path, num)
		ans = combinationInternal(nums, i+1, path, target-num, ans)
		path = path[:len(path)-1]
	}
	return ans
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	path := make([]int, 0)
	ans := make([][]int, 0)
	ans = combinationInternal(candidates, 0, path, target, ans)
	return ans
}

func testCombinationSum2(nums []int, target int) {
	ans := combinationSum2(nums, target)
	fmt.Printf("nums %v, target %d, get %v\n", nums, target, ans)
}

func main() {
	testCombinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
}
