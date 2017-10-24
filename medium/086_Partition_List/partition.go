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
func partition(head *ListNode, x int) *ListNode {
	var low, lprev, high, hprev *ListNode
	for p := head; p != nil; {
		if p.Val < x {
			if lprev == nil {
				low = p
			} else {
				lprev.Next = p
			}
			lprev = p
		} else {
			if hprev == nil {
				high = p
			} else {
				hprev.Next = p
			}
			hprev = p
		}
		p, p.Next = p.Next, nil
	}
	if low != nil {
		head = low
		lprev.Next = high
	} else {
		head = high
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

func testPart(nums []int, x int) {
	head := numArrayToList(nums)
	fmt.Printf("Before partition: ")
	dumpList(head)
	partition(head, x)
	fmt.Printf("After partition %d: ", x)
	dumpList(head)
}

func main() {
	testPart([]int{1, 4, 3, 2, 5, 2}, 3)
}
