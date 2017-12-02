package main

import "fmt"
import "math"
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

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxPathSumInternal(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	} else if root.Left == nil && root.Right == nil {
		return root.Val, root.Val
	}
	var leftMax, leftRootMax, rightMax, rightRootMax int
	maxSum, rootMax := math.MinInt32, math.MinInt32
	if root.Left != nil {
		leftMax, leftRootMax = maxPathSumInternal(root.Left)
		maxSum = max(maxSum, leftMax)
	}
	if root.Right != nil {
		rightMax, rightRootMax = maxPathSumInternal(root.Right)
		maxSum = max(maxSum, rightMax)
	}
	if root.Left != nil && root.Right != nil {
		temp := root.Val
		if leftRootMax > 0 {
			temp += leftRootMax
		}
		if rightRootMax > 0 {
			temp += rightRootMax
		}
		maxSum = max(maxSum, temp)
		temp = max(leftRootMax, rightRootMax)
		if temp > 0 {
			rootMax = root.Val + temp
		} else {
			rootMax = root.Val
		}
	} else if root.Left != nil {
		temp := root.Val
		if leftRootMax > 0 {
			temp += leftRootMax
		}
		maxSum = max(maxSum, temp)
		rootMax = temp
	} else {
		temp := root.Val
		if rightRootMax > 0 {
			temp += rightRootMax
		}
		maxSum = max(maxSum, temp)
		rootMax = temp
	}
	return maxSum, rootMax
}

func maxPathSum(root *TreeNode) int {
	maxSum, _ := maxPathSumInternal(root)
	return maxSum
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

func testMaxSum(vals []string) {
	fmt.Printf("%q, get %d\n", vals, maxPathSum(makeTree(vals)))
}

func main() {
	testMaxSum([]string{"1", "2", "3"})
	testMaxSum([]string{"1", "1", "4", "2", "3", "2", "3"})
	testMaxSum([]string{"10", "1", "4", "2", "3", "2", "3"})
}
