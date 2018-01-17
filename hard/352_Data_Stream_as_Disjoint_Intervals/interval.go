package main

import "fmt"

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */

type Interval struct {
	Start int
	End   int
}

func dumpIntervals(intervals []Interval) {
	for _, interval := range intervals {
		fmt.Printf("[%d, %d] ", interval.Start, interval.End)
	}
	fmt.Println()
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

type SummaryRanges struct {
	intervals []*Interval
}

/** Initialize your data structure here. */
func Constructor() SummaryRanges {
	r := new(SummaryRanges)
	r.intervals = make([]*Interval, 0)
	return *r
}

func (r *SummaryRanges) intersect(i0, i1 *Interval) bool {
	if i0.Start <= i1.Start && i1.Start <= (i0.End+1) {
		return true
	}
	if i1.Start <= i0.Start && i0.Start <= (i1.End+1) {
		return true
	}
	return false
}

func (r *SummaryRanges) merge(i0, i1 *Interval) *Interval {
	return &Interval{
		Start: min(i0.Start, i1.Start),
		End:   max(i0.End, i1.End),
	}
}

func (r *SummaryRanges) Addnum(val int) {
	newIntervals := make([]*Interval, 0)
	p := &Interval{val, val}
	for i := 0; i < len(r.intervals); i++ {
		interval := r.intervals[i]
		if p == nil {
			newIntervals = append(newIntervals, interval)
		} else if r.intersect(interval, p) {
			p = r.merge(interval, p)
		} else if p.Start < interval.Start {
			newIntervals = append(newIntervals, p)
			newIntervals = append(newIntervals, interval)
			p = nil
		} else {
			newIntervals = append(newIntervals, interval)
		}
	}
	if p != nil {
		newIntervals = append(newIntervals, p)
	}
	r.intervals = newIntervals
}

func (r *SummaryRanges) Getintervals() []Interval {
	res := make([]Interval, len(r.intervals))
	for i, interval := range r.intervals {
		res[i] = *interval
	}
	return res
}

/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Addnum(val);
 * param_2 := obj.Getintervals();
 */

func testInterval(r *SummaryRanges, nums []int) {
	fmt.Printf("nums %v\n", nums)
	for _, num := range nums {
		r.Addnum(num)
		dumpIntervals(r.Getintervals())
	}
}

func main() {
	sr := Constructor()
	r := &sr
	testInterval(r, []int{1, 3, 7, 2, 6})
}
