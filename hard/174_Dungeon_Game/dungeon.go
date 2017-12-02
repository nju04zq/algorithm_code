package main

import "fmt"
import "math"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func calculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 || len(dungeon[0]) == 0 {
		return 0
	}
	m, n := len(dungeon), len(dungeon[0])
	dp := make([]int, n)
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dp[j] = max(1, 1-dungeon[i][j])
				continue
			}
			if i < m-1 {
				dp[j] = max(1, dp[j]-dungeon[i][j])
			} else {
				dp[j] = math.MaxInt32
			}
			if j < n-1 {
				hp := max(1, dp[j+1]-dungeon[i][j])
				dp[j] = min(dp[j], hp)
			}
		}
		fmt.Println(dp)
	}
	return dp[0]
}

func testDungeon(dungeon [][]int) {
	fmt.Printf("Dungeon %v\nget %d\n", dungeon, calculateMinimumHP(dungeon))
}

func main() {
	dungeon := [][]int{
		[]int{-2, -3, 3},
		[]int{-5, -10, 1},
		[]int{10, 30, -5},
	}
	testDungeon(dungeon)
	dungeon = [][]int{
		[]int{-1, -102, 900},
		[]int{-1, -1, -1},
		[]int{-1, -1, -100},
	}
	testDungeon(dungeon)
}
