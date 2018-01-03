package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	tbl := make(map[int]int)
	for _, num := range nums {
		if _, ok := tbl[num]; ok {
			tbl[num]++
		} else {
			tbl[num] = 1
		}
	}
	freqs := make([][]int, len(nums))
	for num, cnt := range tbl {
		if freqs[cnt-1] == nil {
			freqs[cnt-1] = make([]int, 0)
		}
		freqs[cnt-1] = append(freqs[cnt-1], num)
	}
	res := make([]int, 0)
	total := 0
	for i := len(freqs) - 1; i >= 0; i-- {
		if freqs[i] == nil {
			continue
		}
		for j := 0; j < len(freqs[i]); j++ {
			res = append(res, freqs[i][j])
			total++
			if total >= k {
				break
			}
		}
		if total >= k {
			break
		}
	}
	return res
}

func testTopK(nums []int, k int) {
	fmt.Printf("%v, k %d, get %v\n", nums, k, topKFrequent(nums, k))
}

func main() {
	testTopK([]int{1, 1, 1, 2, 2, 3}, 1)
	testTopK([]int{1, 1, 1, 2, 2, 3}, 2)
	testTopK([]int{1, 1, 1, 2, 2, 3}, 3)
}
