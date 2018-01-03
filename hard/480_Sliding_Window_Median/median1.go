package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func split(nums []int, low, high int) int {
	pilot := nums[low]
	j := low + 1
	for i := low + 1; i <= high; i++ {
		if nums[i] < pilot {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	nums[low], nums[j-1] = nums[j-1], nums[low]
	return j - 1 - low
}

func getNum(nums []int, low, high int, k int) int {
	//fmt.Printf("nums %v, low %d, high %d, k %d\n", nums, low, high, k)
	j := split(nums, low, high)
	if j == k {
		return nums[low+j]
	} else if j < k {
		return getNum(nums, low+j+1, high, k-j-1)
	} else {
		return getNum(nums, low, low+j-1, k)
	}
}

func median(nums []int) float64 {
	n := len(nums)
	if n%2 != 0 {
		a := getNum(nums, 0, n-1, n/2)
		return float64(a)
	} else {
		a := getNum(nums, 0, n-1, n/2-1)
		b := getNum(nums, 0, n-1, n/2)
		return (float64(a) + float64(b)) / 2
	}
}

func medianSlidingWindow(nums []int, k int) []float64 {
	n := len(nums)
	res := make([]float64, 0)
	temp := make([]int, k)
	for i := 0; i <= n-k; i++ {
		for j := 0; j < k; j++ {
			temp[j] = nums[i+j]
		}
		res = append(res, median(temp))
	}
	return res
}

func bfMedian(nums []int) float64 {
	nums1 := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		nums1[i] = nums[i]
	}
	sort.Ints(nums1)
	n := len(nums1)
	if n%2 != 0 {
		return float64(nums1[n/2])
	} else {
		return (float64(nums1[n/2-1]) + float64(nums1[n/2])) / 2
	}
}

func bf(nums []int, k int) []float64 {
	n := len(nums)
	res := make([]float64, 0)
	for i := 0; i <= n-k; i++ {
		res = append(res, bfMedian(nums[i:i+k]))
	}
	return res
}

func MakeRandInt(i int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int()%i + 1
}

func MakeRandArray() []int {
	//maxLen, maxElement := 20, 20
	maxLen, maxElement := 10, 20
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len := r.Int()%maxLen + 4
	a := make([]int, len)
	for i := 0; i < len; i++ {
		a[i] = r.Int() % maxElement
	}
	return a
}

func testMedian() {
	nums := MakeRandArray()
	k := MakeRandInt(len(nums) / 2)
	res := medianSlidingWindow(nums, k)
	ans := bf(nums, k)
	if len(res) != len(ans) {
		panic(fmt.Errorf("%v, res len %d, ans len %d",
			nums, len(res), len(ans)))
	}
	for i := 0; i < len(res); i++ {
		if res[i] != ans[i] {
			panic(fmt.Errorf("%v, k %d, get %v, expect %v", nums, k, res, ans))
		}
	}
}

func testGetMedian() {
	a := MakeRandArray()
	k := MakeRandInt(len(a) - 1)
	b := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		b[i] = a[i]
	}
	sort.Ints(b)
	res := getNum(a, 0, len(a)-1, k)
	ans := b[k]
	if res != ans {
		panic(fmt.Errorf("%v, k %d, get %d, expect %d", a, k, res, ans))
	}
}

func main() {
	for i := 0; i < 10000; i++ {
		fmt.Printf("\r%d", i)
		testMedian()
	}
	fmt.Println()
}
