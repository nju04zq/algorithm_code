package main

import "fmt"

func leastInterval(tasks []byte, n int) int {
	left := 0
	cnt := make([]int, 26)
	next := make([]int, 26)
	for i := 0; i < len(tasks); i++ {
		j := int(tasks[i] - 'A')
		cnt[j]++
		left++
	}
	round := 0
	for left > 0 {
		round++
		maxCnt, j := 0, 0
		for i := 0; i < len(cnt); i++ {
			if next[i] <= round && cnt[i] > maxCnt {
				j, maxCnt = i, cnt[i]
			}
		}
		if maxCnt == 0 {
			continue
		}
		cnt[j]--
		next[j] = round + n + 1
		left--
	}
	return round
}

func testLeast(tasks []byte, n int) {
	fmt.Printf("%v, n %d, get %d\n", tasks, n, leastInterval(tasks, n))
}

func main() {
	testLeast([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2)
	testLeast([]byte{'A', 'A', 'A', 'B', 'B', 'B', 'C'}, 2)
	testLeast([]byte{'A', 'A', 'B'}, 2)
}
