package main

import "fmt"

func isPerfectSquare(num int) bool {
	r := num
	for r*r > num {
		r = (r + num/r) / 2
	}
	return r*r == num
}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Printf("%d, %t\n", i, isPerfectSquare(i))
	}
}
