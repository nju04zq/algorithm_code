class TreeLinkNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None
        self.next = None

# Definition for binary tree with next pointer.
# class TreeLinkNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
#         self.next = None

class Solution:
    # @param root, a tree link node
    # @return nothing
    def connect(self, root):
        last_level = root
        while last_level is not None:
            prev, this_level = None, None
            node = last_level
            while node is not None:
                if node.left is not None:
                    if prev is None:
                        this_level = node.left
                    else:
                        prev.next = node.left
                    prev = node.left
                if node.right is not None:
                    if prev is None:
                        this_level = node.right
                    else:
                        prev.next = node.right
                    prev = node.right
                node = node.next
            last_level = this_level
