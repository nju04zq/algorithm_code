package main

import "fmt"
import "math/rand"
import "time"

func candy(ratings []int) int {
	total, cur, last, top := 0, 0, 0, 0
	stack := make([]int, 0)
	for i := 0; i < len(ratings); i++ {
		if i == 0 {
			cur = 1
		} else if ratings[i] > ratings[i-1] {
			cur = last + 1
		} else {
			cur = 1
			if len(stack) > 0 && cur == last {
				j := i - stack[len(stack)-1]
				total += (j - 1)
				if top <= j {
					total++
				}
				//fmt.Printf("%d/%d/%d\n", stack[len(stack)-1], j, total)
			}
		}
		total += cur
		last = cur
		//fmt.Printf("%d, %d\n", i, total)
		if i == len(ratings)-1 {
			break
		} else if ratings[i] <= ratings[i+1] {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			if len(stack) == 0 {
				stack = append(stack, i)
				top = cur
			}
		}
	}
	return total
}

func bf(ratings []int) int {
	n := len(ratings)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			res[i] = 1
			continue
		} else if ratings[i] > ratings[i-1] {
			res[i] = res[i-1] + 1
			continue
		} else {
			res[i] = 1
		}
		for j := i; j > 0; j-- {
			if ratings[j-1] > ratings[j] && res[j-1] <= res[j] {
				res[j-1] = res[j] + 1
			}
		}
	}
	for i := 0; i < n; i++ {
		if i > 0 && ratings[i] > ratings[i-1] && res[i] <= res[i-1] {
			panic(fmt.Sprintf("%v, res %v, %d fail", ratings, res, i))
		}
		if i < n-1 && ratings[i] > ratings[i+1] && res[i] <= res[i+1] {
			panic(fmt.Sprintf("%v, res %v, %d fail", ratings, res, i))
		}
	}
	total := 0
	for _, num := range res {
		total += num
	}
	return total
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

func testCandy() {
	a := MakeRandArray()
	//a = []int{2, 10, 15, 17, 14, 8, 6}
	//a = []int{13, 4, 1, 1, 0, 6, 11, 17, 17}
	//a = []int{9, 19, 3, 13, 11, 0}
	res := candy(a)
	ans := bf(a)
	if res != ans {
		panic(fmt.Sprintf("%v, get %d, expect %d\n", a, res, ans))
	}
}

func main() {
	for i := 0; i < 10000; i++ {
		testCandy()
	}
}
