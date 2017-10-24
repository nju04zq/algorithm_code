package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generate(nums []int) []*TreeNode {
	if len(nums) == 0 {
		return []*TreeNode{nil}
	} else if len(nums) == 1 {
		return []*TreeNode{&TreeNode{Val: nums[0]}}
	}
	res := make([]*TreeNode, 0)
	for i := 0; i < len(nums); i++ {
		res1 := generate(nums[:i])
		res2 := generate(nums[i+1:])
		for _, node1 := range res1 {
			for _, node2 := range res2 {
				root := &TreeNode{Val: nums[i], Left: node1, Right: node2}
				res = append(res, root)
			}
		}
	}
	return res
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}
	return generate(nums)
}

func dumpTree(root *TreeNode) {
	if root == nil {
		fmt.Println("Nil")
	}
	nodes := make([]*TreeNode, 0)
	nodes = append(nodes, root)
	for len(nodes) > 0 {
		cnt := len(nodes)
		allNil := true
		for i := 0; i < cnt; i++ {
			if nodes[i] == nil {
				fmt.Printf("# ")
				continue
			} else {
				fmt.Printf("%d ", nodes[i].Val)
			}
			nodes = append(nodes, nodes[i].Left)
			nodes = append(nodes, nodes[i].Right)
			if nodes[i].Left != nil || nodes[i].Right != nil {
				allNil = false
			}
		}
		if allNil {
			break
		}
		nodes = nodes[cnt:]
	}
	fmt.Println()
}

func testGenerate(n int) {
	res := generateTrees(n)
	for _, root := range res {
		dumpTree(root)
	}
}

func main() {
	testGenerate(3)
}
