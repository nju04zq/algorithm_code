package main

import "fmt"

func integerBreak(n int) int {
	if n == 2 {
		return 1
	} else if n == 3 {
		return 2
	}
	product := 1
	for n > 4 {
		product *= 3
		n -= 3
	}
	product *= n
	return product
}

func testBreak(n int) {
	fmt.Printf("%d, get %d\n", n, integerBreak(n))
}

func main() {
	for i := 2; i <= 20; i++ {
		testBreak(i)
	}
}
