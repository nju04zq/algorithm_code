package main

import "fmt"

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
	diff := len(p)
	for i := 0; i < len(s); i++ {
		if i >= len(p) {
			j := int(s[i-len(p)])
			tblw[j]--
			if tblw[j] < tblp[j] && tblp[j] > 0 {
				diff++
			}
		}
		j := int(s[i])
		tblw[j]++
		if tblw[j] <= tblp[j] {
			diff--
		}
		if diff == 0 {
			res = append(res, i-len(p)+1)
		}
		//fmt.Printf("i %d, diff %d\n", i, diff)
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
