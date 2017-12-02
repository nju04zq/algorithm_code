package main

import "fmt"
import "math"
import "sort"

var pStats = make(map[string]int)
var cStats = make(map[string]int)

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func isPalindrome(s string) bool {
	pStats[s]++
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func cut(s string, total *int, res *int) {
	cStats[s]++
	if *total > *res {
		return
	} else if isPalindrome(s) {
		*res = min(*res, *total)
		return
	}
	(*total)++
	for i := len(s) - 1; i >= 0; i-- {
		if isPalindrome(s[:i]) {
			cut(s[i:], total, res)
		}
	}
	(*total)--
}

func minCut(s string) int {
	var total, res int
	res = math.MaxInt32
	cut(s, &total, &res)
	return res
}

func testMinCut(s string) {
	fmt.Printf("On %q, get %d\n", s, minCut(s))
}

func rankByWordCount(wordFrequencies map[string]int) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	s := "eegiicgaeadbcfacfhi"
	testMinCut(s)
	fmt.Println("======pStats======")
	a := rankByWordCount(pStats)
	for _, p := range a {
		fmt.Printf("%21q, %d\n", p.Key, p.Value)
	}
	fmt.Println("======cStats======")
	a = rankByWordCount(cStats)
	for _, p := range a {
		fmt.Printf("%21q, %d\n", p.Key, p.Value)
	}
	return
	testMinCut("")
	testMinCut("aaa")
	testMinCut("aab")
}
