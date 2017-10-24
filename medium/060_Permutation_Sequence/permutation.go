package main

import "fmt"

var f = []int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}

func getPermutation(n int, k int) string {
	if k > f[n] {
		return ""
	}
	a := make([]byte, n)
	for i := 0; i < n; i++ {
		a[i] = byte(i) + '1'
	}
	res := make([]byte, n)
	idx, k := 0, k-1
	for ; n > 0; n-- {
		i := k / f[n-1]
		res[idx] = a[i]
		for j := i; j < len(a)-1; j++ {
			a[j] = a[j+1]
		}
		a = a[:len(a)-1]
		k = k % f[n-1]
		idx++
	}
	return string(res)
}

func testPermutation(n int, k int) {
	fmt.Printf("n %d, k %2d, %q\n", n, k, getPermutation(n, k))
}

func main() {
	n := 4
	for k := 1; k <= f[n]; k++ {
		testPermutation(n, k)
	}
}
