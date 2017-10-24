# Comments
#   x0       x1
# 0 --> 1 --> 2 --> 3 --> 1
# Head              MeetAt
# x0 is the distance from head to cycle entry
# x1 is the distance from cycle entry to meet point
# n  is the length for the cycle
# we have 2(x0+x1) = x0 + x1 + kn
# note, while slow pointer enter the cycle, it can't finish one loop
# consider the worest case, fast pointer is just in front of slow pointer
# we can get x0+x1=kn
# that means, pointer travels x1 from cycle entry, then travel x0,
# can reach the cycle entry, so one from head, one from meet point, both
# travel x0 distance, they certainly meet at the cycle entry

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
    def detectCycle(self, head):
        if head is None or head.next is None or head.next.next is None:
            return None
        slow, fast = head.next, head.next.next
        while slow != fast:
            slow, fast = slow.next, fast.next
            if slow is None or fast is None:
                return None
            fast = fast.next
            if fast is None:
                return None
        fast = head
        while slow != fast:
            slow = slow.next
            fast = fast.next
        return slow

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
