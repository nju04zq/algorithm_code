package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, lp1, lp2, cur, prev *ListNode
	lp1, lp2 = l1, l2
	for lp1 != nil || lp2 != nil {
		if lp1 == nil {
			cur = &ListNode{Val: lp2.Val}
			lp2 = lp2.Next
		} else if lp2 == nil {
			cur = &ListNode{Val: lp1.Val}
			lp1 = lp1.Next
		} else if lp1.Val < lp2.Val {
			cur = &ListNode{Val: lp1.Val}
			lp1 = lp1.Next
		} else {
			cur = &ListNode{Val: lp2.Val}
			lp2 = lp2.Next
		}
		if prev == nil {
			head = cur
		} else {
			prev.Next = cur
		}
		prev = cur
	}
	return head
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

func dumpList(head *ListNode) {
	if head == nil {
		fmt.Println("Nil")
		return
	}
	for p := head; p != nil; p = p.Next {
		if p == head {
			fmt.Printf("%d", p.Val)
		} else {
			fmt.Printf(" -> %d", p.Val)
		}
	}
	fmt.Println()
}

func testMerge(a, b []int) {
	l1 := numArrayToList(a)
	l2 := numArrayToList(b)
	l3 := mergeTwoLists(l1, l2)
	fmt.Println("l1, l2, l3")
	dumpList(l1)
	dumpList(l2)
	dumpList(l3)
}

func main() {
	testMerge([]int{}, []int{})
	testMerge([]int{}, []int{1, 2, 3})
	testMerge([]int{1, 2, 3}, []int{})
	testMerge([]int{2, 4, 6}, []int{1, 3, 5})
}
