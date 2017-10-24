package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	buf := make([]byte, 0)
	for ; l != nil; l = l.Next {
		buf = append(buf, byte('0'+l.Val))
	}
	return string(buf)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func getListLen(l *ListNode) int {
	len := 0
	for ; l != nil; l = l.Next {
		len++
	}
	return len
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}

func skipNodes(l *ListNode, count int) *ListNode {
	for i := 0; i < count && l != nil; i, l = i+1, l.Next {
	}
	return l
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1, s2 := getListLen(l1), getListLen(l2)
	smax := max(s1, s2)
	nodes := make([]ListNode, smax+1)
	diff := abs(s1 - s2)
	var l *ListNode
	if s1 > s2 {
		l = l1
		l1 = skipNodes(l1, diff)
	} else {
		l = l2
		l2 = skipNodes(l2, diff)
	}
	i := diff + 1
	for l1 != nil || l2 != nil {
		var val1, val2 int
		if l1 != nil {
			val1 = l1.Val
		}
		if l2 != nil {
			val2 = l2.Val
		}
		val := val1 + val2
		nodes[i].Val = val
		i++
		if i < len(nodes) {
			nodes[i-1].Next = &nodes[i]
		}
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	j := 1
	for i = 0; i < diff; i++ {
		nodes[j].Val = l.Val
		j++
		if j < len(nodes) {
			nodes[j-1].Next = &nodes[j]
		}
		l = l.Next
	}
	flag := 0
	for i = len(nodes) - 1; i >= 0; i-- {
		nodes[i].Val += flag
		if nodes[i].Val >= 10 {
			nodes[i].Val -= 10
			flag = 1
		} else {
			flag = 0
		}
	}
	if len(nodes) > 1 && nodes[0].Val == 0 {
		return &nodes[1]
	} else {
		nodes[0].Next = &nodes[1]
		return &nodes[0]
	}
}

func numArrayToList(num []int) *ListNode {
	var l, prev *ListNode
	for _, n := range num {
		node := &ListNode{Val: n}
		if prev == nil {
			l = node
		} else {
			prev.Next = node
		}
		prev = node
	}
	return l
}

func testAddTwoNumbers(num1 []int, num2 []int) {
	l1 := numArrayToList(num1)
	l2 := numArrayToList(num2)
	l3 := addTwoNumbers(l1, l2)
	fmt.Printf("%s + %s = %s\n", l1, l2, l3)
}

func main() {
	testAddTwoNumbers([]int{9, 9, 9}, []int{1})          //1, 0, 0, 0
	testAddTwoNumbers([]int{7, 2, 4, 3}, []int{5, 6, 4}) // 7, 8, 0, 7
	testAddTwoNumbers([]int{0}, []int{0})                // 0
}
