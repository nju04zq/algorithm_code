package main

import "fmt"

func sortColors(nums []int) {
	cnts := []int{0, 0, 0}
	for _, num := range nums {
		cnts[num]++
	}
	i := 0
	for j, cnt := range cnts {
		for k := 0; k < cnt; k++ {
			nums[i] = j
			i++
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
}
