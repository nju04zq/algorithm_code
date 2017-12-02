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

func copyArray(a []int) []int {
	b := make([]int, len(a))
	for i, val := range a {
		b[i] = val
	}
	return b
}

func pathSumInternal(root *TreeNode, sum int, path []int, res [][]int) [][]int {
	if root == nil {
		return res
	}
	sum -= root.Val
	path = append(path, root.Val)
	if sum == 0 && (root.Left == nil && root.Right == nil) {
		res = append(res, copyArray(path))
	} else {
		if root.Left != nil {
			res = pathSumInternal(root.Left, sum, path, res)
		}
		if root.Right != nil {
			res = pathSumInternal(root.Right, sum, path, res)
		}
	}
	path = path[:len(path)-1]
	return res
}

func pathSum(root *TreeNode, sum int) [][]int {
	path := make([]int, 0)
	res := make([][]int, 0)
	return pathSumInternal(root, sum, path, res)
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

func testPathSum(vals []string, sum int) {
	res := pathSum(makeTree(vals), sum)
	fmt.Printf("%q, get %v\n", vals, res)
}

func main() {
	testPathSum([]string{"5", "4", "8", "11", "#", "13", "4", "7", "2", "#", "#", "5", "1"}, 22)
}
