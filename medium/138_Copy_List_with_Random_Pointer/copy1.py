# Definition for singly-linked list with a random pointer.
class RandomListNode(object):
    def __init__(self, x):
        self.label = x
        self.next = None
        self.random = None

class Solution(object):
    def copyRandomList(self, head):
        """
        :type head: RandomListNode
        :rtype: RandomListNode
        """
        p = head
        while p is not None:
            next = p.next
            node = RandomListNode(p.label)
            p.next = node
            node.next = next
            p = next
        p = head
        while p is not None:
            if p.random is not None:
                p.next.random = p.random.next
            p = p.next.next
        p = head
        head1 = None
        while p is not None:
            if head1 is None:
                head1 = p.next
            p1 = p.next
            p.next = p.next.next
            if p.next is None:
                p1.next = None
            else:
                p1.next = p.next.next
            p = p.next
        return head1

Solution().copyRandomList(None)
