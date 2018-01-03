package main

import "fmt"

func dfs(num string, target, eval, mul int, path string, res []string) []string {
	if len(num) == 0 {
		if target == eval {
			res = append(res, path)
		}
		return res
	}
	j := 0
	for i := 0; i < len(num); i++ {
		if i > 0 && num[0] == '0' {
			break
		}
		j = j*10 + int(num[i]-'0')
		cur := num[:i+1]
		if len(path) == 0 {
			res = dfs(num[i+1:], target, eval+j, j, path+cur, res)
		} else {
			res = dfs(num[i+1:], target, eval+j, j, path+"+"+cur, res)
			res = dfs(num[i+1:], target, eval-j, -j, path+"-"+cur, res)
			res = dfs(num[i+1:], target, eval-mul+mul*j, mul*j, path+"*"+cur, res)
		}
	}
	return res
}

func addOperators(num string, target int) []string {
	if len(num) == 0 {
		return []string{}
	}
	res := make([]string, 0)
	res = dfs(num, target, 0, 0, "", res)
	return res
}

func testAdd(num string, target int) {
	fmt.Printf("%q, target %d, get %v\n", num, target, addOperators(num, target))
}

func main() {
	testAdd("123", 6)
	testAdd("232", 8)
	testAdd("105", 5)
	testAdd("00", 0)
}
