package main

import "fmt"

func containsDuplicate(nums []int) bool {
	tbl := make(map[int]bool)
	for _, num := range nums {
		if _, ok := tbl[num]; ok {
			return true
		} else {
			tbl[num] = true
		}
	}
	return false
}

func main() {
	fmt.Println("vim-go")
}
