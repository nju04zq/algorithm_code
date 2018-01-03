package main

import "fmt"

func isDigit(d byte) bool {
	if d >= '0' && d <= '9' {
		return true
	} else {
		return false
	}
}

func skipSpaces(s string, i int) int {
	for i < len(s) && s[i] == ' ' {
		i++
	}
	return i
}

func getNum(s string, i int) (int, int) {
	num := 0
	for i < len(s) && isDigit(s[i]) {
		num = num*10 + int(s[i]-'0')
		i++
	}
	return num, i
}

func calculate(s string) int {
	var num1, num2 int
	sign := 1
	stack := make([]int, 0)
	for i := 0; i < len(s); {
		if s[i] == ' ' {
			i = skipSpaces(s, i)
		} else if isDigit(s[i]) {
			num1, i = getNum(s, i)
			stack = append(stack, sign*num1)
		} else if s[i] == '+' {
			sign = 1
			i++
		} else if s[i] == '-' {
			sign = -1
			i++
		} else {
			op := s[i]
			num1, stack = stack[len(stack)-1], stack[:len(stack)-1]
			i = skipSpaces(s, i+1)
			num2, i = getNum(s, i)
			if op == '*' {
				stack = append(stack, num1*num2)
			} else {
				stack = append(stack, num1/num2)
			}
		}
	}
	sum := 0
	for _, num := range stack {
		sum += num
	}
	return sum
}

func testCalc(s string) {
	fmt.Printf("%s, get %d\n", s, calculate(s))
}

func main() {
	testCalc("3+2*2")
	testCalc(" 3/2 ")
	testCalc(" 3+5 / 2 ")
}
