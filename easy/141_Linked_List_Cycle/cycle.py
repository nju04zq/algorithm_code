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
        """
        :type head: ListNode
        :rtype: bool
        """
        if head is None:
            return False
        p1, p2 = head, head
        while True:
            p1 = p1.next
            if p2.next is None:
                return False
            p2 = p2.next.next
            if p1 is None or p2 is None:
                return False
            elif p1 == p2:
                return True

a = [ListNode(1), ListNode(2), ListNode(3), ListNode(4)]
a[0].next = a[1]
a[1].next = a[2]
a[2].next = a[3]
a[3].next = a[0]
print Solution().hasCycle(a[0])
