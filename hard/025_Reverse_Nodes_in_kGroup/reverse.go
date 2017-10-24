package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(start, end *ListNode) (head, tail *ListNode) {
	var prev, cur, next *ListNode
	for cur = start; ; {
		if prev == nil {
			tail = cur
		}
		next = cur.Next
		cur.Next = prev
		if cur == end {
			break
		}
		prev = cur
		cur = next
	}
	head = cur
	return
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	var prev, cur, next, start, end *ListNode
	var i int
	for cur = head; cur != nil; {
		if start == nil {
			start = cur
		}
		i++
		if i < k {
			cur = cur.Next
			continue
		}
		next = cur.Next
		end = cur
		start, end = reverse(start, end)
		if prev == nil {
			head = start
		} else {
			prev.Next = start
		}
		prev = end
		i = 0
		start = nil
		cur = next
	}
	if prev != nil {
		prev.Next = start
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

func testReverse(nums []int, k int) {
	list := numArrayToList(nums)
	fmt.Printf("*****Reverse %d*****\n", k)
	dumpList(list)
	ans := reverseKGroup(list, k)
	dumpList(ans)
}

func main() {
	testReverse([]int{}, 0)
	testReverse([]int{}, 1)
	testReverse([]int{1, 2, 3, 4, 5}, 0)
	testReverse([]int{1, 2, 3, 4, 5}, 1)
	testReverse([]int{1, 2, 3, 4, 5}, 2)
	testReverse([]int{1, 2, 3, 4, 5}, 3)
	testReverse([]int{1, 2, 3, 4, 5}, 4)
	testReverse([]int{1, 2, 3, 4, 5}, 5)
	testReverse([]int{1, 2, 3, 4, 5}, 6)
}
