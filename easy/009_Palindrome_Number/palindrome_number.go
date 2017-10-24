package main

import "fmt"

import "math"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x == 0 {
		return true
	}
	input := int64(x)
	var reversed int64
	for x > 0 {
		y := x % 10
		reversed = reversed*10 + int64(y)
		if reversed > math.MaxInt32 {
			return false
		}
		x /= 10
	}
	if reversed == input {
		return true
	} else {
		return false
	}
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
