package main

import "fmt"

func reverse(s string) string {
	buf := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		buf[i] = s[len(s)-i-1]
	}
	return string(buf)
}

func computePrefix(s string) []int {
	prefix := make([]int, len(s))
	prefix[0] = -1
	for i := 1; i < len(s); i++ {
		j := prefix[i-1]
		for j != -1 {
			if s[j+1] == s[i] {
				break
			}
			j = prefix[j]
		}
		if s[j+1] == s[i] {
			prefix[i] = j + 1
		} else {
			prefix[i] = -1
		}
	}
	return prefix
}

func shortestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	r := reverse(s)
	prefix := computePrefix(s)
	var i, j int
	for i < len(s) {
		if r[i] == s[j] {
			i++
			j++
		} else if j == 0 {
			i++
		} else {
			j = prefix[j-1] + 1
		}
	}
	buf := make([]byte, len(s)+len(s)-j)
	k := 0
	for i := len(s) - 1; i >= j; i-- {
		buf[k] = s[i]
		k++
	}
	for i := 0; i < len(s); i++ {
		buf[k] = s[i]
		k++
	}
	return string(buf)
}

func testShortest(s string) {
	fmt.Printf("%q, get %s\n", s, shortestPalindrome(s))
}

func main() {
	testShortest("")
	testShortest("a")
	testShortest("acca")
	testShortest("accay")
	testShortest("accayy")
	testShortest("accayyy")
	testShortest("aacecaaa")
}
