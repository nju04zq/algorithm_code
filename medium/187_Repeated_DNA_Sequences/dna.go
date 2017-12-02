package main

import "fmt"

func findRepeatedDnaSequences(s string) []string {
	tbl := make(map[string]int)
	res := make([]string, 0)
	for i := 10; i <= len(s); i++ {
		start, end := i-10, i
		tok := s[start:end]
		cnt, ok := tbl[tok]
		if !ok {
			tbl[tok] = 1
		} else {
			if cnt == 1 {
				res = append(res, tok)
			}
			tbl[tok]++
		}
	}
	return res
}

func testDNA(s string) {
	fmt.Printf("%q, get %v\n", s, findRepeatedDnaSequences(s))
}

func main() {
	testDNA("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT")
}
