package main

import "fmt"

func hammingDistance(x int, y int) int {
	z := x ^ y
	cnt := 0
	for z > 0 {
		cnt++
		z -= (z & -z)
	}
	return cnt
}

func testHamming(x, y int) {
	fmt.Printf("x %d, y %d, get %d\n", x, y, hammingDistance(x, y))
}

func main() {
	testHamming(1, 4)
}
