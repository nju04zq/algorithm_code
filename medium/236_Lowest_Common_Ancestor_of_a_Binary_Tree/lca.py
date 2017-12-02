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
    def get_path(self, root, p, q, path, path_p, path_q):
        if root is None:
            return
        path.append(root)
        if root == p:
            path_p[:] = path[:]
        if root == q:
            path_q[:] = path[:]
        self.get_path(root.left, p, q, path, path_p, path_q)
        self.get_path(root.right, p, q, path, path_p, path_q)
        path.pop()

    def lowestCommonAncestor(self, root, p, q):
        """
        :type root: TreeNode
        :type p: TreeNode
        :type q: TreeNode
        :rtype: TreeNode
        """
        path_p, path_q, path = [], [], []
        self.get_path(root, p, q, path, path_p, path_q)
        common = None
        for i in xrange(min(len(path_p), len(path_q))):
            if path_p[i] == path_q[i]:
                common = path_p[i]
            else:
                break
        return common
            
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
