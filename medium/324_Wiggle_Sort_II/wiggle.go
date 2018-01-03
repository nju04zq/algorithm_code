package main

import "fmt"
import "sort"

// X0 X1 Y0 Y1
// X0 Y0 X1 Y1
// X1 Y1 X0 Y0
// avoid X1/Y0 sit aside
func wiggleSort(nums []int) {
	n := len(nums)
	temp := make([]int, n)
	for i := 0; i < n; i++ {
		temp[i] = nums[i]
	}
	sort.Ints(temp)
	mid := (n + 1) / 2
	left, right, i := mid-1, n-1, 0
	for left >= 0 || right >= mid {
		if left >= 0 {
			nums[i] = temp[left]
			left--
			i++
		}
		if right >= mid {
			nums[i] = temp[right]
			right--
			i++
		}
	}
}

func verify(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if i%2 == 0 && nums[i] >= nums[i+1] {
			return false
		}
		if i%2 == 1 && nums[i] <= nums[i+1] {
			return false
		}
	}
	return true
}

func testWiggle(nums []int) {
	fmt.Printf("%v\n", nums)
	wiggleSort(nums)
	fmt.Printf("get %v\n", nums)
	if !verify(nums) {
		panic("Not qualified")
	}
}

func main() {
	testWiggle([]int{1, 5, 1, 1, 6, 4})
	testWiggle([]int{1, 3, 2, 2, 3, 1})
	testWiggle([]int{1, 5, 1, 1, 6, 4, 7})
	testWiggle([]int{1, 1, 2, 1, 2, 2, 1})
	testWiggle([]int{5, 3, 1, 2, 6, 7, 8, 5, 5})
	testWiggle([]int{4, 5, 5, 6})
}
