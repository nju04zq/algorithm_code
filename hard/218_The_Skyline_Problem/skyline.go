package main

import "fmt"

type Range struct {
	x, y, h int
	next    *Range
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func makeRanges(a [][]int) *Range {
	dummy := new(Range)
	prev := dummy
	for _, b := range a {
		r := &Range{b[0], b[1], b[2], nil}
		prev.next = r
		prev = r
	}
	return dummy.next
}

func intersect(r1, r2 *Range) bool {
	if r1.x <= r2.x && r2.x < r1.y {
		return true
	} else if r2.x <= r1.x && r1.x < r2.y {
		return true
	} else {
		return false
	}
}

func contains(r1, r2 *Range) bool {
	if r1.x <= r2.x && r2.y <= r1.y {
		return true
	} else {
		return false
	}
}

func insert(r1, r2 *Range) {
	prev := r1
	for r := r1; r != nil; r = r.next {
		if r.x > r2.x {
			prev.next = r2
			r2.next = r
			return
		}
		prev = r
	}
	prev.next = r2
}

func dump(head *Range) {
	for r := head; r != nil; r = r.next {
		fmt.Printf("[%d, %d, %d], ", r.x, r.y, r.h)
	}
	fmt.Println()
}

func getSkyline(buildings [][]int) [][]int {
	head := makeRanges(buildings)
	var prev *Range
	for r := head; r != nil; {
		if prev == nil {
			prev = r
			r = r.next
			continue
		} else if !intersect(r, prev) {
			if r.x == prev.y && r.h == prev.h {
				prev.y = r.y
				prev.next = r.next
				r = prev.next
			} else {
				prev = r
				r = r.next
			}
			continue
		}
		//fmt.Printf("prev %d,%d,%d, r %d,%d,%d\n", prev.x, prev.y, prev.h, r.x, r.y, r.h)
		if r.h == prev.h {
			prev.y = max(prev.y, r.y)
			prev.next = r.next
			r = prev.next
		} else if r.h > prev.h {
			if contains(prev, r) {
				insert(r, &Range{r.y, prev.y, prev.h, nil})
				prev.y = r.x
				prev = r
				r = r.next
			} else {
				prev.y = r.x
				prev = r
				r = r.next
			}
		} else {
			if contains(prev, r) {
				prev.next = r.next
				r = r.next
			} else {
				r.x = prev.y
				prev.next = r.next
				r.next = nil
				insert(prev, r)
				if prev.next == r {
					prev = r
					r = r.next
				} else {
					r = prev.next
				}
			}
		}
		//dump(head)
	}
	res := make([][]int, 0)
	prev = nil
	for r := head; r != nil; r = r.next {
		//fmt.Printf("r %d,%d,%d\n", r.x, r.y, r.h)
		if r.x == r.y {
			continue
		}
		if prev != nil && prev.y != r.x {
			res = append(res, []int{prev.y, 0})
		}
		res = append(res, []int{r.x, r.h})
		prev = r
	}
	if prev != nil {
		res = append(res, []int{prev.y, 0})
	}
	return res
}

func main() {
	a := [][]int{
		[]int{2, 9, 10},
		[]int{3, 7, 15},
		[]int{5, 12, 12},
		[]int{15, 20, 10},
		[]int{19, 24, 8},
	}
	fmt.Println(getSkyline(a))
	a = [][]int{
		[]int{2, 9, 10},
		[]int{2, 7, 15},
		[]int{5, 12, 12},
		[]int{15, 20, 10},
		[]int{19, 24, 8},
	}
	fmt.Println(getSkyline(a))
	a = [][]int{
		[]int{60, 100, 41},
		[]int{60, 80, 91},
		[]int{70, 90, 72},
		[]int{85, 120, 59},
	}
	fmt.Println(getSkyline(a))
	a = [][]int{
		[]int{0, 2, 3},
		[]int{2, 5, 3},
	}
	fmt.Println(getSkyline(a))
}
