package main

import "fmt"

func toInt(r bool) int {
	if r {
		return 1
	} else {
		return 0
	}
}

func dumpHeader(s string) {
	fmt.Printf("    ")
	for _, r := range s {
		fmt.Printf("%c ", r)
	}
	fmt.Println()
}

func dumpArray(c byte, a []bool) {
	fmt.Printf("%c ", c)
	for _, r := range a {
		fmt.Printf("%d ", toInt(r))
	}
	fmt.Println()
}

func isMatch(s string, p string) bool {
	//dumpHeader(s)
	a := make([]bool, len(s)+1)
	a[0] = true
	//dumpArray(' ', a)
	for i := 1; i <= len(p); i++ {
		prev := a[0]
		if p[i-1] == '*' {
			a[0] = prev
		} else {
			a[0] = false
		}
		for j := 1; j <= len(s); j++ {
			tmp := a[j]
			if p[i-1] != '*' {
				if p[i-1] == '?' || p[i-1] == s[j-1] {
					a[j] = prev
				} else {
					a[j] = false
				}
			} else if prev || a[j-1] || a[j] {
				a[j] = true
			} else {
				a[j] = false
			}
			prev = tmp
		}
		//dumpArray(p[i-1], a)
	}
	return a[len(s)]
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
