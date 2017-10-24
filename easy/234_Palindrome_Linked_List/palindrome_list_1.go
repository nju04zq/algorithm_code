package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	s := ""
	for p := l; p != nil; p = p.Next {
		s += fmt.Sprintf("%d ", p.Val)
	}
	return s
}

func getMedian(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverse(head *ListNode) *ListNode {
	var prev, cur *ListNode
	for cur = head; cur != nil; {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	tail := prev
	return tail
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	mid := getMedian(head)
	tail := reverse(mid.Next)
	pLeft, pRight := head, tail
	for pRight != nil {
		if pLeft.Val != pRight.Val {
			break
		}
		pLeft, pRight = pLeft.Next, pRight.Next
	}
	reverse(tail)
	if pRight != nil {
		return false
	} else {
		return true
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

func testIsPalindrome(a []int, expect bool) {
	l1 := numArrayToList(a)
	result := isPalindrome(l1)
	if result != expect {
		fmt.Printf("IsPalindrome %v, get %t, expect %t\n", a, result, expect)
	}
	i := 0
	for p := l1; p != nil; p = p.Next {
		if p.Val != a[i] {
			fmt.Printf("List changed, get %v, expect %v\n", l1, a)
			return
		}
		i++
	}
	if i != len(a) {
		fmt.Printf("List changed, get %v, expect %v\n", l1, a)
	}
}

func main() {
	testIsPalindrome([]int{}, true)
	testIsPalindrome([]int{0}, true)
	testIsPalindrome([]int{1}, true)
	testIsPalindrome([]int{1, 1}, true)
	testIsPalindrome([]int{1, 2}, false)
	testIsPalindrome([]int{1, 2, 1}, true)
	testIsPalindrome([]int{1, 2, 2}, false)
	testIsPalindrome([]int{1, 2, 2, 1}, true)
	testIsPalindrome([]int{1, 2, 3, 1}, false)
	testIsPalindrome([]int{1, 2, 2, 3}, false)
	testIsPalindrome([]int{1, 2, 3, 2, 1}, true)
	testIsPalindrome([]int{1, 2, 3, 1, 1}, false)
	testIsPalindrome([]int{1, 2, 3, 3, 2, 1}, true)
	testIsPalindrome([]int{1, 2, 3, 4, 2, 1}, false)
	testIsPalindrome([]int{1, 2, 3, 3, 4, 1}, false)
}
