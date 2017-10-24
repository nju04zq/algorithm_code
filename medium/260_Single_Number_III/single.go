package main

import "fmt"

func singleNumber(nums []int) []int {
	var xor int
	for _, num := range nums {
		xor ^= num
	}
	lowest := xor & -xor
	var a, b int
	for _, num := range nums {
		if num&lowest == 0 {
			a ^= num
		} else {
			b ^= num
		}
	}
	return []int{a, b}
}

func testSingleNumber(nums []int) {
	result := singleNumber(nums)
	fmt.Printf("nums %v, get %v\n", nums, result)
}

func main() {
	testSingleNumber([]int{1, 2, 1, 3, 2, 5})
}
