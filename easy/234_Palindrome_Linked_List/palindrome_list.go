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

func reverseTill(head *ListNode, end *ListNode) {
	var prev, cur *ListNode
	for cur = head; cur != end; {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
}

func getListLen(head *ListNode) int {
	var i int
	for i = 0; head != nil; head = head.Next {
		i++
	}
	return i
}

func getListMedian(head *ListNode, n int) (left, right, mid *ListNode) {
	leftIdx, rightIdx, midIdx := 0, 0, 0
	if n%2 == 0 {
		rightIdx = n / 2
		leftIdx = n/2 - 1
		midIdx = rightIdx
	} else {
		midIdx = n / 2
		rightIdx = midIdx + 1
		leftIdx = midIdx - 1
	}
	p := head
	for i := 0; i <= rightIdx; i++ {
		if i == midIdx {
			mid = p
		}
		if i == leftIdx {
			left = p
		} else if i == rightIdx {
			right = p
			break
		}
		p = p.Next
	}
	return left, right, mid
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	n := getListLen(head)
	if n <= 1 {
		return true
	}
	left, right, mid := getListMedian(head, n)
	reverseTill(head, left.Next)
	pLeft, pRight := left, right
	for pLeft != nil && pRight != nil {
		if pLeft.Val != pRight.Val {
			break
		}
		pLeft = pLeft.Next
		pRight = pRight.Next
	}
	reverseTill(left, nil)
	left.Next = mid
	if pLeft == nil && pRight == nil {
		return true
	} else {
		return false
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
