package main

import "fmt"

func reverse(nums []int) {
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	i0 := 0
	for i, num := range nums {
		if i == 0 {
			continue
		}
		if num > nums[i-1] {
			i0 = i
		}
	}
	if i0 == 0 {
		reverse(nums)
		return
	}
	i1 := i0
	for i := i0; i < len(nums); i++ {
		if nums[i] > nums[i0-1] {
			i1 = i
		}
	}
	nums[i0-1], nums[i1] = nums[i1], nums[i0-1]
	reverse(nums[i0:])
}

func testNextPermutation(nums []int) {
	fmt.Println("==============")
	fmt.Printf("Before: %v\n", nums)
	nextPermutation(nums)
	fmt.Printf("After: %v\n", nums)
}

func main() {
	testNextPermutation([]int{1, 2, 3})    // 1, 3, 2
	testNextPermutation([]int{3, 2, 1})    // 1, 2, 3
	testNextPermutation([]int{1, 1, 5})    // 1, 5, 1
	testNextPermutation([]int{1, 5, 1})    // 5, 1, 1
	testNextPermutation([]int{1, 5, 2, 1}) // 2, 1, 1, 5
}
