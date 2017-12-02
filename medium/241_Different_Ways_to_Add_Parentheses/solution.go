package main

import "fmt"

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

func computeInternal(nums []int, opers []byte) []int {
	if len(nums) == 1 {
		return []int{nums[0]}
	}
	res := make([]int, 0)
	for i := 1; i < len(nums); i++ {
		res1 := computeInternal(nums[:i], opers[:i-1])
		res2 := computeInternal(nums[i:], opers[i:])
		for _, num1 := range res1 {
			for _, num2 := range res2 {
				num := cal(num1, num2, opers[i-1])
				res = append(res, num)
			}
		}
	}
	return res
}

func diffWaysToCompute(input string) []int {
	nums, opers := split(input)
	return computeInternal(nums, opers)
}

func testSolution(s string) {
	res := diffWaysToCompute(s)
	fmt.Printf("%q, get %v\n", s, res)
}

func main() {
	testSolution("2*3-4*5")
	testSolution("2-1-1")
}
