package main

import "fmt"

func getGCD(x, y int) int {
	for y != 0 {
		temp := y
		y = x % y
		x = temp
	}
	return x
}

func canMeasureWater(x int, y int, z int) bool {
	if (x + y) < z {
		return false
	}
	if x == z || y == z || x+y == z {
		return true
	}
	gcd := getGCD(x, y)
	return z%gcd == 0
}

func testFill(x, y, z int) {
	fmt.Printf("%d, %d, %d, %t\n", x, y, z, canMeasureWater(x, y, z))
}

func main() {
	testFill(3, 5, 4)
	testFill(2, 6, 5)
	testFill(0, 0, 1)
	testFill(1, 2, 3)
	testFill(4, 6, 8)
	testFill(34, 6, 8)
	testFill(11, 3, 13)
	testFill(11, 13, 1)
}
