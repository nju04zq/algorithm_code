package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		cnt := len(nodes)
		for i := 0; i < cnt; i++ {
			nodes[i].Left, nodes[i].Right = nodes[i].Right, nodes[i].Left
			if nodes[i].Left != nil {
				nodes = append(nodes, nodes[i].Left)
			}
			if nodes[i].Right != nil {
				nodes = append(nodes, nodes[i].Right)
			}
		}
		nodes = nodes[cnt:]
	}
	return root
}

func main() {
	fmt.Println("vim-go")
}
