package main

import "fmt"

func pow(x float64, n int) float64 {
	f := x
	var res float64 = 1
	for {
		if n <= 0 {
			break
		}
		i := n & 0x1
		if i == 1 {
			res *= f
		}
		f = f * f
		n >>= 1
	}
	return res
}

func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1 / pow(x, -n)
	} else {
		return pow(x, n)
	}
}

func myPowBF(x float64, n int) float64 {
	var m int
	if n < 0 {
		m = -n
	} else {
		m = n
	}
	var res float64 = 1
	for i := 0; i < m; i++ {
		res *= x
	}
	if n < 0 {
		return 1 / res
	} else {
		return res
	}
}

func testMyPow(x float64, n int) {
	res := myPow(x, n)
	ans := myPowBF(x, n)
	if res != ans {
		panic(fmt.Sprintf("Fail on %f**%d, get %f, should %f", x, n, res, ans))
	}
}

func main() {
	testMyPow(2, 0)
	testMyPow(2, 1)
	testMyPow(2, 2)
	testMyPow(2, 3)
	testMyPow(2, 4)
	testMyPow(2, 5)
	testMyPow(2, 6)
	testMyPow(2, 7)
	testMyPow(2, 8)
	testMyPow(2, 9)
	testMyPow(2, -1)
	testMyPow(2, -2)
}
