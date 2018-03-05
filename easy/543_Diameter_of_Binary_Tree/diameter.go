package main

import "fmt"
import "strings"
import "strconv"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func dfs(root *TreeNode, d *int) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left, d)
	right := dfs(root.Right, d)
	*d = max(*d, left+right+1)
	return max(left, right) + 1
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	d := 0
	dfs(root, &d)
	return d - 1
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func makeTree(s string) *TreeNode {
	s = strings.Replace(s, " ", "", -1)
	vals := strings.Split(s, ",")
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

func testDiameter(s string) {
	root := makeTree(s)
	fmt.Printf("%q, get %d\n", s, diameterOfBinaryTree(root))
}

func main() {
	testDiameter("1, 2, 3, 4, 5")
	testDiameter("1, 2, 3, 4, 5, #, #, 6")
}
