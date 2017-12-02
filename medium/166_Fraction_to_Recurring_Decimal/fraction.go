package main

import "fmt"

func fractionToDecimal(a int, b int) string {
	sign := ""
	if a < 0 && b < 0 {
		a, b = -a, -b
	} else if a < 0 {
		a = -a
		sign = "-"
	} else if b < 0 {
		b = -b
		if a != 0 {
			sign = "-"
		}
	}
	x := a / b
	a = a % b
	if a == 0 {
		return fmt.Sprintf("%s%d", sign, x)
	}
	tbl := make(map[int]int)
	digits := make([]byte, 0)
	i := 0
	for a != 0 {
		if _, ok := tbl[a]; ok {
			break
		} else {
			tbl[a] = i
		}
		a *= 10
		digits = append(digits, '0'+byte(a/b))
		i++
		if a >= b {
			a = a % b
		}
	}
	if a == 0 {
		return fmt.Sprintf("%s%d.%s", sign, x, string(digits))
	} else {
		i := tbl[a]
		return fmt.Sprintf("%s%d.%s(%s)", sign, x,
			string(digits[:i]),
			string(digits[i:]))
	}
}

func testFraction(a, b int) {
	fmt.Printf("%d/%d, get %s\n", a, b, fractionToDecimal(a, b))
}

func main() {
	testFraction(1, 2)
	testFraction(2, 1)
	testFraction(2, 3)
	testFraction(1, 6)
	testFraction(7, 6)
	testFraction(1, 100)
	testFraction(1, 99)
	testFraction(1, 90)
	testFraction(1, 900)
	testFraction(1, 990)
	testFraction(1, 999)
	testFraction(-1, 999)
	testFraction(-1, -999)
	testFraction(0, -999)
}
