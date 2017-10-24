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
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var node, left *TreeNode
	stack := []*TreeNode{root}
	visited := []bool{false}
	left = root.Left
	res := make([]int, 0)
	for len(stack) > 0 {
		if left != nil {
			stack = append(stack, left)
			visited = append(visited, false)
			left = left.Left
			continue
		}
		top := len(stack) - 1
		node = stack[top]
		if visited[top] == true || node.Right == nil {
			res = append(res, node.Val)
			stack = stack[:top]
			visited = visited[:top]
		} else {
			visited[top] = true
			stack = append(stack, node.Right)
			visited = append(visited, false)
			left = node.Right.Left
		}
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

func testPostorder(vals []string) {
	fmt.Printf("Tree: %v\n", vals)
	root := makeTree(vals)
	res := postorderTraversal(root)
	fmt.Println(res)
}

func main() {
	testPostorder([]string{"1", "2", "3", "#", "#", "4", "5", "#", "6", "7"})
}
