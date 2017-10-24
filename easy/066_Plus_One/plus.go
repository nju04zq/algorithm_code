package main

import "fmt"

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{1}
	}
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += carry
		if digits[i] >= 10 {
			digits[i] -= 10
			carry = 1
		} else {
			carry = 0
			break
		}
	}
	if carry > 0 {
		return append([]int{1}, digits...)
	} else {
		return digits
	}
}

func testPlus(digits []int) {
	fmt.Println(digits)
	fmt.Printf("Get %v\n", plusOne(digits))
}

func main() {
	testPlus([]int{1, 9, 9})
	testPlus([]int{9, 9})
}
