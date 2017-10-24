package main

import "fmt"

func isValid(s string) bool {
	var rLeft rune
	stack := make([]rune, len(s))
	i := -1
	for _, r := range s {
		if r == '(' || r == '{' || r == '[' {
			i++
			stack[i] = r
			continue
		}
		switch r {
		case ')':
			rLeft = '('
		case '}':
			rLeft = '{'
		case ']':
			rLeft = '['
		default:
			return false
		}
		if i < 0 {
			return false
		}
		if stack[i] != rLeft {
			return false
		}
		i--
	}
	if i >= 0 {
		return false
	} else {
		return true
	}
}

func testIsValid(s string, ans bool) {
	res := isValid(s)
	if res != ans {
		panic(fmt.Errorf("s %q, get %t, ans %t", s, res, ans))
	}
}

func main() {
	testIsValid("()", true)
	testIsValid("{}", true)
	testIsValid("[]", true)
	testIsValid("()[]{}", true)
	testIsValid("([{}])", true)
	testIsValid("(()[{}{}])", true)
	testIsValid("(()[{}{})", false)
	testIsValid("(()[{}{}]", false)
	testIsValid("(]", false)
	testIsValid("([)]", false)
}
