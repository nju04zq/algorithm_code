package main

import "fmt"

func countBase(n int, base int, tbl map[int]int) int {
	if base == 1 {
		return 1
	}
	if cnt, ok := tbl[n]; ok {
		return cnt
	}
	cnt := 10*countBase(n/10, base/10, tbl) + base
	tbl[n] = cnt
	return cnt
}

func countNonBase(n int, tbl map[int]int) int {
	if n <= 0 {
		return 0
	}
	base := 1
	for m := n; m >= 10; m /= 10 {
		base *= 10
	}
	if base == 1 {
		return 1
	}
	total := 0
	n1 := n / base * base
	n2 := n - n1
	if n/base == 1 {
		total += (n2 + 1)
	} else {
		total += base
	}
	//fmt.Println(n, base, n1, n2, total)
	total += (countBase(base-1, base/10, tbl) * (n1 / base))
	//fmt.Println(n, base, "add", base-1, total)
	total += countNonBase(n2, tbl)
	//fmt.Println(n, base, "add", n2, total)
	return total
}

func countDigitOne(n int) int {
	tbl := make(map[int]int)
	return countNonBase(n, tbl)
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
	testCount(101)
	for i := 0; i <= 10000; i++ {
		fmt.Printf("\r%d", i)
		testCount(i)
	}
	fmt.Println()
}
