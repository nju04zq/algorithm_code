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
func deleteDuplicates(head *ListNode) *ListNode {
	var prev, cur0, cur *ListNode
	var cnt int
	for cur = head; cur != nil; cur = cur.Next {
		if cur0 == nil {
			cur0 = cur
		} else if cur.Val == cur0.Val {
			cnt++
		} else {
			if cnt == 0 {
				if prev == nil {
					head = cur0
				} else {
					prev.Next = cur0
				}
				prev = cur0
			}
			cnt = 0
			cur0 = cur
		}
	}
	if cur0 != nil && cnt == 0 {
		if prev == nil {
			head = cur0
		} else {
			prev.Next = cur0
		}
		cur0.Next = nil
	} else if prev != nil {
		prev.Next = nil
	} else {
		head = nil
	}
	return head
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

func testDelete(nums []int) {
	l := numArrayToList(nums)
	fmt.Printf("Before: ")
	dumpList(l)
	l = deleteDuplicates(l)
	fmt.Printf("After: ")
	dumpList(l)
}

func main() {
	testDelete([]int{})
	testDelete([]int{1})
	testDelete([]int{1, 1})
	testDelete([]int{1, 1, 2, 3, 3})
}
