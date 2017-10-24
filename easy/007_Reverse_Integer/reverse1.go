package main

import "fmt"

import "math"

func splitInt32(x int32) []int32 {
	if x == 0 {
		return []int32{0}
	}
	a := make([]int32, 0)
	for x > 0 {
		a = append(a, x%10)
		x /= 10
	}
	return a
}

func multiply10(x int32) (int32, bool) {
	x1 := int64(x)
	y1 := x1 * 10
	if y1 > math.MaxInt32 {
		return 0, true
	} else if y1 < math.MinInt32 {
		return 0, true
	} else {
		return int32(y1), false
	}
}

func add(x int32, y int32) (int32, bool) {
	z := x + y
	if (z^x) >= 0 || (z^y) >= 0 {
		return z, false
	} else {
		return 0, true
	}
}

func composeInt32(a []int32, neg bool) int32 {
	var y int32
	for i := 0; i < len(a); i++ {
		x := a[i]
		if neg {
			x = -x
		}
		z, overflow := multiply10(y)
		if overflow {
			return 0
		}
		y, overflow = add(z, x)
		if overflow {
			return 0
		}
	}
	return y
}

func reverseInt32(x int32) int32 {
	neg := false
	if x < 0 {
		x = -x
		neg = true
	}
	a := splitInt32(x)
	y := composeInt32(a, neg)
	return y
}

func reverse(x int) int {
	y := reverseInt32(int32(x))
	return int(y)
}

func testReverse(x int) {
	fmt.Printf("Reverse %d, get %d\n", x, reverse(x))
}

func main() {
	testReverse(-123)
	testReverse(-1)
	testReverse(1)
	testReverse(0)
	testReverse(1534236469)
}
