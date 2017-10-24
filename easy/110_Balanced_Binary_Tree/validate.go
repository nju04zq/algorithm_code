package main

import "fmt"
import "strconv"

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

func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func treeDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := treeDepth(root.Left)
	if leftDepth == -1 {
		return -1
	}
	rightDepth := treeDepth(root.Right)
	if rightDepth == -1 {
		return -1
	}
	if abs(rightDepth-leftDepth) > 1 {
		return -1
	}
	return max(leftDepth, rightDepth) + 1
}

func isBalanced(root *TreeNode) bool {
	if treeDepth(root) != -1 {
		return true
	} else {
		return false
	}
}

func makeTree(vals []string) *TreeNode {
	makeNode := func(val string) *TreeNode {
		if val == "#" {
			return nil
		}
		node := new(TreeNode)
		if i, err := strconv.ParseInt(val, 10, 32); err != nil {
			panic(fmt.Errorf("Fail to parse %q, %v", val, err))
		} else {
			node.Val = int(i)
		}
		return node
	}
	var root, node *TreeNode
	lastLevel := make([]*TreeNode, 0)
	for i := 0; i < len(vals); {
		if len(lastLevel) == 0 {
			if root = makeNode(vals[i]); root == nil {
				return nil
			}
			lastLevel = append(lastLevel, root)
			i++
			continue
		}
		cnt := len(lastLevel)
		for j := 0; j < cnt && i < len(vals); j++ {
			node = makeNode(vals[i])
			lastLevel[j].Left = node
			if node != nil {
				lastLevel = append(lastLevel, node)
			}
			i++
			if i >= len(vals) {
				break
			}
			node = makeNode(vals[i])
			lastLevel[j].Right = node
			if node != nil {
				lastLevel = append(lastLevel, node)
			}
			i++
		}
		lastLevel = lastLevel[cnt:]
	}
	return root
}

func testIsBalanced(vals []string) {
	fmt.Printf("%q, %t\n", vals, isBalanced(makeTree(vals)))
}

func main() {
	testIsBalanced([]string{"1", "2", "3", "4"})
	testIsBalanced([]string{"1", "2", "#", "3", "4", "5", "6"})
}
