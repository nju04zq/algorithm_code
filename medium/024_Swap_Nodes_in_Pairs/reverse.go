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
func swapPairs(head *ListNode) *ListNode {
	var p1, p2, prev, next *ListNode
	for p1 = head; p1 != nil; {
		p2 = p1.Next
		if p2 == nil {
			if prev != nil {
				prev.Next = p1
			}
			break
		}
		next = p2.Next
		if prev == nil {
			head = p2
		} else {
			prev.Next = p2
		}
		p2.Next = p1
		p1.Next = nil
		prev = p1
		p1 = next
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
	fmt.Println("")
}

func testReverse(nums []int) {
	list := numArrayToList(nums)
	fmt.Printf("*****Reverse*****\n")
	dumpList(list)
	ans := swapPairs(list)
	dumpList(ans)
}

func main() {
	testReverse([]int{})
	testReverse([]int{1})
	testReverse([]int{1, 2})
	testReverse([]int{1, 2, 3, 4, 5})
}
