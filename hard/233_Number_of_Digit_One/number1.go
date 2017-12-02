package main

import "fmt"

func countDigitOne(n int) int {
	total := 0
	for m := 1; m <= n; m *= 10 {
		a, b := n/m, n%m
		total += (a + 8) / 10 * m
		if a%10 == 1 {
			total += (b + 1)
		}
	}
	return total
}

func splitInt(n int) []int {
	digits := make([]int, 0)
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	return digits
}

func bf(n int) int {
	cnt := 0
	for i := 1; i <= n; i++ {
		digits := splitInt(i)
		for _, digit := range digits {
			if digit == 1 {
				cnt++
			}
		}
	}
	return cnt
}

func testCount(n int) {
	res := countDigitOne(n)
	ans := bf(n)
	if res != ans {
		panic(fmt.Errorf("%d, get %d, ans %d", n, res, ans))
	}
}

func main() {
	for i := 0; i <= 10000; i++ {
		fmt.Printf("\r%d", i)
		testCount(i)
	}
	fmt.Println()
}
