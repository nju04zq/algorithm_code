package main

import "fmt"

func buildGraph(n int, edges [][]int) map[int][]int {
	graph := make(map[int][]int)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, 0)
	}
	for i := 0; i < len(edges); i++ {
		v0, v1 := edges[i][0], edges[i][1]
		graph[v0] = append(graph[v0], v1)
		graph[v1] = append(graph[v1], v0)
	}
	return graph
}

func findMinHeightTrees(n int, edges [][]int) []int {
	graph := buildGraph(n, edges)
	leaves := []int{}
	degrees := make(map[int]int)
	visited := make(map[int]bool)
	for i, neighbors := range graph {
		degrees[i] = len(neighbors)
		if len(neighbors) <= 1 {
			leaves = append(leaves, i)
			visited[i] = true
		}
	}
	for n > 2 {
		cnt := len(leaves)
		n -= cnt
		for _, i := range leaves {
			for _, neighbor := range graph[i] {
				if _, ok := visited[neighbor]; ok {
					continue
				}
				degrees[neighbor]--
				if degrees[neighbor] <= 1 {
					leaves = append(leaves, neighbor)
					visited[neighbor] = true
				}
			}
		}
		leaves = leaves[cnt:]
	}
	return leaves
}

func testFind(n int, edges [][]int) {
	fmt.Printf("n %d, edges %v, get %d\n", n, edges, findMinHeightTrees(n, edges))
}

func main() {
	n := 4
	edges := [][]int{
		[]int{1, 0}, []int{1, 2}, []int{1, 3},
	}
	testFind(n, edges)
	n = 6
	edges = [][]int{
		[]int{0, 3}, []int{1, 3}, []int{2, 3}, []int{4, 3}, []int{5, 4},
	}
	testFind(n, edges)
	n = 3
	edges = [][]int{
		[]int{0, 1}, []int{0, 2},
	}
	testFind(n, edges)
}
