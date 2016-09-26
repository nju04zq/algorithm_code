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

def test_tree_make_dump(s):
    root = make_tree(s)
    print dump_tree(root)

if __name__ == "__main__":
    test_tree_make_dump("1")
    test_tree_make_dump("1, #, 2")
    test_tree_make_dump("1, #, 2, 3, #, 4")
    test_tree_make_dump("1, 2, 3, 4, 5, 6, 7")
