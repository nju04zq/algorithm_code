package main

import "fmt"

func isValid(s string) bool {
	stack := 0
	for _, c := range s {
		if c == '(' {
			stack++
		} else if c == ')' {
			stack--
		}
		if stack < 0 {
			return false
		}
	}
	if stack > 0 {
		return false
	} else {
		return true
	}
}

func removeInvalidParentheses(s string) []string {
	ans := make([]string, 0)
	queue := make([]string, 0)
	visited := make(map[string]bool)
	queue = append(queue, s)
	for len(queue) > 0 {
		queue0 := make([]string, 0)
		for _, s := range queue {
			if _, ok := visited[s]; ok {
				continue
			}
			visited[s] = true
			if isValid(s) {
				ans = append(ans, s)
			}
			for i, c := range s {
				if c == '(' || c == ')' {
					s1 := s[:i] + s[i+1:]
					queue0 = append(queue0, s1)
				}
			}
		}
		if len(ans) > 0 {
			break
		}
		queue = queue0
	}
	if len(ans) == 0 {
		return []string{""}
	} else {
		return ans
	}
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
