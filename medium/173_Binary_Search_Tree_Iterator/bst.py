import collections

class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

# Definition for a  binary tree node
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class BSTIterator(object):
    def __init__(self, root):
        """
        :type root: TreeNode
        """
        self.stack = []
        node = root
        while node is not None:
            self.stack.append(node)
            node = node.left

    def hasNext(self):
        """
        :rtype: bool
        """
        return len(self.stack) > 0
        

    def next(self):
        """
        :rtype: int
        """
        if len(self.stack) == 0:
            return None
        next_node = self.stack.pop()
        node = next_node.right
        while node is not None:
            self.stack.append(node)
            node = node.left
        return next_node.val

def create_node(s):
    if s == "#":
        return None
    node = TreeNode(int(s))
    return node
        
def make_tree(s):
    nodes = [x.strip() for x in s.split(",")]
    root = None
    if len(nodes) == 0:
        return None
    root = create_node(nodes[0])
    parents = collections.deque([root])
    prev_is_left = False
    for i in xrange(1, len(nodes)):
        p = create_node(nodes[i])
        if prev_is_left:
            parents[0].right = p
            parents.popleft()
        else:
            parents[0].left = p
        prev_is_left = not prev_is_left
        if p is not None:
            parents.append(p)
    return root

# Your BSTIterator will be called like this:
# i, v = BSTIterator(root), []
# while i.hasNext(): v.append(i.next())
root = make_tree("4, 2, 6, 1, 3, 5, 7")
i, v = BSTIterator(root), []
while i.hasNext(): v.append(i.next())
print v
