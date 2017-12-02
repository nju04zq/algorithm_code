package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

type Range struct {
	start, end int
}

func (r *Range) length() int {
	return r.end - r.start
}

func (r *Range) intersect(r1 *Range) bool {
	if r.start <= r1.start && r1.start < r.end {
		return true
	} else if r1.start <= r.start && r.start < r1.end {
		return true
	} else {
		return false
	}
}

func (r *Range) overlap(r1 *Range) int {
	if r1.start < r.start {
		r, r1 = r1, r
	}
	return min(r.end, r1.end) - r1.start
}

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	rax := &Range{A, C}
	ray := &Range{B, D}
	rbx := &Range{E, G}
	rby := &Range{F, H}
	total := rax.length()*ray.length() + rbx.length()*rby.length()
	if rax.intersect(rbx) && ray.intersect(rby) {
		total -= rax.overlap(rbx) * ray.overlap(rby)
	}
	return total
}

func testCompute(A, B, C, D, E, F, G, H int) {
	fmt.Printf("(%d, %d), (%d, %d), (%d, %d), (%d, %d), get %d\n",
		A, B, C, D, E, F, G, H, computeArea(A, B, C, D, E, F, G, H))
}

func main() {
	testCompute(-2, -2, 2, 2, -2, -2, 2, 2)
	testCompute(0, 0, 0, 0, -1, -1, 1, 1)
}
