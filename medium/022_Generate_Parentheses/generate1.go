package main

import "fmt"

func generate(left, right int, path []byte, ans []string) []string {
	if left > right {
		return ans
	}
	if left == 0 && right == 0 {
		ans = append(ans, string(path))
		return ans
	}
	if left > 0 {
		path1 := append(path, '(')
		ans = generate(left-1, right, path1, ans)
	}
	if right > 0 {
		path1 := append(path, ')')
		ans = generate(left, right-1, path1, ans)
	}
	return ans
}

func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	path := make([]byte, 0)
	ans = generate(n, n, path, ans)
	return ans
}

func testGenerateParenthesis(n int) {
	ans := generateParenthesis(n)
	fmt.Println(n, ans)
}

func main() {
	testGenerateParenthesis(0)
	testGenerateParenthesis(1)
	testGenerateParenthesis(2)
	testGenerateParenthesis(3)
}
