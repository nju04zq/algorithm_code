package main

import "fmt"

func countBits(num int) []int {
	counts := make([]int, num+1)
	for i := 1; i <= num; i++ {
		counts[i] = counts[i-(i&-i)] + 1
	}
	return counts
}

func testCount(num int) {
	fmt.Printf("%d, get %v\n", num, countBits(num))
}

func main() {
	testCount(5)
}
