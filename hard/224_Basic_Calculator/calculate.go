package main

import "fmt"
import "strconv"

const (
	tokenNum = iota
	tokenOperator
)

const (
	ADD = iota
	SUB
	MUL
	DIV
	LPAREN
	RPAREN
)

var priority = map[int]int{
	RPAREN: 3,
	MUL:    2,
	DIV:    2,
	ADD:    1,
	SUB:    1,
	LPAREN: 0,
}

type Token struct {
	tokenType int
	val       int
}

func (t *Token) eval(a, b int) int {
	switch t.val {
	case ADD:
		return a + b
	case SUB:
		return a - b
	case MUL:
		return a * b
	case DIV:
		return a / b
	default:
		return 0
	}
}

func (t *Token) str() string {
	if t.tokenType == tokenNum {
		return fmt.Sprintf("%d", t.val)
	}
	switch t.val {
	case ADD:
		return "+"
	case SUB:
		return "-"
	case MUL:
		return "*"
	case DIV:
		return "/"
	case LPAREN:
		return "("
	case RPAREN:
		return ")"
	}
	return "#"
}

func makeToken(s string) *Token {
	if len(s) > 1 || isDigit(s[0]) {
		val, _ := strconv.ParseInt(s, 10, 32)
		return &Token{tokenNum, int(val)}
	}
	var val int
	switch s[0] {
	case '+':
		val = ADD
	case '-':
		val = SUB
	case '*':
		val = MUL
	case '/':
		val = DIV
	case '(':
		val = LPAREN
	case ')':
		val = RPAREN
	}
	return &Token{tokenOperator, val}
}

func isDigit(d byte) bool {
	if d >= '0' && d <= '9' {
		return true
	} else {
		return false
	}
}

func skipSpaces(s string, start int) int {
	i := start
	for ; i < len(s); i++ {
		if s[i] != ' ' {
			break
		}
	}
	return i
}

func getToken(s string, start int, prev *Token) (string, int) {
	//fmt.Println(s[start:])
	negative := false
	i := skipSpaces(s, start)
	if i >= len(s) {
		return "", i
	}
	start = i
	if prev != nil && prev.tokenType == tokenOperator && prev.val != RPAREN {
		if s[i] == '+' {
			i++
			i = skipSpaces(s, i)
			start = i
		} else if s[i] == '-' {
			negative = true
			i++
			i = skipSpaces(s, i)
			start = i
		}
	}
	for ; i < len(s); i++ {
		if !isDigit(s[i]) {
			break
		}
	}
	if i > start {
		if negative {
			return "-" + s[start:i], i
		} else {
			return s[start:i], i
		}
	}
	return s[i : i+1], i + 1
}

func split(s string) []*Token {
	var tokStr string
	var prev *Token
	toks := make([]*Token, 0)
	for i := 0; i < len(s); {
		tokStr, i = getToken(s, i, prev)
		if tokStr != "" {
			tok := makeToken(tokStr)
			toks = append(toks, tok)
			prev = tok
		}
	}
	return toks
}

func rpn(toks []*Token) []*Token {
	var tok1 *Token
	sRes := make([]*Token, 0)
	sOpe := make([]*Token, 0)
	for _, tok := range toks {
		if tok.tokenType == tokenNum {
			sRes = append(sRes, tok)
		} else if tok.val == LPAREN {
			sOpe = append(sOpe, tok)
		} else if tok.val == RPAREN {
			for len(sOpe) > 0 {
				cnt := len(sOpe)
				tok1, sOpe = sOpe[cnt-1], sOpe[:cnt-1]
				if tok1.val == LPAREN {
					break
				} else {
					sRes = append(sRes, tok1)
				}
			}
		} else {
			for len(sOpe) > 0 {
				cnt := len(sOpe)
				tok1 = sOpe[cnt-1]
				if priority[tok.val] > priority[tok1.val] {
					break
				}
				sOpe = sOpe[:cnt-1]
				sRes = append(sRes, tok1)
			}
			sOpe = append(sOpe, tok)
		}
	}
	for len(sOpe) > 0 {
		cnt := len(sOpe)
		tok1, sOpe = sOpe[cnt-1], sOpe[:cnt-1]
		sRes = append(sRes, tok1)
	}
	return sRes
}

func evalRPN(toks []*Token) int {
	var a, b, cnt int
	stack := make([]int, 0)
	for _, tok := range toks {
		//fmt.Println(stack)
		if tok.tokenType == tokenNum {
			stack = append(stack, tok.val)
			continue
		}
		cnt = len(stack)
		b, stack = stack[cnt-1], stack[:cnt-1]
		cnt = len(stack)
		a, stack = stack[cnt-1], stack[:cnt-1]
		res := tok.eval(a, b)
		stack = append(stack, res)
	}
	return stack[len(stack)-1]
}

func dumpTokens(toks []*Token) {
	for _, tok := range toks {
		fmt.Printf("%s ", tok.str())
	}
	fmt.Println()
}

func calculate(s string) int {
	toks := split(s)
	//dumpTokens(toks)
	toks = rpn(toks)
	//dumpTokens(toks)
	res := evalRPN(toks)
	return res
}

func testCalculate(s string) {
	fmt.Printf("%q, get %d\n", s, calculate(s))
}

func main() {
	//2
	testCalculate(" ( 2 + 3*4 )  / 5  ")
	//3
	testCalculate(" 1 + ( 2 * 3 + 4 )  / 5  ")
	//1
	testCalculate(" 1 + ( 2 * 3 + - 4 )  / 5  ")
	//-4
	testCalculate("1-(5)")
	//4
	testCalculate("(3)+1")
}
