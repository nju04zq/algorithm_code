package main

import "fmt"

func getRow(rowIndex int) []int {
	rowIndex++
	res := make([]int, rowIndex)
	res[0] = 1
	for i := 2; i <= rowIndex; i++ {
		var j, prev, temp int
		for j = 0; j < i-1; j++ {
			temp = res[j]
			res[j] += prev
			prev = temp
		}
		res[j] = 1
	}
	return res
}

func testGet(rowIndex int) {
	fmt.Printf("%d, get %v\n", rowIndex, getRow(rowIndex))
}

func main() {
	testGet(0)
	testGet(1)
	testGet(2)
	testGet(3)
	testGet(4)
	testGet(5)
}
