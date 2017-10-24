package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	var i int
	val := postorder[len(postorder)-1]
	for i, _ = range inorder {
		if inorder[i] == val {
			break
		}
	}
	root := &TreeNode{Val: val}
	root.Left = buildTree(inorder[:i], postorder[:i])
	root.Right = buildTree(inorder[i+1:], postorder[i:len(postorder)-1])
	return root
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

func testBuild(inorder, postorder []int) {
	root := buildTree(inorder, postorder)
	fmt.Printf("%v, %v, get: ", inorder, postorder)
	dumpTree(root)
}

func main() {
	// 1, 2, 3, 4, #, #, 5
	testBuild([]int{4, 2, 1, 3, 5}, []int{4, 2, 5, 3, 1})
}
