package main

import "fmt"

func isScramble(s1 string, s2 string) bool {
	//fmt.Printf("s1 %q, s2 %q\n", s1, s2)
	if s1 == s2 {
		return true
	}
	cnt := make([]int, 256)
	for i := 0; i < len(s1); i++ {
		cnt[s1[i]]++
	}
	for i := 0; i < len(s2); i++ {
		if cnt[s2[i]] == 0 {
			//fmt.Printf("s1 %q, s2 %q, false\n", s1, s2)
			return false
		}
		cnt[s2[i]]--
	}
	for i := 1; i < len(s1); i++ {
		if isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]) {
			//fmt.Printf("s1 %q, s2 %q, true\n", s1, s2)
			return true
		}
		if isScramble(s1[:i], s2[len(s2)-i:]) && isScramble(s1[i:], s2[:len(s2)-i]) {
			//fmt.Printf("s1 %q, s2 %q, true\n", s1, s2)
			return true
		}
	}
	//fmt.Printf("s1 %q, s2 %q, false\n", s1, s2)
	return false
}

func testIsScramble(s1, s2 string) {
	fmt.Printf("%q, %q, get %t\n", s1, s2, isScramble(s1, s2))
}

func main() {
	testIsScramble("rgtae", "great")
	testIsScramble("rgeat", "great")
	testIsScramble("rgtae", "rgeat")
}
