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
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}
	var dummy = &ListNode{Next: head}
	var prev, next, cur, first, tail *ListNode
	prev = dummy
	i := 1
	for cur = head; cur != nil; {
		next = cur.Next
		if i == m {
			first, tail = cur, prev
		} else if i > m {
			cur.Next = prev
			if i == n {
				first.Next, tail.Next = next, cur
				break
			}
		}
		prev, cur = cur, next
		i++
	}
	return dummy.Next
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

func testReverse(nums []int, m, n int) {
	head := numArrayToList(nums)
	fmt.Printf("Before m %d, n %d: ", m, n)
	dumpList(head)
	head = reverseBetween(head, m, n)
	fmt.Printf("After m %d, n %d:  ", m, n)
	dumpList(head)
}

func main() {
	testReverse([]int{1, 2, 3, 4, 5}, 1, 1)
	testReverse([]int{1, 2, 3, 4, 5}, 1, 2)
	testReverse([]int{1, 2, 3, 4, 5}, 1, 5)
	testReverse([]int{1, 2, 3, 4, 5}, 2, 3)
	testReverse([]int{1, 2, 3, 4, 5}, 2, 5)
	testReverse([]int{1, 2, 3, 4, 5}, 4, 5)
	testReverse([]int{1, 2, 3, 4, 5}, 5, 5)
}
