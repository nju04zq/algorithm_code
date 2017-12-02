package main

import "fmt"

func addDigits(num int) int {
	if num < 10 {
		return num
	}
	num1 := 0
	for num > 0 {
		num1 += (num % 10)
		num /= 10
	}
	return addDigits(num1)
}

func main() {
	fmt.Printf("%d, get %d\n", 38, addDigits(38))
}
