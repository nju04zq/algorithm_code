package main

import "fmt"

func stripSpaces(s string) string {
	start, end, i := 0, 0, 0
	for i = 0; i < len(s); i++ {
		if s[i] != ' ' {
			start = i
			break
		}
	}
	start = i
	for i = len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			end = i
			break
		}
	}
	end = i
	if start >= len(s) {
		return ""
	} else {
		return s[start : end+1]
	}
}

const (
	NUM = iota
	DOT
	SIG
	EXP
	NIL
	ALL
)

var matrix = [][]int{
	//    N  D  S  E  L
	[]int{1, 1, 0, 1, 1}, //NUM
	[]int{1, 0, 0, 1, 1}, //DOT
	[]int{1, 1, 0, 0, 0}, //SIG
	[]int{1, 0, 1, 0, 0}, //EXP
	[]int{0, 0, 0, 0, 0}, //NIL
}

func classify(c byte) int {
	if c >= '0' && c <= '9' {
		return NUM
	} else if c == '.' {
		return DOT
	} else if c == '+' {
		return SIG
	} else if c == '-' {
		return SIG
	} else if c == 'e' {
		return EXP
	} else if c == '0' {
		return NIL
	} else {
		return -1
	}
}

type context struct {
	cnts []int
}

func (c *context) init() *context {
	c.cnts = make([]int, ALL)
	return c
}

type handler func(c *context) bool

var handlers = []handler{
	nil,        //NUM
	dotHandler, //DOT
	sigHandler, //SIG
	expHandler, //EXP
	nilHandler, //NIL
}

func dotHandler(c *context) bool {
	if c.cnts[EXP] >= 1 {
		return false
	}
	if c.cnts[DOT] > 1 {
		return false
	} else {
		return true
	}
}

func sigHandler(c *context) bool {
	return true
}

func expHandler(c *context) bool {
	if c.cnts[NUM] == 0 {
		return false
	}
	if c.cnts[EXP] > 1 {
		return false
	} else {
		return true
	}
}

func nilHandler(c *context) bool {
	if c.cnts[NUM] == 0 {
		return false
	} else {
		return true
	}
}

func validTransitiion(prev, cur int) bool {
	if matrix[prev][cur] == 1 {
		return true
	} else {
		return false
	}
}

func judge(c *context, prev, cur int) bool {
	if prev != -1 && !validTransitiion(prev, cur) {
		return false
	}
	h := handlers[cur]
	if h != nil && !h(c) {
		return false
	}
	return true
}

func isNumber(s string) bool {
	s = stripSpaces(s)
	if len(s) == 0 {
		return false
	}
	prev, cur := -1, -1
	c := new(context).init()
	for i, _ := range s {
		cur := classify(s[i])
		if cur == -1 {
			return false
		}
		c.cnts[cur]++
		if !judge(c, prev, cur) {
			return false
		}
		prev = cur
	}
	cur = NIL
	if !judge(c, prev, cur) {
		return false
	}
	return true

}

func testValid(s string, ans bool) {
	res := isNumber(s)
	if res != ans {
		panic(fmt.Errorf("%q, get %t, should %t", s, res, ans))
	}
}

func main() {
	testValid(" 123  ", true)
	testValid(" 1  ", true)
	testValid(" +1  ", true)
	testValid(" -1  ", true)
	testValid(".1", true)
	testValid("1.1", true)
	testValid("-1.1", true)
	testValid("1.", true)
	testValid(".", false)
	testValid("1e1", true)
	testValid("1e-1", true)
	testValid("1e-1.0", false)
	testValid("+1e-1", true)
	testValid("e1", false)
}
