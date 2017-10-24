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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var flag = 0
	var l3, prev *ListNode
	for l1 != nil || l2 != nil {
		var val1, val2 int
		if l1 != nil {
			val1 = l1.Val
		}
		if l2 != nil {
			val2 = l2.Val
		}
		val := val1 + val2 + flag
		if val >= 10 {
			val -= 10
			flag = 1
		} else {
			flag = 0
		}
		node := &ListNode{Val: val}
		if prev == nil {
			l3 = node
		} else {
			prev.Next = node
		}
		prev = node
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if flag > 0 {
		prev.Next = &ListNode{Val: flag}
	}
	return l3
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
	testAddTwoNumbers([]int{0}, []int{0})             // 0
	testAddTwoNumbers([]int{9, 9, 9}, []int{1})       // 0, 0, 0, 1
	testAddTwoNumbers([]int{2, 4, 3}, []int{5, 6, 4}) //7, 0, 8
}
