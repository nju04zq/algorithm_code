package main

import "fmt"

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	} else if n == 1 {
		return true
	}
	for n > 1 {
		if n%3 != 0 {
			return false
		}
		n /= 3
	}
	return true
}

func main() {
	for i := 0; i < 100; i++ {
		if isPowerOfThree(i) {
			fmt.Println(i)
		}
	}
}
