package main

import "fmt"

func makeGray(digits []int) int {
	n := 0
	for i := 0; i < len(digits); i++ {
		n += (1 << uint(i)) * digits[len(digits)-i-1]
	}
	return n
}

func gray(digits []int, res []int, n int) []int {
	if n == 0 {
		res = append(res, makeGray(digits))
		return res
	}
	res = gray(digits, res, n-1)
	if digits[len(digits)-n] == 0 {
		digits[len(digits)-n] = 1
	} else {
		digits[len(digits)-n] = 0
	}
	res = gray(digits, res, n-1)
	return res
}

func grayCode(n int) []int {
	res := make([]int, 0, 1<<uint(n))
	digits := make([]int, n)
	res = gray(digits, res, n)
	return res
}

func testGray(n int) {
	fmt.Println(n, grayCode(n))
}

func main() {
	testGray(0)
	testGray(1)
	testGray(2)
	testGray(3)
	testGray(4)
}
