package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func compare(a, b, c string) int {
	d := make([]byte, max(len(a), len(b))+1)
	i, j, k := len(a)-1, len(b)-1, 0
	var x, y, z, carry int
	for i >= 0 || j >= 0 {
		if i >= 0 {
			x = int(a[i] - '0')
			i--
		} else {
			x = 0
		}
		if j >= 0 {
			y = int(b[j] - '0')
			j--
		} else {
			y = 0
		}
		z = x + y + carry
		if z >= 10 {
			z -= 10
			carry = 1
		} else {
			carry = 0
		}
		d[k] = byte(z + '0')
		k++
	}
	if carry > 0 {
		d[k] = byte(carry + '0')
		k++
	}
	i, j = 0, k-1
	for i < j {
		d[i], d[j] = d[j], d[i]
		i++
		j--
	}
	d = d[:k]
	if len(d) < len(c) {
		return -1
	} else if len(d) > len(c) {
		return 1
	}
	for i = 0; i < len(d); i++ {
		if d[i] < c[i] {
			return -1
		} else if d[i] > c[i] {
			return 1
		}
	}
	return 0
}

func dfs(num string, path []string) bool {
	if len(num) == 0 && len(path) >= 3 {
		return true
	}
	n := len(path)
	last0 := ""
	last1 := ""
	if n > 0 {
		last0 = path[n-1]
	}
	if n > 1 {
		last1 = path[n-2]
	}
	for i := 0; i < len(num); i++ {
		if i > 0 && num[0] == '0' {
			break
		}
		cur := num[:i+1]
		if len(last0) > 0 && len(last1) > 0 {
			rc := compare(last0, last1, cur)
			//fmt.Println("compare", last0, last1, cur, "rc", rc)
			if rc < 0 {
				break
			} else if rc > 0 {
				continue
			}
		}
		path = append(path, cur)
		if dfs(num[i+1:], path) {
			return true
		}
		path = path[:len(path)-1]

	}
	return false
}

func isAdditiveNumber(num string) bool {
	path := make([]string, 0)
	return dfs(num, path)
}

func testAdd(num string) {
	fmt.Printf("%q, get %t\n", num, isAdditiveNumber(num))
}

func main() {
	testAdd("112358")
	testAdd("199100199")
	testAdd("1203")
	testAdd("199111992")
}
