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

func maxPathSumInternal(root *TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}
	left := max(0, maxPathSumInternal(root.Left, maxSum))
	right := max(0, maxPathSumInternal(root.Right, maxSum))
	*maxSum = max(*maxSum, left+right+root.Val)
	return max(left, right) + root.Val
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	maxPathSumInternal(root, &maxSum)
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
