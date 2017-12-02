import collections

# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

    def __repr__(self):
        return str(self.val)

class Solution(object):
    def lowestCommonAncestor(self, root, p, q):
        """
        :type root: TreeNode
        :type p: TreeNode
        :type q: TreeNode
        :rtype: TreeNode
        """
        if root is None or root == p or root == q:
            return root
        left = self.lowestCommonAncestor(root.left, p, q)
        right = self.lowestCommonAncestor(root.right, p, q)
        if left is None:
            return right
        elif right is None:
            return left
        else:
            return root
            
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

def find(root, i):
    if root is None:
        return None
    elif root.val == i:
        return root
    res = find(root.left, i)
    if res is not None:
        return res
    res = find(root.right, i)
    if res is not None:
        return res
    return None

def testLCA(s, i, j):
    root = make_tree(s)
    p = find(root, i)
    q = find(root, j)
    res = Solution().lowestCommonAncestor(root, p, q)
    print "In {0}, LCA for {1} and {2}, {3}".format(root, i, j, res.val)

testLCA("1, 2, 3, 4, 5, 6, 7", 3, 6)
testLCA("1, 2, 3, 4, 5, 6, 7", 4, 6)
