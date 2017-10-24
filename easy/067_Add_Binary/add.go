package main

import "fmt"

func addBinary(a string, b string) string {
	add := func(a, b, carry byte) (byte, byte) {
		a0 := a - '0'
		b0 := b - '0'
		c0 := carry - '0'
		sum := a0 + b0 + c0
		c0 = sum / 2
		sum = sum % 2
		return sum + '0', c0 + '0'
	}
	res := []byte{}
	i, j := len(a)-1, len(b)-1
	sum, carry := byte('0'), byte('0')
	for i >= 0 && j >= 0 {
		sum, carry = add(a[i], b[j], carry)
		res = append(res, sum)
		i--
		j--
	}
	for ; i >= 0; i-- {
		sum, carry = add(a[i], '0', carry)
		res = append(res, sum)
	}
	for ; j >= 0; j-- {
		sum, carry = add(b[j], '0', carry)
		res = append(res, sum)
	}
	if carry != '0' {
		res = append(res, carry)
	}
	i, j = 0, len(res)-1
	for i <= j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return string(res)
}

func testAdd(a, b string) {
	c := addBinary(a, b)
	fmt.Printf("a %q, b %q, c %q\n", a, b, c)
}

func main() {
	testAdd("11", "1")
}
