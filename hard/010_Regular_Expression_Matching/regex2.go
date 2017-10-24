package main

/*
 * twice time cost compared with regex1
 */

import "fmt"

func compare(p, s byte) bool {
	if p == '.' || p == s {
		return true
	} else {
		return false
	}
}

type dpContainer struct {
	m, n int
	res  []bool
}

func (dp *dpContainer) get(i, j int) bool {
	return dp.res[i*dp.n+j]
}

func (dp *dpContainer) set(i, j int, val bool) {
	dp.res[i*dp.n+j] = val
}

func makeDpContainer(m, n int) *dpContainer {
	dp := dpContainer{m: m, n: n}
	dp.res = make([]bool, m*n)
	return &dp
}

func isMatch(s string, p string) bool {
	dp := makeDpContainer(len(p)+1, len(s)+1)
	dp.set(0, 0, true)
	for i := 1; i <= len(p); i++ {
		for j := 0; j <= len(s); j++ {
			if p[i-1] != '*' {
				if j > 0 && compare(p[i-1], s[j-1]) && dp.get(i-1, j-1) {
					dp.set(i, j, true)
				}
				continue
			}
			if dp.get(i-2, j) {
				dp.set(i, j, true)
				continue
			}
			for k := 1; j-k >= 0; k++ {
				if !compare(p[i-2], s[j-k]) {
					break
				}
				if dp.get(i-2, j-k) {
					dp.set(i, j, true)
					break
				}
			}
		}
	}
	//verifyDp(dp, s, p)
	return dp.get(len(p), len(s))
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

func isMatchRecursive(s string, p string) bool {
	return isMatchInternal(s, p, 0, 0)
}

func verifyDp(dp *dpContainer, s, p string) {
	for i := 0; i <= len(p); i++ {
		for j := 0; j <= len(s); j++ {
			p0, s0 := p[:i], s[:j]
			ans := isMatchRecursive(s0, p0)
			if ans != dp.get(i, j) {
				panic(fmt.Sprintf("i %d %q, j %d %q, should be %t\n", i, p0,
					j, s0, ans))
			}
		}
	}
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
	//testIsMatch("aab", "c*a*b", true)
	testIsMatch("aaba", "ab*a*c*a", false)

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
