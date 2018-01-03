package main

import "fmt"

func paint(n, k int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return k
	}
	same := k
	diff := k * (k - 1)
	for i := 2; i < n; i++ {
		temp := diff
		diff = (same + diff) * (k - 1)
		same = temp
	}
	return same + diff
}

func main() {
	fmt.Println("vim-go")
}
