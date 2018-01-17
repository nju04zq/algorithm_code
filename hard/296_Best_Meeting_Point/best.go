package main

import "fmt"
import "sort"

// A group of two or more people wants to meet and minimize the total travel distance. You are given a 2D grid of values 0 or 1, where each 1 marks the home of someone in the group. The distance is calculated using Manhattan Distance, where distance(p1, p2) = |p2.x - p1.x| + |p2.y - p1.y|.
//
// For example, given three people living at (0,0), (0,4), and (2,2):
//
// 1 - 0 - 0 - 0 - 1
// |   |   |   |   |
// 0 - 0 - 0 - 0 - 0
// |   |   |   |   |
// 0 - 0 - 1 - 0 - 0
// The point (0,2) is an ideal meeting point, as the total travel distance of 2+2+2=6 is minimal. So return 6.
//
// Hint:
//
// Try to solve it in one dimension first. How can this solution apply to the two dimension case?

func BestMeetingPoint(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}
	m, n := len(grid), len(grid[0])
	x, y := make([]int, 0), make([]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				x = append(x, j)
				y = append(y, i)
			}
		}
	}
	sort.Ints(x)
	total := 0
	i, j := 0, len(x)-1
	for i < j {
		total += (x[j] - x[i])
		total += (y[j] - y[i])
		i++
		j--
	}
	return total
}

func main() {
	grid := [][]int{
		[]int{1, 0, 0, 0, 1},
		[]int{0, 0, 0, 0, 0},
		[]int{0, 0, 1, 0, 0},
	}
	fmt.Println(BestMeetingPoint(grid))
}
