package main

import "fmt"

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	marks := make([]bool, n)
	marks[0] = true
	marks[1] = true
	for i := 2; i*i < n; i++ {
		if marks[i] {
			continue
		}
		for j := 2; i*j < n; j++ {
			marks[j*i] = true
		}
	}
	cnt := 0
	for i := 0; i < n; i++ {
		if !marks[i] {
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Printf("Primes %d, get %d\n", 13, countPrimes(13))
	fmt.Printf("Primes %d, get %d\n", 100, countPrimes(100))
	fmt.Printf("Primes %d, get %d\n", 1, countPrimes(1))
}
