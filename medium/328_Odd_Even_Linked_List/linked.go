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
func oddEvenList(head *ListNode) *ListNode {
	var ohead, ehead *ListNode
	opprev, epprev := &ohead, &ehead
	p, i := head, 1
	for p != nil {
		if i%2 == 1 {
			*opprev = p
			opprev = &p.Next
		} else {
			*epprev = p
			epprev = &p.Next
		}
		next := p.Next
		p.Next = nil
		p = next
		i++
	}
	*opprev = ehead
	return ohead
}

func numArrayToList(nums []int) *ListNode {
	var l, prev *ListNode
	for _, n := range nums {
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

func testLink(nums []int) {
	head := numArrayToList(nums)
	dumpList(head)
	head = oddEvenList(head)
	dumpList(head)
}

func main() {
	testLink([]int{1, 2, 3, 4, 5})
}
