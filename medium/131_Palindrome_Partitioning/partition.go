package main

import "fmt"

func isPalindrome(s string) bool {
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

func makeCopy(src []string) []string {
	dst := make([]string, len(src))
	for i, s := range src {
		dst[i] = s
	}
	return dst
}

func partitionInternal(s string, path []string, res [][]string) [][]string {
	if len(s) == 0 {
		res = append(res, makeCopy(path))
		return res
	}
	for i := 0; i < len(s); i++ {
		part := s[:i+1]
		if isPalindrome(part) {
			path = append(path, part)
			res = partitionInternal(s[i+1:], path, res)
			path = path[:len(path)-1]
		}
	}
	return res
}

func partition(s string) [][]string {
	path := make([]string, 0)
	res := make([][]string, 0)
	res = partitionInternal(s, path, res)
	return res
}

func testPartition(s string) {
	res := partition(s)
	fmt.Printf("%q, get %v\n", s, res)
}

func main() {
	testPartition("aab")
}
