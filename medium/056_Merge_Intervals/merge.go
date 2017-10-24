package main

import "fmt"
import "sort"

type Interval struct {
	Start, End int
}

func dumpIntervals(a []Interval) {
	for _, x := range a {
		fmt.Printf("[%d, %d] ", x.Start, x.End)
	}
	fmt.Println()
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func intersect(x, y *Interval) bool {
	if x.Start <= y.Start && y.Start <= x.End {
		return true
	} else if y.Start <= x.Start && x.Start <= y.End {
		return true
	} else {
		return false
	}
}

/*
 * Definition for an interval.
 * type Interval struct {
 *     Start int
 *     End   int
 * }
 */
func merge(a []Interval) []Interval {
	sort.Slice(a, func(i, j int) bool {
		if a[i].Start < a[j].Start {
			return true
		} else {
			return false
		}
	})
	b := make([]Interval, 0)
	var prev *Interval
	for i := 0; i < len(a); i++ {
		cur := &a[i]
		if prev == nil {
			prev = cur
			continue
		}
		if !intersect(prev, cur) {
			b = append(b, *prev)
			prev = cur
		} else {
			prev.End = max(prev.End, cur.End)
		}
	}
	if prev != nil {
		b = append(b, *prev)
	}
	return b
}

func testMerge(a []Interval) {
	fmt.Println("Before merge:")
	dumpIntervals(a)
	res := merge(a)
	fmt.Println("After merge:")
	dumpIntervals(res)
}

func main() {
	testMerge([]Interval{})
	testMerge([]Interval{Interval{2, 4}, Interval{1, 3}, Interval{3, 5}})
	testMerge([]Interval{Interval{1, 2}, Interval{3, 4}, Interval{5, 6}})
	testMerge([]Interval{Interval{2, 3}, Interval{4, 5}, Interval{1, 6}})
}
