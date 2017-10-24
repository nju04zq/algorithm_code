package main

import "fmt"

func generate(left, right int, path []byte, ans []string) []string {
	if left == 0 && right == 0 {
		ans = append(ans, string(path))
		return ans
	}
	for i := left; i > 0; i-- {
		path = append(path, '(')
		mark := len(path)
		for j := right; j >= i; j-- {
			path = append(path, ')')
			ans = generate(i-1, j-1, path, ans)
		}
		path = path[:mark]
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
