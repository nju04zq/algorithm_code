package main

import "fmt"
import "math"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func nthSuperUglyNumber(n int, primes []int) int {
	if n == 1 {
		return 1
	}
	idxs := make([]int, len(primes))
	for i := 0; i < len(primes); i++ {
		idxs[i] = 1
	}
	res := []int{0, 1}
	getVal := func(j int) int {
		return res[idxs[j]] * primes[j]
	}
	for i := 2; i <= n; i++ {
		minVal := math.MaxInt32
		for j := 0; j < len(primes); j++ {
			minVal = min(minVal, getVal(j))
		}
		for j := 0; j < len(primes); j++ {
			if getVal(j) == minVal {
				idxs[j]++
			}
		}
		res = append(res, minVal)
	}
	return res[len(res)-1]
}

func testSuper(n int, primes []int) {
	fmt.Printf("n %d, %v, get %d\n", n, primes, nthSuperUglyNumber(n, primes))
}

func main() {
	primes := []int{2, 7, 13, 19}
	for i := 1; i <= 12; i++ {
		testSuper(i, primes)
	}
}
