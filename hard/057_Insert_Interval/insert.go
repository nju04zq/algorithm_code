package main

import "fmt"

func dumpIntervals(a []Interval) {
	for _, x := range a {
		fmt.Printf("[%d, %d] ", x.Start, x.End)
	}
	fmt.Println()
}

type Interval struct {
	Start, End int
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
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

func merge(x, y *Interval) *Interval {
	x.Start = min(x.Start, y.Start)
	x.End = max(x.End, y.End)
	return x
}

/**
 * Definition for an interval.
 * type Interval struct {
 *     Start int
 *     End   int
 * }
 */
func insert(a []Interval, in Interval) []Interval {
	b := make([]Interval, 0)
	p := &in
	for i, _ := range a {
		cur := &a[i]
		if p == nil {
			b = append(b, *cur)
		} else if !intersect(cur, p) {
			if p.Start < cur.Start {
				b = append(b, *p)
				p = nil
			}
			b = append(b, *cur)
		} else {
			p = merge(p, cur)
		}
	}
	if p != nil {
		b = append(b, *p)
	}
	return b
}

func testInsert(a []Interval, in Interval) {
	fmt.Printf("Before insert [%d, %d]:\n", in.Start, in.End)
	dumpIntervals(a)
	res := insert(a, in)
	fmt.Printf("After insert:\n")
	dumpIntervals(res)
}

func main() {
	testInsert([]Interval{}, Interval{1, 2})
	testInsert([]Interval{Interval{2, 3}, Interval{4, 5}}, Interval{0, 1})
	testInsert([]Interval{Interval{2, 3}, Interval{4, 5}}, Interval{1, 2})
	testInsert([]Interval{Interval{2, 3}, Interval{4, 5}}, Interval{1, 6})
	testInsert([]Interval{Interval{2, 3}, Interval{4, 5}}, Interval{7, 8})
	testInsert([]Interval{Interval{2, 3}, Interval{7, 8}}, Interval{4, 5})
	testInsert([]Interval{Interval{2, 3}, Interval{7, 8}}, Interval{3, 7})
}
