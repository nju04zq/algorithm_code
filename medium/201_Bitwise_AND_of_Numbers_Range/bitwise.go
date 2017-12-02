package main

import "fmt"
import "math/rand"
import "time"

func rangeBitwiseAnd(m int, n int) int {
	res := 1
	for m < n {
		m >>= 1
		n >>= 1
		res <<= 1
	}
	return res * m
}

func bf(m, n int) int {
	res := m
	for i := m; i <= n; i++ {
		res &= i
	}
	return res
}

func MakeRandInt() int {
	maxNum := 10000
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int() % maxNum
}

func testBitwise() {
	m, n := MakeRandInt(), 0
	for {
		n = MakeRandInt()
		if n >= m {
			break
		}
	}
	res := rangeBitwiseAnd(m, n)
	ans := bf(m, n)
	if res != ans {
		panic(fmt.Sprintf("m %d, n %d, get %d, ans %d\n", m, n, res, ans))
	}
}

func main() {
	for i := 0; i < 10000; i++ {
		testBitwise()
	}
}
