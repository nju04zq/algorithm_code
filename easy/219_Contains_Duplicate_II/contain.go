package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	tbl := make(map[int]int)
	for i, num := range nums {
		j, ok := tbl[num]
		if !ok {
			tbl[num] = i
			continue
		}
		if i-j <= k {
			return true
		}
		tbl[num] = i
	}
	return false
}

func main() {
	fmt.Println("vim-go")
}
