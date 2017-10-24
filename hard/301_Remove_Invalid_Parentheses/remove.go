package main

import "fmt"

// removeLeft can be combined with removeRight
func removeLeft(s string, last_j int, ans []string) []string {
	right := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' {
			right++
		} else if s[i] == '(' {
			right--
		}
		if right >= 0 {
			continue
		}
		for j := last_j; j >= i; j-- {
			if s[j] == '(' && (j == last_j || s[j+1] != '(') {
				s1 := s[:j] + s[j+1:]
				// should be j-1, since ( removed on the right side
				ans = removeLeft(s1, j-1, ans)
			}
		}
		return ans
	}
	ans = append(ans, s)
	return ans
}

// last_j is necessary, to remove duplicate results
// otherwise might remove two ) in different orders,
// like remove 1st then 2nd, and remove 2nd then the 1st
func removeRight(s string, last_j int, ans []string) []string {
	left := 0
	// do not need to count i from 0, because duplicate work has been done
	// previous part has been validated
	for i, c := range s {
		if c == '(' {
			left++
		} else if c == ')' {
			left--
		}
		if left >= 0 {
			continue
		}
		for j := last_j; j <= i; j++ {
			// s[j-1] != ')' is necessary to remove duplcates
			if s[j] == ')' && (j == last_j || s[j-1] != ')') {
				s1 := s[0:j] + s[j+1:]
				// no need to pass in j-1, since a ')' already removed
				ans = removeRight(s1, j, ans)
			}
		}
		return ans
	}
	ans = append(ans, s)
	return ans
}

func removeInvalidParentheses(s string) []string {
	ans0 := make([]string, 0)
	ans0 = removeRight(s, 0, ans0)
	ans := make([]string, 0)
	for _, s := range ans0 {
		ans = removeLeft(s, len(s)-1, ans)
	}
	return ans
}

func testRemove(s string) {
	result := removeInvalidParentheses(s)
	fmt.Println(s, result)
}

func main() {
	testRemove("")
	testRemove("(")
	testRemove(")")
	testRemove("(((k()((")
	testRemove("(()")
	testRemove("()())()")
	testRemove("(a)())()")
	testRemove(")(")
}
