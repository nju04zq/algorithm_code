package main

import "fmt"
import "sort"
import "math/rand"
import "time"

// n = 3*m1 + m2, 0 <= m2 <= 2
// n1 > n/3, n1 > m1
// two such elements, total count > 2*(m1), that's >= 2*(m1+1) = 2*m1+2
// left elements, <= m1 + m2 -2, which <= m1
func majorityElement(nums []int) []int {
	var num0, num1 int
	cnt0, cnt1 := 0, 0
	for _, num := range nums {
		if cnt0 > 0 && num == num0 {
			cnt0++
		} else if cnt1 > 0 && num == num1 {
			cnt1++
		} else if cnt0 == 0 {
			num0, cnt0 = num, 1
		} else if cnt1 == 0 {
			num1, cnt1 = num, 1
		} else {
			cnt0--
			cnt1--
		}
	}
	n0, n1 := 0, 0
	for _, num := range nums {
		if num == num0 {
			n0++
		} else if num == num1 {
			n1++
		}
	}
	res := make([]int, 0)
	if n0 > len(nums)/3 {
		res = append(res, num0)
	}
	if n1 > len(nums)/3 {
		res = append(res, num1)
	}
	return res
}

func bf(nums []int) []int {
	res := make([]int, 0)
	tbl := make(map[int]int)
	for _, num := range nums {
		if _, ok := tbl[num]; !ok {
			tbl[num] = 1
		} else {
			tbl[num]++
		}
	}
	for num, cnt := range tbl {
		if cnt > len(nums)/3 {
			res = append(res, num)
		}
	}
	return res
}

func testMajority(nums []int) {
	//nums = []int{0, 9, 9, 9, 15, 14, 7, 6}
	nums = []int{17, 1, 1, 10, 1, 14, 4, 12}
	res := majorityElement(nums)
	ans := bf(nums)
	sort.Ints(res)
	sort.Ints(ans)
	if len(res) != len(ans) {
		panic(fmt.Sprintf("%v, get %v, expect %v\n", nums, res, ans))
	}
	for i := 0; i < len(res); i++ {
		if res[i] != ans[i] {
			panic(fmt.Sprintf("%v, get %v, expect %v\n", nums, res, ans))
		}
	}
}

func MakeRandArray() []int {
	maxLen, maxElement := 10, 20
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len := r.Int() % maxLen
	a := make([]int, len)
	for i := 0; i < len; i++ {
		a[i] = r.Int() % maxElement
	}
	return a
}

func main() {
	for i := 0; i < 10000; i++ {
		a := MakeRandArray()
		testMajority(a)
	}
}
