package main

import "fmt"

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	cnt := 0
	for n > 0 {
		if n&0x1 == 1 {
			cnt++
		}
		if cnt > 1 {
			return false
		}
		n >>= 1
	}
	return cnt == 1
}

func testPower(n int) {
	fmt.Printf("%d, get %t\n", n, isPowerOfTwo(n))
}

func main() {
	for i := 0; i <= 32; i++ {
		testPower(i)
	}
}
