package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	tbl := make(map[byte]int)
	for i, _ := range s {
		if _, ok := tbl[s[i]]; !ok {
			tbl[s[i]] = 1
		} else {
			tbl[s[i]]++
		}
	}
	for i, _ := range t {
		if _, ok := tbl[t[i]]; !ok {
			return false
		} else if tbl[t[i]] <= 0 {
			return false
		} else {
			tbl[t[i]]--
		}
	}
	return true
}

func main() {
	fmt.Println("vim-go")
}
