package main

import "fmt"

func generate(numRows int) [][]int {
	res := make([][]int, 0)
	if numRows == 0 {
		return res
	}
	last := []int{1}
	res = append(res, last)
	for i := 2; i <= numRows; i++ {
		var j int
		cur := make([]int, i)
		cur[0] = 1
		for j = 1; j < len(last); j++ {
			cur[j] = last[j-1] + last[j]
		}
		cur[j] = 1
		res = append(res, cur)
		last = cur
	}
	return res
}

func testGenerate(k int) {
	res := generate(k)
	fmt.Printf("%d, get:\n", k)
	fmt.Println(res)
}

func main() {
	testGenerate(1)
	testGenerate(2)
	testGenerate(3)
	testGenerate(4)
}
