package main

import "fmt"
import "strings"

func dfs(s string, dict []string, record map[string][]string) []string {
	res := make([]string, 0)
	if _, ok := record[s]; ok {
		return record[s]
	} else if len(s) == 0 {
		res = append(res, "")
		return res
	}
	for _, word := range dict {
		if !strings.HasPrefix(s, word) {
			continue
		}
		sublist := dfs(s[len(word):], dict, record)
		if len(sublist) == 0 {
			continue
		}
		for _, sub := range sublist {
			if len(sub) == 0 {
				res = append(res, fmt.Sprintf("%s", word))
			} else {
				res = append(res, fmt.Sprintf("%s %s", word, sub))
			}
		}
	}
	record[s] = res
	return res
}

func wordBreak(s string, wordDict []string) []string {
	record := make(map[string][]string)
	return dfs(s, wordDict, record)
}

func testBreak(s string, wordDict []string) {
	res := wordBreak(s, wordDict)
	fmt.Printf("%q, dict %v, get:\n", s, wordDict)
	for _, s := range res {
		fmt.Println(s)
	}
}

func main() {
	testBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"})
}
