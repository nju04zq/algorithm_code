package main

import "fmt"

var roots = []int{2, 3, 5}

func isUgly(num int) bool {
	if num == 0 {
		return false
	} else if num == 1 {
		return true
	}
	for _, root := range roots {
		if num%root == 0 {
			return isUgly(num / root)
		}
	}
	return false
}

func testUgly(num int) {
	fmt.Printf("%d, ugly %t\n", num, isUgly(num))
}

func main() {
	testUgly(6)
	testUgly(8)
	testUgly(14)
}
