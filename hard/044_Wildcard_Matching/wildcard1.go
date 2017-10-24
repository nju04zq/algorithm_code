package main

import "fmt"

func toInt(x bool) int {
	if x {
		return 1
	} else {
		return 0
	}
}

func dump(a [][]bool, s, p string) {
	fmt.Printf("    ")
	for i, _ := range s {
		fmt.Printf("%c ", s[i])
	}
	fmt.Printf("\n")
	for i := 0; i <= len(p); i++ {
		if i == 0 {
			fmt.Printf("  ")
		} else {
			fmt.Printf("%c ", p[i-1])
		}
		for j := 0; j <= len(s); j++ {
			fmt.Printf("%d ", toInt(a[i][j]))
		}
		fmt.Printf("\n")
	}
}

func make2dArray(m, n int) [][]bool {
	a := make([][]bool, m)
	for i, _ := range a {
		a[i] = make([]bool, n)
	}
	return a
}

func isMatch(s string, p string) bool {
	a := make2dArray(len(p)+1, len(s)+1)
	a[0][0] = true
	for i := 1; i <= len(p); i++ {
		if p[i-1] == '*' {
			a[i][0] = a[i-1][0]
		}
	}
	for i := 1; i <= len(p); i++ {
		for j := 1; j <= len(s); j++ {
			if p[i-1] != '*' {
				if p[i-1] == '?' || p[i-1] == s[j-1] {
					a[i][j] = a[i-1][j-1]
				}
			} else if a[i-1][j-1] || a[i][j-1] || a[i-1][j] {
				a[i][j] = true
			}
		}
	}
	//dump(a, s, p)
	return a[len(p)][len(s)]
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
