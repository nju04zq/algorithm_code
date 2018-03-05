// inorder not work
// [3,4,5,1,2,null,null,0], [4,1,2]
// preorder & postorder can work

package main

import "fmt"
import "bytes"
import "strings"
import "strconv"

func preorderDfs(s *TreeNode, buf *bytes.Buffer) {
	if s == nil {
		buf.WriteString(",#")
		return
	}
	buf.WriteString(fmt.Sprintf(",%d", s.Val))
	preorderDfs(s.Left, buf)
	preorderDfs(s.Right, buf)
}

func preorder(s *TreeNode) string {
	buf := bytes.NewBuffer(nil)
	preorderDfs(s, buf)
	return buf.String()
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(s *TreeNode, t *TreeNode) bool {
	s0 := preorder(s)
	s1 := preorder(t)
	return strings.Contains(s0, s1)
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

func testSubTree(s0, s1 string) {
	s, t := makeTree(s0), makeTree(s1)
	fmt.Printf("s %q, t %q, get %t\n", s0, s1, isSubtree(s, t))
}

func main() {
	testSubTree("3, 4, 5, 1, 2", "4, 1, 2")
	testSubTree("3, 4, 5, 1, 2, #, #, #, #, 0", "4, 1, 2")
	testSubTree("1", "1")
	testSubTree("12", "2")
}
