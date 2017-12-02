package main

import "fmt"

func happyOneRound(n int) int {
	m := 0
	for n > 0 {
		x := n % 10
		m += x * x
		n /= 10
	}
	return m
}

func isHappy(n int) bool {
	visited := make(map[int]bool)
	visited[n] = true
	for {
		n = happyOneRound(n)
		if _, ok := visited[n]; ok {
			break
		}
		visited[n] = true
		if n == 1 {
			break
		}
	}
	return n == 1
}

func main() {
	fmt.Printf("%d, get %t\n", 19, isHappy(19))
	fmt.Printf("%d, get %t\n", 2, isHappy(2))
}
