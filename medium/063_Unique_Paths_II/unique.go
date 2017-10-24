package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return 0
	}
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
				continue
			}
			if i == 0 && j == 0 {
				dp[j] = 1
				continue
			}
			if i > 0 && obstacleGrid[i-1][j] == 1 {
				dp[j] = 0
			}
			if j > 0 && obstacleGrid[i][j-1] != 1 {
				dp[j] += dp[j-1]
			}

		}
	}
	return dp[n-1]
}

func dumpArray(a [][]int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			fmt.Printf("%d ", a[i][j])
		}
		fmt.Println()
	}
}

func testUnique(a [][]int) {
	dumpArray(a)
	fmt.Printf("get %d\n", uniquePathsWithObstacles(a))
}

func main() {
	a := [][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 0},
		[]int{0, 0, 0},
	}
	testUnique(a)
	a = [][]int{
		[]int{1, 0},
	}
	testUnique(a)
}
