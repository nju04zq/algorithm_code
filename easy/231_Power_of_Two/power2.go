package main

import "fmt"

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	if (n - (n & -n)) == 0 {
		return true
	} else {
		return false
	}
}

func testPower(n int) {
	fmt.Printf("%d, get %t\n", n, isPowerOfTwo(n))
}

func main() {
	for i := 0; i <= 32; i++ {
		testPower(i)
	}
}
