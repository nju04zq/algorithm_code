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

func calc(s string, i, total, prev int) int {
	i = skipSpaces(s, i)
	if i == len(s) {
		return total
	}
	op := s[i]
	i++
	i = skipSpaces(s, i)
	num, i := getNum(s, i)
	if op == '+' {
		return calc(s, i, total+num, num)
	} else if op == '-' {
		return calc(s, i, total-num, -num)
	} else {
		total -= prev
		if op == '*' {
			num = prev * num
		} else {
			num = prev / num
		}
		return calc(s, i, total+num, num)
	}
}

func calculate(s string) int {
	i := skipSpaces(s, 0)
	num, i := getNum(s, i)
	return calc(s, i, num, num)
}

func testCalc(s string) {
	fmt.Printf("%s, get %d\n", s, calculate(s))
}

func main() {
	testCalc("3+2*2")
	testCalc(" 3/2 ")
	testCalc(" 3+5 / 2 ")
}
