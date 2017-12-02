package main

import "fmt"

func majorityElement(nums []int) int {
	var majority, cnt int
	for _, num := range nums {
		if cnt == 0 {
			majority = num
			cnt++
		} else if majority == num {
			cnt++
		} else {
			cnt--
		}
	}
	return majority
}

func testMajority(nums []int) {
	fmt.Printf("%v, %d\n", nums, majorityElement(nums))
}

func main() {
	testMajority([]int{2, 3, 2, 2, 4})
	testMajority([]int{3, 3, 4})
}
