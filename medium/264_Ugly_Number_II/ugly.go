package main

import "fmt"
import "math"

var roots = []int{2, 3, 5}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func makeTbl(nums []int) map[int]bool {
	tbl := make(map[int]bool)
	for _, num := range nums {
		tbl[num] = true
	}
	return tbl
}

func nthUglyNumber(n int) int {
	nums := []int{1, 2, 3, 4, 5}
	next := []int{2, 2, 4}
	tbl := makeTbl(nums)
	if n <= len(nums) {
		return nums[n-1]
	}
	for i := len(nums); i < n; i++ {
		for j, root := range roots {
			for {
				x := root * nums[next[j]]
				if _, ok := tbl[x]; !ok {
					break
				}
				next[j]++
			}
		}
		a := roots[0] * nums[next[0]]
		b := roots[1] * nums[next[1]]
		c := roots[2] * nums[next[2]]
		//fmt.Println(i, a, b, c)
		minNum := min(a, min(b, c))
		if a == minNum {
			next[0]++
		}
		if b == minNum {
			next[1]++
		}
		if c == minNum {
			next[2]++
		}
		nums = append(nums, minNum)
		tbl[minNum] = true
	}
	return nums[len(nums)-1]
}

func isUgly(num int) bool {
	if num == 0 {
		return false
	} else if num == 1 {
		return true
	}
	for _, root := range roots {
		for num%root == 0 {
			num /= root
		}
	}
	if num == 1 {
		return true
	} else {
		return false
	}
}

func bf(n int) int {
	for i, j := 1, 1; i < math.MaxInt32; i++ {
		if isUgly(i) {
			if j == n {
				return i
			}
			j++
		}
	}
	return -1
}

func testUgly(n int) {
	res := nthUglyNumber(n)
	ans := bf(n)
	if res != ans {
		fmt.Println()
		panic(fmt.Errorf("the %d, get %d, expect %d\n", n, res, ans))
	}
}

func main() {
	for i := 1; i < 1000; i++ {
		fmt.Printf("\r%d", i)
		testUgly(i)
	}
	fmt.Println()
}
