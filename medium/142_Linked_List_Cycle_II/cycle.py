class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def hasCycle(self, head):
        if head is None or head.next is None or head.next.next is None:
            return None
        p1, p2 = head.next, head.next.next
        while p1 != p2:
            p1, p2 = p1.next, p2.next
            if p1 is None or p2 is None:
                return None
            p2 = p2.next
            if p2 is None:
                return None
        return p1

    def getMeetpoint(self, h1, h2):
        p1, p2 = h1, h2
        while True:
            if p1 == p2:
                break
            if p1 is None:
                p1 = h2
            else:
                p1 = p1.next
            if p2 is None:
                p2 = h1
            else:
                p2 = p2.next
        return p1

    def detectCycle(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        p1 = self.hasCycle(head)
        if p1 is None:
            return None
        p2 = p1.next
        p1.next = None
        p = self.getMeetpoint(head, p2)
        p1.next = p2
        return p
        
a = [ListNode(0), ListNode(1), ListNode(2), ListNode(3)]
a[0].next = a[1]
a[1].next = a[2]
a[2].next = a[3]
print Solution().detectCycle(a[0])

a = [ListNode(0), ListNode(1), ListNode(2), ListNode(3)]
a[0].next = a[1]
a[1].next = a[2]
a[2].next = a[3]
a[3].next = a[0]
print Solution().detectCycle(a[0]).val

a = [ListNode(0), ListNode(1), ListNode(2), ListNode(3)]
a[0].next = a[1]
a[1].next = a[2]
a[2].next = a[3]
a[3].next = a[1]
print Solution().detectCycle(a[0]).val

a = [ListNode(0), ListNode(1), ListNode(2), ListNode(3)]
a[0].next = a[1]
a[1].next = a[2]
a[2].next = a[3]
a[3].next = a[2]
print Solution().detectCycle(a[0]).val

a = [ListNode(0), ListNode(1), ListNode(2), ListNode(3)]
a[0].next = a[1]
a[1].next = a[2]
a[2].next = a[3]
a[3].next = a[3]
print Solution().detectCycle(a[0]).val
