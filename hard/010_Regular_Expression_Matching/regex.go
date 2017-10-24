package main

import "fmt"

func compare(p, s byte) bool {
	if p == '.' || p == s {
		return true
	} else {
		return false
	}
}

func isMatchInternal(s, p string, sStart, pStart int) bool {
	if pStart >= len(p) && sStart >= len(s) {
		return true
	} else if pStart >= len(p) {
		return false
	}
	if pStart == (len(p)-1) || p[pStart+1] != '*' {
		if sStart >= len(s) {
			return false
		} else if !compare(p[pStart], s[sStart]) {
			return false
		} else {
			return isMatchInternal(s, p, sStart+1, pStart+1)
		}
	}
	// handle something like .*, a*
	for i := -1; sStart+i < len(s); i++ {
		if i >= 0 && !compare(p[pStart], s[sStart+i]) {
			return false
		}
		if isMatchInternal(s, p, sStart+i+1, pStart+2) {
			return true
		}
	}
	return false
}

func isMatch(s string, p string) bool {
	return isMatchInternal(s, p, 0, 0)
}

type testError struct {
	err string
}

func testIsMatch(s, p string, ans bool) {
	var res bool
	defer func() {
		switch pa := recover(); pa {
		case nil:
		case testError{}:
			panic(pa)
		default:
			fmt.Printf("p %q, s %q, ans %t, get %t\n", p, s, ans, res)
			panic(pa)
		}
	}()
	res = isMatch(s, p)
	if res != ans {
		s := fmt.Sprintf("p %q, s %q, ans %t, get %t\n", p, s, ans, res)
		panic(testError{s})
	}
}

func main() {
	testIsMatch("aab", "c*a*b", true)

	testIsMatch("aa", "a", false)
	testIsMatch("aa", "aa", true)
	testIsMatch("aa", "aaa", false)
	testIsMatch("aa", "a*", true)
	testIsMatch("aa", ".*", true)
	testIsMatch("ab", ".*", true)
	testIsMatch("aab", "c*a*b", true)
	testIsMatch("aa", "c*ab*", false)
	testIsMatch("a", "c*ab*", true)
	testIsMatch("ca", "c*ab*", true)
	testIsMatch("cabb", "c*ab*", true)
	testIsMatch("", "c*", true)
}
