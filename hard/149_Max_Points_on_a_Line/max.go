package main

import "fmt"

type Point struct {
	X int
	Y int
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxThrough(points []Point, i int, dx, dy []int, d []float64) int {
	cnt0, same := 0, 0
	k := 0
	for j := 0; j < i; j++ {
		dx[j] = points[j].X - points[i].X
		dy[j] = points[j].Y - points[i].Y
		if dx[j] == 0 && dy[j] == 0 {
			same++
		} else if dy[j] == 0 {
			cnt0++
		} else {
			d[k] = float64(dx[j]) / float64(dy[j])
			k++
		}
	}
	cnt1 := 0
	tbl := make(map[float64]int)
	for j := 0; j < k; j++ {
		if _, ok := tbl[d[j]]; !ok {
			tbl[d[j]] = 1
		} else {
			tbl[d[j]]++
		}
		cnt1 = max(cnt1, tbl[d[j]])
	}
	return max(cnt0+same+1, cnt1+same+1)
}

/**
 * Definition for a point.
 * type Point struct {
 *     X int
 *     Y int
 * }
 */
func maxPoints(points []Point) int {
	n := len(points)
	if n < 2 {
		return n
	} else if n == 2 {
		return 2
	}
	dx := make([]int, n)
	dy := make([]int, n)
	d := make([]float64, n)
	dp := make([]int, n)
	dp[1] = 2
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], maxThrough(points, i, dx, dy, d))
	}
	return dp[n-1]
}

func testMaxPoints(points []Point) {
	fmt.Println("======")
	for _, point := range points {
		fmt.Printf("(%d, %d) ", point.X, point.Y)
	}
	fmt.Println()
	fmt.Printf("get %d\n", maxPoints(points))
}

func main() {
	points := []Point{
		Point{1, 2},
		Point{1, 1},
		Point{2, 4},
		Point{2, 2},
		Point{3, 6},
	}
	testMaxPoints(points)
	points = []Point{
		Point{0, 0},
		Point{1, 1},
		Point{0, 0},
	}
	testMaxPoints(points)
	points = []Point{
		Point{1, 3},
		Point{2, 3},
		Point{3, 3},
	}
	testMaxPoints(points)
}
