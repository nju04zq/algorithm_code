package main

import "fmt"

func sortColors(nums []int) {
	var white, blue int
	for i, num := range nums {
		if num == 0 {
			nums[i] = nums[blue]
			nums[blue] = nums[white]
			nums[white] = 0
			blue++
			white++
		} else if num == 1 {
			nums[i] = nums[blue]
			nums[blue] = 1
			blue++
		}
	}
}

func testSort(nums []int) {
	fmt.Printf("Before sort: %v\n", nums)
	sortColors(nums)
	fmt.Printf("After sort: %v\n", nums)
}

func main() {
	testSort([]int{1, 2, 0})
	testSort([]int{1, 2})
	testSort([]int{1})
	testSort([]int{0})
	testSort([]int{1, 1, 1, 2, 2, 0, 0})
}
