package main

import "fmt"

func isMatch(s string, p string) bool {
	slast, mark := -1, -1
	var i, j int
	for i < len(s) {
		if j < len(p) && p[j] == '*' {
			slast, mark = i, j
			j++
		} else if j < len(p) && (p[j] == '?' || p[j] == s[i]) {
			i++
			j++
		} else if mark != -1 {
			i = slast
			j = mark + 1
			slast++
		} else {
			return false
		}
	}
	for j < len(p) {
		if p[j] == '*' {
			j++
		} else {
			return false
		}
	}
	return true
}

func testIsMatch(s, p string, ans bool) {
	res := isMatch(s, p)
	if res != ans {
		panic(fmt.Errorf("s %q, p %q, get %t, should be %t\n", s, p, res, ans))
	}
}

func main() {
	testIsMatch("", "", true)
	testIsMatch("a", "", false)
	testIsMatch("aa", "aa", true)
	testIsMatch("aa", "aaa", false)
	testIsMatch("aaa", "aa", false)
	testIsMatch("aa", "*", true)
	testIsMatch("aa", "a*", true)
	testIsMatch("aa", "a*a*", true)
	testIsMatch("ab", "*?", true)
	testIsMatch("b", "c*a*b", false)
	testIsMatch("aab", "c*a*b", false)
}
