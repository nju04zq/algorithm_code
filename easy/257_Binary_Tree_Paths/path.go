package main

import "fmt"
import "strconv"
import "strings"
import "bytes"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func makePath(path []int) string {
	buf := bytes.NewBuffer(nil)
	for i, val := range path {
		if i > 0 {
			buf.WriteString("->")
		}
		buf.WriteString(fmt.Sprintf("%d", val))
	}
	return buf.String()
}

func pathInternal(root *TreeNode, path []int, res []string) []string {
	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil {
		res = append(res, makePath(path))
	}
	if root.Left != nil {
		res = pathInternal(root.Left, path, res)
	}
	if root.Right != nil {
		res = pathInternal(root.Right, path, res)
	}
	path = path[:len(path)-1]
	return res
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	path := make([]int, 0)
	res := make([]string, 0)
	return pathInternal(root, path, res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func makeTree(s string) *TreeNode {
	s = strings.Replace(s, " ", "", -1)
	vals := strings.Split(s, ",")
	fmt.Println(vals)
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

func testPath(vals string) {
	root := makeTree(vals)
	paths := binaryTreePaths(root)
	fmt.Printf("%q, get %v\n", vals, paths)
}

func main() {
	testPath("1, 2, 3, 4, #, 6")
}
