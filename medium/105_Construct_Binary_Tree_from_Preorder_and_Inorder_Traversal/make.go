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
func buildTree(preorder []int, inorder []int) *TreeNode {
	var i int
	if len(preorder) == 0 {
		return nil
	}
	for i, _ = range inorder {
		if inorder[i] == preorder[0] {
			break
		}
	}
	leftCnt := i
	root := &TreeNode{Val: preorder[0]}
	root.Left = buildTree(preorder[1:leftCnt+1], inorder[:i])
	root.Right = buildTree(preorder[leftCnt+1:], inorder[i+1:])
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

func testBuild(preorder, inorder []int) {
	root := buildTree(preorder, inorder)
	fmt.Printf("%v, %v, get: ", preorder, inorder)
	dumpTree(root)
}

func main() {
	// 1, 2, 3, 4, #, #, 5
	testBuild([]int{1, 2, 4, 3, 5}, []int{4, 2, 1, 3, 5})
}
