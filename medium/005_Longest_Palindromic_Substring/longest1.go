package main

import "fmt"

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func getMaxPalindrome(s string, i int) int {
	left, right := i, i
	maxLen := 0
	for j := 0; left >= 0 && right < len(s); j++ {
		if s[i-j] != s[i+j] {
			break
		}
		maxLen = max(maxLen, 2*j+1)
		left--
		right++
	}
	left, right = i, i+1
	for j := 0; left >= 0 && right < len(s); j++ {
		if s[i-j] != s[i+1+j] {
			break
		}
		maxLen = max(maxLen, 2*j+2)
		left--
		right++
	}
	return maxLen
}

func longestPalindrome(s string) string {
	maxLen, maxStr := 0, ""
	for i := 0; i < len(s); i++ {
		if (i + maxLen/2) >= len(s) {
			break
		}
		pLen := getMaxPalindrome(s, i)
		if pLen > maxLen {
			maxLen = pLen
			start := i - (maxLen-1)/2
			maxStr = s[start : start+maxLen]
		}
	}
	return maxStr
}

func testLongestPalindrome(s string) {
	p := longestPalindrome(s)
	fmt.Printf("s: %q, p %q\n", s, p)
}

func main() {
	testLongestPalindrome("babad")
	testLongestPalindrome("cbbd")
	testLongestPalindrome("abcda")
}
