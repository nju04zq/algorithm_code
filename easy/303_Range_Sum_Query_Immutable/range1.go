package main

import (
	"fmt"
	"math/rand"
	"time"
)

type NumArray struct {
	buf []int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	a := new(NumArray)
	a.buf = make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			a.buf[i] = nums[i]
		} else {
			a.buf[i] = a.buf[i-1] + nums[i]
		}
	}
	return *a
}

func (a *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return a.buf[j]
	} else {
		return a.buf[j] - a.buf[i-1]
	}
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */

func MakeRandArray() []int {
	maxLen, maxElement := 1000, 10
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len := r.Int()%maxLen + 1
	a := make([]int, len)
	for i := 0; i < len; i++ {
		a[i] = r.Int() % maxElement
	}
	return a
}

func MakeRandInt(a int) int {
	if a == 0 {
		return 0
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int() % a
}

func bf(nums []int, i, j int) int {
	total := 0
	for x := i; x <= j; x++ {
		total += nums[x]
	}
	return total
}

func testSum() {
	nums := MakeRandArray()
	j := MakeRandInt(len(nums))
	i := MakeRandInt(j + 1)
	//nums, i, j = []int{0, 9}, 0, 1
	obj := Constructor(nums)
	res := obj.SumRange(i, j)
	ans := bf(nums, i, j)
	if res != ans {
		panic(fmt.Errorf("Fail on %v, i %d, j %d, get %d, ans %d",
			nums, i, j, res, ans))
	}
}

func main() {
	for i := 0; i < 10000; i++ {
		testSum()
	}
}
