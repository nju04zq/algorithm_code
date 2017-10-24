package main

import "fmt"

func twoSum(nums []int, target int) []int {
	tbl := make(map[int][]int, len(nums))
	for i, num := range nums {
		_, ok := tbl[num]
		if !ok {
			tbl[num] = []int{i}
		} else {
			tbl[num] = append(tbl[num], i)
		}
		wanted := target - num
		a, ok := tbl[wanted]
		if !ok {
			continue
		}
		for _, j := range a {
			if j != i {
				return []int{j, i}
			}
		}
	}
	return nil
}

func main() {
	nums := []int{1, 5, 7, 9}
	target := 12
	fmt.Printf("twoSum %v\n", twoSum(nums, target))
}
