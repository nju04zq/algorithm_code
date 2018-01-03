import collections

# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Codec:
    def serialize(self, root):
        """Encodes a tree to a single string.
        
        :type root: TreeNode
        :rtype: str
        """
        res = []
        nodes = [root]
        while len(nodes) > 0:
            cnt = len(nodes)
            for i in xrange(cnt):
                node = nodes[i]
                if node is None:
                    res.append("None")
                else:
                    res.append(str(node.val))
                    nodes.append(node.left)
                    nodes.append(node.right)
            nodes = nodes[cnt:]
        for j in xrange(len(res)-1, -1, -1):
            if res[j] != "None":
                break
        if j > 0:
            res = res[:j+1]
        return ",".join(res)
        
    def deserialize(self, data):
        """Decodes your encoded data to tree.
        
        :type data: str
        :rtype: TreeNode
        """
        vals = data.split(",")
        root = self.make_tree_node(vals[0])
        nodes, i = [root], 1
        while i < len(vals): 
            cnt = len(nodes)
            j, leftNode = 0, True
            while j < cnt:
                if i >= len(vals):
                    break
                node = self.make_tree_node(vals[i]) 
                if leftNode:
                    nodes[j].left = node
                else:
                    nodes[j].right = node
                    j+=1
                if node is not None:
                    nodes.append(node)
                leftNode = not leftNode
                i+=1
            nodes = nodes[cnt:]
        return root

    def make_tree_node(self, val):
        if val == "None":
            return None
        else:
            return TreeNode(int(val))

def dump_node(node):
    if node is None:
        return "#"
    else:
        return str(node.val)

def dump_tree(root):
    if root is None:
        return ""
    output = []
    queue = collections.deque([root])
    while len(queue) > 0:
        node = queue.popleft()
        output.append(dump_node(node))
        if node is not None:
            queue.append(node.left)
            queue.append(node.right)
    i = len(output) - 1
    while i >= 0:
        if output[i] != "#":
            break
        else:
            del output[i]
        i -= 1
    return ", ".join(output)

# Your Codec object will be instantiated and called as such:
# codec = Codec()
# codec.deserialize(codec.serialize(root))
codec = Codec()
root = codec.deserialize("1,2,3,4,None,5")
#print dump_tree(root)
print codec.serialize(root)
