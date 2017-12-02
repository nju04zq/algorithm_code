package main

import "fmt"

func trailingZeroes(n int) int {
	cnt := 0
	for n >= 5 {
		cnt += (n / 5)
		n /= 5
	}
	return cnt
}

func testTrailingZeros(n int) {
	fmt.Printf("n %d, get %d\n", n, trailingZeroes(n))
}

func main() {
	testTrailingZeros(10)
	testTrailingZeros(100)
	testTrailingZeros(900)
	testTrailingZeros(1000)
}
