import sys

# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def depth(self, p):
        d = 0
        while p is not None:
            d += 1
            p = p.left
        return d

    def countNodes(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        if root is None:
            return 0
        elif root.left is None:
            return 1
        elif root.right is None:
            return 2
        dLeft = self.depth(root.left)
        dRight = self.depth(root.right)
        if dLeft == dRight:
            leftCnt = (1<<dLeft) - 1
            rightCnt = self.countNodes(root.right)
        else:
            leftCnt = self.countNodes(root.left)
            rightCnt = (1<<dRight) - 1
        return leftCnt + rightCnt + 1

import collections

class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

def create_node(s):
    if s == "#":
        return None
    node = TreeNode(int(s))
    return node

def dump_node(node):
    if node is None:
        return "#"
    else:
        return str(node.val)

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

def test_tree_nodes(n):
    s = ", ".join([str(i+1) for i in xrange(n)])

    root = make_tree(s)
    cnt = Solution().countNodes(root)
    if cnt != n:
        raise Exception("{0}, get {1}, expect {2}".format(s, cnt, n))

for i in xrange(10000):
    sys.stdout.write("\r{0:03d}".format(i))
    test_tree_nodes(i+1)
print
