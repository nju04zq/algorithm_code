package main

import "fmt"

var pairLR = []byte{'(', ')'}
var pairRL = []byte{')', '('}

func reverse(s string) string {
	buf := []byte(s)
	i, j := 0, len(buf)-1
	for i < j {
		buf[i], buf[j] = buf[j], buf[i]
		i++
		j--
	}
	return string(buf)
}

func remove(s string, last_i, last_j int, ans []string, pair []byte) []string {
	stack := 0
	for i := last_i; i < len(s); i++ {
		if s[i] == pair[0] {
			stack++
		} else if s[i] == pair[1] {
			stack--
		}
		if stack >= 0 {
			continue
		}
		for j := last_j; j <= i; j++ {
			if s[j] == pair[1] && (j == last_j || s[j-1] != pair[1]) {
				s1 := s[:j] + s[j+1:]
				ans = remove(s1, i, j, ans, pair)
			}
		}
		return ans
	}
	reversed := reverse(s)
	if pair[0] == '(' {
		ans = remove(reversed, 0, 0, ans, pairRL)
	} else {
		ans = append(ans, reverse(s))
	}
	return ans
}

func removeInvalidParentheses(s string) []string {
	ans := make([]string, 0)
	ans = remove(s, 0, 0, ans, pairLR)
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
