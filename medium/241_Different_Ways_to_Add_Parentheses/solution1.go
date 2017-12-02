package main

import "fmt"
import "bytes"

func isDigit(ch byte) bool {
	if ch >= '0' && ch <= '9' {
		return true
	} else {
		return false
	}
}

func split(s string) ([]int, []byte) {
	nums, opers := make([]int, 0), make([]byte, 0)
	for i := 0; i < len(s); {
		if !isDigit(s[i]) {
			opers = append(opers, s[i])
			i++
			continue
		}
		num := 0
		for ; i < len(s) && isDigit(s[i]); i++ {
			num = num*10 + int(s[i]-'0')
		}
		nums = append(nums, num)
	}
	return nums, opers
}

func cal(a, b int, oper byte) int {
	switch oper {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	}
	return 0
}

func compose(nums []int, opers []byte) string {
	buf := bytes.NewBuffer(nil)
	for i := 0; i < len(nums); i++ {
		buf.WriteString(fmt.Sprintf("%d", nums[i]))
		if i < len(opers) {
			buf.WriteByte(opers[i])
		}
	}
	return buf.String()
}

func computeInternal(nums []int, opers []byte, dp map[string][]int) []int {
	if len(nums) == 1 {
		return []int{nums[0]}
	}
	s := compose(nums, opers)
	if res, ok := dp[s]; ok {
		return res
	}
	res := make([]int, 0)
	for i := 1; i < len(nums); i++ {
		res1 := computeInternal(nums[:i], opers[:i-1], dp)
		res2 := computeInternal(nums[i:], opers[i:], dp)
		for _, num1 := range res1 {
			for _, num2 := range res2 {
				num := cal(num1, num2, opers[i-1])
				res = append(res, num)
			}
		}
	}
	dp[s] = res
	return res
}

func diffWaysToCompute(input string) []int {
	nums, opers := split(input)
	dp := make(map[string][]int)
	return computeInternal(nums, opers, dp)
}

func testSolution(s string) {
	res := diffWaysToCompute(s)
	fmt.Printf("%q, get %v\n", s, res)
}

func main() {
	testSolution("2*3-4*5")
	testSolution("2-1-1")
}
