package main

import "fmt"
import "sort"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func leastInterval(tasks []byte, n int) int {
	cnt := make([]int, 26)
	for i := 0; i < len(tasks); i++ {
		j := int(tasks[i] - 'A')
		cnt[j]++
	}
	sort.Ints(cnt)
	j, m := 1, len(cnt)
	for i := m - 2; i >= 0; i-- {
		if cnt[i] != cnt[m-1] {
			break
		}
		j++
	}
	return max(len(tasks), ((cnt[m-1]-1)*(n+1) + j))
}

func testLeast(tasks []byte, n int) {
	fmt.Printf("%v, n %d, get %d\n", tasks, n, leastInterval(tasks, n))
}

func main() {
	testLeast([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2)
	testLeast([]byte{'A', 'A', 'A', 'B', 'B', 'B', 'C'}, 2)
	testLeast([]byte{'A', 'A', 'B'}, 2)
}
