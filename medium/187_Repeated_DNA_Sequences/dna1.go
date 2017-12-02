package main

import "fmt"

func transform(s string, lookup map[byte]int) int {
	mark := 0
	for i := 0; i < len(s); i++ {
		k := lookup[s[i]]
		mark = (mark << 2) | k
	}
	return mark
}

func findRepeatedDnaSequences(s string) []string {
	lookup := map[byte]int{'A': 0, 'C': 1, 'G': 2, 'T': 3}
	tbl := make(map[int]int)
	res := make([]string, 0)
	for i := 10; i <= len(s); i++ {
		start, end := i-10, i
		tok := s[start:end]
		mark := transform(tok, lookup)
		cnt, ok := tbl[mark]
		if !ok {
			tbl[mark] = 1
		} else {
			if cnt == 1 {
				res = append(res, tok)
			}
			tbl[mark]++
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
