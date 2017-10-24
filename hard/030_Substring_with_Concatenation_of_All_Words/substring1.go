package main

import "fmt"

func checkWindow(s string, start, end, k int, tbl map[string]int) bool {
	occurence := make(map[string]int)
	for i := start; i < end; i += k {
		word := s[i : i+k]
		if _, ok := tbl[word]; !ok {
			return false
		}
		occurence[word]++
		if occurence[word] > tbl[word] {
			return false
		}
	}
	return true
}

func findSubstring(s string, words []string) []int {
	if s == "" || len(words) == 0 {
		return []int{}
	}
	k := len(words[0])
	if len(s) < k {
		return []int{}
	}
	tbl := make(map[string]int)
	for _, word := range words {
		tbl[word]++
	}
	total := k * len(words)
	res := make([]int, 0)
	for i := 0; i <= len(s)-total; i++ {
		end := i + total
		if checkWindow(s, i, end, k, tbl) {
			res = append(res, i)
		}
	}
	return res
}

func testFind(s string, words []string) {
	ans := findSubstring(s, words)
	fmt.Printf("Find words %v in %q, get %v\n", words, s, ans)
}

func main() {
	testFind("", []string{})
	testFind("barfoothefoobarman", []string{"foo", "bar"})              //[0, 9]
	testFind("barfoothefoo", []string{"foo", "bar"})                    //[0]
	testFind("barfoothefoomanbar", []string{"foo", "bar"})              //[0]
	testFind("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}) //[6, 9, 12]
	testFind("aaaaaabaa", []string{"aa", "aa"})                         //[0, 1, 2]
}
