# Definition for singly-linked list with a random pointer.
class RandomListNode(object):
    def __init__(self, x):
        self.label = x
        self.next = None
        self.random = None

class Solution(object):
    def get_nodes(self, node, nodes):
        if node is None:
            return None
        if node in nodes:
            return nodes[node]
        else:
            node1 = RandomListNode(node.label)
            nodes[node] = node1
            return node1

    def copyRandomList(self, head):
        """
        :type head: RandomListNode
        :rtype: RandomListNode
        """
        nodes = {}
        p = head
        head1 = None
        while p is not None:
            p1 = self.get_nodes(p, nodes)
            p1.next = self.get_nodes(p.next, nodes)
            p1.random = self.get_nodes(p.random, nodes)
            if head1 is None:
                head1 = p1
            p = p.next
        return head1

Solution().copyRandomList(None)
