package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x == 0 {
		return true
	}
	wReverse := 1
	for y := x; y >= 10; {
		y /= 10
		wReverse *= 10
	}
	w := 1
	for wReverse > w {
		yHigh := (x / wReverse) % 10
		yLow := (x / w) % 10
		if yHigh != yLow {
			return false
		}
		wReverse /= 10
		w *= 10
	}
	return true
}

func testIsPalindrome(x int) {
	fmt.Printf("%d isPalindrome %t\n", x, isPalindrome(x))
}

func main() {
	testIsPalindrome(-1)
	testIsPalindrome(0)
	testIsPalindrome(1)
	testIsPalindrome(10)
	testIsPalindrome(11)
	testIsPalindrome(121)
	testIsPalindrome(12321)
	testIsPalindrome(2<<31 - 1)
}
