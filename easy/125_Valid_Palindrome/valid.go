package main

import "fmt"

func isAplhanumeric(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	} else if c >= 'a' && c <= 'z' {
		return true
	} else if c >= 'A' && c <= 'Z' {
		return true
	} else {
		return false
	}
}

func toLower(c byte) byte {
	if c >= '0' && c <= '9' {
		return c
	} else if c >= 'a' && c <= 'z' {
		return c
	} else {
		return c - 'A' + 'a'
	}
}

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	i, j := 0, len(s)-1
	for i < j {
		c1, c2 := s[i], s[j]
		if !isAplhanumeric(c1) {
			i++
		} else if !isAplhanumeric(c2) {
			j--
		} else if toLower(c1) != toLower(c2) {
			return false
		} else {
			i++
			j--
		}
	}
	return true
}

func testIsPalindrome(s string) {
	fmt.Printf("%q, get %t\n", s, isPalindrome(s))
}

func main() {
	testIsPalindrome("A man, a plan, a canal: Panama")
	testIsPalindrome("race a car")
}
