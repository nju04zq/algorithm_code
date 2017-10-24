package main

import "fmt"

import "math"

func splitInt64(x int64) []int64 {
	if x == 0 {
		return []int64{0}
	}
	a := make([]int64, 0)
	for x > 0 {
		a = append(a, x%10)
		x /= 10
	}
	return a
}

func composeInt64(a []int64) int64 {
	var y int64
	for i := 0; i < len(a); i++ {
		y = y*10 + a[i]
	}
	return y
}

func reverseInt64(x int64) int64 {
	a := splitInt64(x)
	y := composeInt64(a)
	return y
}

func reverse(x int) int {
	x1 := int64(x)
	neg := false
	if x1 < 0 {
		neg = true
		x1 = -x1
	}
	y1 := reverseInt64(x1)
	if neg {
		y1 = -y1
	}
	if y1 > math.MaxInt32 {
		return 0
	} else if y1 < math.MinInt32 {
		return 0
	} else {
		return int(y1)
	}
}

func testReverse(x int) {
	fmt.Printf("Reverse %d, get %d\n", x, reverse(x))
}

func main() {
	testReverse(123)
	testReverse(-123)
	testReverse(-1)
	testReverse(1)
	testReverse(0)
	testReverse(1534236469)
}
