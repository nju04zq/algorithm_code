package main

import "fmt"
import "strings"

func dfs(s string, dp [][]int, i int, res, path []string) []string {
	if i == len(s) {
		res = append(res, strings.Join(path, " "))
		return res
	}
	pos := dp[i]
	for j := 0; j < len(pos); j++ {
		k := pos[j]
		path = append(path, s[i:k])
		res = dfs(s, dp, k, res, path)
		path = path[:len(path)-1]
	}
	return res
}

func wordBreak(s string, wordDict []string) []string {
	tbl := make(map[string]bool)
	for _, word := range wordDict {
		tbl[word] = true
	}
	dp := make([][]int, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		pos := make([]int, 0)
		for j := i + 1; j <= len(s); j++ {
			if _, ok := tbl[s[i:j]]; !ok {
				continue
			}
			if j == len(s) || len(dp[j]) > 0 {
				pos = append(pos, j)
			}
		}
		dp[i] = pos
	}
	res := make([]string, 0)
	if len(dp[0]) == 0 {
		return res
	}
	path := make([]string, 0)
	res = dfs(s, dp, 0, path, res)
	return res
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
