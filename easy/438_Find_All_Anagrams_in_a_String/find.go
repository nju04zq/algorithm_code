package main

import "fmt"

func isTheSame(t1, t2 []int) bool {
	for i := 0; i < len(t1); i++ {
		if t1[i] != t2[i] {
			return false
		}
	}
	return true
}

func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	if len(s) < len(p) {
		return res
	}
	tblp := make([]int, 255)
	tblw := make([]int, 255)
	for i := 0; i < len(p); i++ {
		tblp[int(p[i])]++
	}
	for i := 0; i < len(s); i++ {
		tblw[int(s[i])]++
		if i >= len(p) {
			tblw[int(s[i-len(p)])]--
		}
		if i >= len(p)-1 && isTheSame(tblw, tblp) {
			res = append(res, i-len(p)+1)
		}
	}
	return res
}

func testFind(s, p string) {
	fmt.Printf("s %q, p %q, get %v\n", s, p, findAnagrams(s, p))
}

func main() {
	testFind("cbaebabacd", "abc")
	testFind("abab", "ab")
}
