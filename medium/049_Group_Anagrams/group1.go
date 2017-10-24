package main

import "fmt"
import "sort"

func makeKey(s string) string {
	buf := []byte(s)
	sort.Slice(buf, func(i, j int) bool {
		if buf[i] > buf[j] {
			return true
		} else {
			return false
		}
	})
	return string(buf)
}

func groupAnagrams(strs []string) [][]string {
	res := make(map[string][]string)
	for _, s := range strs {
		key := makeKey(s)
		if _, ok := res[key]; !ok {
			res[key] = []string{s}
		} else {
			res[key] = append(res[key], s)
		}
	}
	ans := make([][]string, 0)
	for _, a := range res {
		ans = append(ans, a)
	}
	return ans
}

func testGroup(strs []string) {
	ans := groupAnagrams(strs)
	fmt.Printf("strs: %v\n", strs)
	fmt.Printf("get: %v\n", ans)
}

func main() {
	testGroup([]string{"eeat", "teea", "tan", "atee", "nat", "bat"})
}
