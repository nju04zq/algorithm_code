package main

import (
	"fmt"
	"math/rand"
	"time"
)

type NumArray struct {
	nums []int
	buf  []int
}

func copyArray(nums []int) []int {
	nums1 := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		nums1[i] = nums[i]
	}
	return nums1
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	a := new(NumArray)
	a.buf = make([]int, n+1)
	nums1 := copyArray(nums)
	for i := 1; i < n; i++ {
		nums1[i] += nums1[i-1]
	}
	for i := 1; i <= n; i++ {
		j := i - (i & -i)
		if j == 0 {
			a.buf[i] = nums1[i-1]
		} else {
			a.buf[i] = nums1[i-1] - nums1[j-1]
		}
	}
	a.nums = copyArray(nums)
	return *a
}

func (a *NumArray) sum(i int) int {
	if i <= 0 {
		return 0
	}
	total := 0
	for i > 0 {
		total += a.buf[i]
		i = i - (i & -i)
	}
	return total
}

func (a *NumArray) SumRange(i int, j int) int {
	return a.sum(j+1) - a.sum(i)
}

func (a *NumArray) Update(i int, val int) {
	diff := val - a.nums[i]
	a.nums[i] = val
	i++
	for i < len(a.buf) {
		a.buf[i] += diff
		i = i + (i & -i)
	}
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */

func MakeRandArray() []int {
	maxLen, maxElement := 1000, 10
	//maxLen, maxElement := 10, 10
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
	k := MakeRandInt(j - i)
	nums[k] = MakeRandInt(100)
	obj.Update(k, nums[k])
	res = obj.SumRange(i, j)
	ans = bf(nums, i, j)
	if res != ans {
		panic(fmt.Errorf("After update, fail on %v, i %d, j %d, get %d, ans %d",
			nums, i, j, res, ans))
	}
}

func testOne() {
	nums := []int{7, 2, 7, 2, 0}
	obj := Constructor(nums)
	obj.Update(4, 6)
	obj.Update(0, 2)
	obj.Update(0, 9)
	nums = []int{9, 2, 7, 2, 6}
	obj.Update(3, 8)
	nums = []int{9, 2, 7, 8, 6}
	fmt.Println(obj.SumRange(0, 4))
}

func main() {
	testOne()
	return
	for i := 0; i < 10000; i++ {
		testSum()
	}
}
