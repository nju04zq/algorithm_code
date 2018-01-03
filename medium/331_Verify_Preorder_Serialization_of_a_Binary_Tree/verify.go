package main

import "fmt"
import "strings"

type Stack struct {
	buf []int
}

func (s *Stack) init() *Stack {
	s.buf = make([]int, 0)
	return s
}

func (s *Stack) size() int {
	return len(s.buf)
}

func (s *Stack) push(val int) {
	s.buf = append(s.buf, val)
}

func (s *Stack) pop() int {
	if len(s.buf) == 0 {
		return -1
	}
	val := s.buf[len(s.buf)-1]
	s.buf = s.buf[:len(s.buf)-1]
	return val
}

func (s *Stack) top() int {
	if len(s.buf) == 0 {
		return -1
	}
	return s.buf[len(s.buf)-1]
}

func (s *Stack) incTop() {
	if len(s.buf) == 0 {
		return
	}
	s.buf[len(s.buf)-1]++
}

func isValidSerialization(preorder string) bool {
	if len(preorder) == 0 {
		return true
	}
	toks := strings.Split(preorder, ",")
	if toks[0] == "#" {
		if len(toks) == 1 {
			return true
		} else {
			return false
		}
	}
	stack := new(Stack).init()
	stack.push(0)
	for i := 1; i < len(toks); i++ {
		//fmt.Println(toks[i], stack.buf)
		if stack.size() == 0 {
			return false
		}
		stack.incTop()
		if toks[i] != "#" {
			stack.push(0)
			continue
		}
		if stack.top() == 2 {
			for stack.top() == 2 {
				stack.pop()
			}
		}
	}
	if stack.size() == 0 {
		return true
	} else {
		return false
	}
}

func testIsValid(s string) {
	fmt.Printf("%q, get %t\n", s, isValidSerialization(s))
}

func main() {
	testIsValid("9,3,4,#,#,1,#,#,2,#,6,#,#")
	testIsValid("1,#")
	testIsValid("1,#,#")
	testIsValid("1,#,#,#")
	testIsValid("9,#,#,1")
	testIsValid("#")
}
