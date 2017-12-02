package main

import "fmt"

func validInternal(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func validPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			break
		}
		i++
		j--
	}
	if i >= j {
		return true
	}
	rc := validInternal(s, i+1, j)
	if rc {
		return true
	}
	rc = validInternal(s, i, j-1)
	if rc {
		return true
	}
	return false
}

func main() {
	s := "aba"
	fmt.Printf("%q, get %t\n", s, validPalindrome(s))
	s = "abca"
	fmt.Printf("%q, get %t\n", s, validPalindrome(s))
}
