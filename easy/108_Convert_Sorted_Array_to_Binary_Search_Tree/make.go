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
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	i := len(nums) / 2
	root := &TreeNode{Val: nums[i]}
	root.Left = sortedArrayToBST(nums[:i])
	if i+1 < len(nums) {
		root.Right = sortedArrayToBST(nums[i+1:])
	}
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

func testMake(nums []int) {
	root := sortedArrayToBST(nums)
	fmt.Printf("%v, get: ", nums)
	dumpTree(root)
}

func main() {
	testMake([]int{1})
	testMake([]int{1, 2, 3})
	testMake([]int{1, 2, 3, 4, 5})
	testMake([]int{1, 2, 3, 4, 5, 6, 7})
}
