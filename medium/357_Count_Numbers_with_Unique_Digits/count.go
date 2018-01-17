package main

import "fmt"

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return 10
	}
	total, prev := 10+9*9, 9*9
	for i := 3; i <= n && i <= 10; i++ {
		cur := prev * (10 - i + 1)
		total += cur
		prev = cur
	}
	return total
}

func testCount(n int) {
	fmt.Printf("%d, get %d\n", n, countNumbersWithUniqueDigits(n))
}

func main() {
	for i := 1; i <= 11; i++ {
		testCount(i)
	}
}
