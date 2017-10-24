package main

import "fmt"
import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	cur := root
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		top := len(stack) - 1
		cur, stack = stack[top], stack[:top]
		res = append(res, cur.Val)
		cur = cur.Right
	}
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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

func testInorder(vals []string) {
	fmt.Printf("Tree %v\n", vals)
	root := makeTree(vals)
	res := inorderTraversal(root)
	fmt.Println(res)
}

func main() {
	testInorder([]string{"1", "2", "3", "#", "#", "4", "5", "#", "6", "7"})
}
