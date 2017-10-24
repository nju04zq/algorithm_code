package main

import "fmt"

func climbStairs(n int) int {
	a, b := 1, 2
	if n == 0 {
		return 0
	} else if n == 1 {
		return a
	} else if n == 2 {
		return b
	}
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func main() {
	fmt.Printf("%d\n", climbStairs(3))
}
