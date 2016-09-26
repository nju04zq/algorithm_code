import sys
sys.path.append("../../common/")
from tree import TreeNode, make_tree, dump_tree

# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def inorderTraversal(self, root):
        """
        :type root: TreeNode
        :rtype: List[int]
        """
        result = []
        if root is None:
            return result
        stack = [(root, False)]
        while len(stack) > 0:
            node, visited = stack.pop()
            if node is None:
                continue
            if visited:
                result.append(node.val)
                stack.append((node.right, False))
            else:
                stack.append((node, True))
                stack.append((node.left, False))
        return result
    
    def preorderTraversal(self, root):
        result = []
        if root is None:
            return result
        stack = [root]
        while len(stack) > 0:
            node = stack.pop()
            if node is None:
                continue
            result.append(node.val)
            stack.append(node.right)
            stack.append(node.left)
        return result

    def postorderTraversal(self, root):
        result = []
        if root is None:
            return result
        stack = [(root, False)]
        while len(stack) > 0:
            node, visited = stack.pop()
            if node is None:
                continue
            if visited:
                result.append(node.val)
            else:
                stack.append((node, True))
                stack.append((node.right, False))
                stack.append((node.left, False))
        return result

def inorder_traversal_recursive(root, result):
    if root is None:
        return
    inorder_traversal_recursive(root.left, result)
    result.append(root.val)
    inorder_traversal_recursive(root.right, result)

def preorder_traversal_recursive(root, result):
    if root is None:
        return
    result.append(root.val)
    preorder_traversal_recursive(root.left, result)
    preorder_traversal_recursive(root.right, result)

def postorder_traversal_recursive(root, result):
    if root is None:
        return
    postorder_traversal_recursive(root.left, result)
    postorder_traversal_recursive(root.right, result)
    result.append(root.val)

def test_traversal(s):
    root = make_tree(s)
    result = solution.inorderTraversal(root)
    answer = []
    inorder_traversal_recursive(root, answer)
    if result != answer:
        print "Inorder result for {}".format(dump_tree(root))
        print "Get {}, shoule be {}".format(result, answer)
    result = solution.preorderTraversal(root)
    answer = []
    preorder_traversal_recursive(root, answer)
    if result != answer:
        print "Preorder result for {}".format(dump_tree(root))
        print "Get {}, shoule be {}".format(result, answer)
    result = solution.postorderTraversal(root)
    answer = []
    postorder_traversal_recursive(root, answer)
    if result != answer:
        print "Postorder result for {}".format(dump_tree(root))
        print "Get {}, shoule be {}".format(result, answer)

solution = Solution()
test_traversal("1")
test_traversal("1, 2")
test_traversal("1, #, 2")
test_traversal("1, #, 2, 3, 4, 5")
