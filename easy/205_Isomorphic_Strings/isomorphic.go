package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	tbls := make(map[byte]byte)
	tblt := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if _, ok := tbls[s[i]]; ok {
			if tbls[s[i]] != t[i] {
				return false
			}
		} else if _, ok := tblt[t[i]]; ok {
			return false
		} else {
			tbls[s[i]] = t[i]
			tblt[t[i]] = s[i]
		}
	}
	return true
}

func testIsIsomorphic(s, t string) {
	fmt.Printf("%q, %q, get %t\n", s, t, isIsomorphic(s, t))
}

func main() {
	testIsIsomorphic("egg", "add")
	testIsIsomorphic("foo", "bar")
	testIsIsomorphic("paper", "title")
	testIsIsomorphic("aa", "ab")
}
