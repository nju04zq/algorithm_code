package main

import "fmt"

func skip(s string, start, k int, word string, occurence map[string]int) int {
	var i int
	for i = start; i <= len(s)-k; i += k {
		w := s[i : i+k]
		occurence[w]--
		if w == word {
			break
		}
	}
	return i + k
}

func clearOccurence(occurence map[string]int) {
	for word, _ := range occurence {
		occurence[word] = 0
	}
}

func find(s string, offset, k, total int, tbl map[string]int) []int {
	start := -1
	res := make([]int, 0)
	occurence := make(map[string]int)
	for i := 0; i <= len(s)-k; i += k {
		word := s[i : i+k]
		fmt.Printf("word %s, i %d, start %d, occurence %v\n", word, i, start, occurence)
		if _, ok := tbl[word]; !ok {
			start = -1
			continue
		}
		if start == -1 {
			start = i
			clearOccurence(occurence)
		}
		occurence[word] += 1
		if occurence[word] > tbl[word] {
			start = skip(s, start, k, word, occurence)
		}
		if i+k-start == total {
			res = append(res, start+offset)
		}
	}
	return res
}

func findSubstring(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 {
		return []int{}
	}
	k := len(words[0])
	if len(s) < k {
		return []int{}
	}
	tbl := make(map[string]int)
	for _, word := range words {
		tbl[word] += 1
	}
	total := len(words) * k
	ans := make([]int, 0)
	for i := 0; i < k; i++ {
		res := find(s[i:], i, k, total, tbl)
		ans = append(ans, res...)
	}
	return ans
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
