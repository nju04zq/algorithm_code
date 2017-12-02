package main

import "fmt"
import "strconv"

const (
	NIL = iota
	ADD
	SUB
	MUL
	DIV
)

func operate(num1, num2, operator int) int {
	switch operator {
	case ADD:
		return num1 + num2
	case SUB:
		return num1 - num2
	case MUL:
		return num1 * num2
	case DIV:
		return num1 / num2
	}
	return 0
}

func evalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}
	var num1, num2, num int
	stack := make([]int, 0)
	for _, s := range tokens {
		operator := NIL
		switch s {
		case "+":
			operator = ADD
		case "-":
			operator = SUB
		case "*":
			operator = MUL
		case "/":
			operator = DIV
		default:
			i, _ := strconv.ParseInt(s, 10, 32)
			num = int(i)
			stack = append(stack, num)
		}
		if operator != NIL {
			num1, stack = stack[len(stack)-1], stack[:len(stack)-1]
			num2, stack = stack[len(stack)-1], stack[:len(stack)-1]
			num = operate(num2, num1, operator)
			stack = append(stack, num)
		}
	}
	return stack[len(stack)-1]
}

func testRPN(toks []string) {
	fmt.Printf("For %v, get %d\n", toks, evalRPN(toks))
}

func main() {
	testRPN([]string{"2", "1", "+", "3", "*"})
	testRPN([]string{"4", "13", "5", "/", "+"})
}
