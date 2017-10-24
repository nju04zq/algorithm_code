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
		path = append(path, num)
		ans = combinationInternal(nums, i, path, target-num, ans)
		path = path[:len(path)-1]
	}
	return ans
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	path := make([]int, 0)
	ans := make([][]int, 0)
	ans = combinationInternal(candidates, 0, path, target, ans)
	return ans
}

func testCombinationSum(nums []int, target int) {
	ans := combinationSum(nums, target)
	fmt.Printf("nums %v, target %d, ans %v\n", nums, target, ans)
}

func main() {
	testCombinationSum([]int{1, 2}, 4)
	testCombinationSum([]int{2, 3, 6, 7}, 1)
	testCombinationSum([]int{2, 3, 6, 7}, 2)
	testCombinationSum([]int{2, 3, 6, 7}, 7)
	testCombinationSum([]int{2, 3, 6, 7}, 8)
	testCombinationSum([]int{2, 3, 6, 7}, 20)
}
