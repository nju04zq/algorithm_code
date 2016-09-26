import collections

class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.parent = None
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
        if p is not None:
            p.parent = parents[0]
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

def path_sum_up_internal(root, total, path):
    if root is None:
        return

    path.append(root.val)

    path_total = 0
    for i in xrange(len(path)-1, -1, -1):
        path_total += path[i]
        if path_total == total:
            print "Path: {}".format(path[i:])

    path_sum_up_internal(root.left, total, path)
    path_sum_up_internal(root.right, total, path)

    path.pop()


def path_sum_up(root, total):
    if root is None:
        return
    path = []
    path_sum_up_internal(root, total, path)

if __name__ == "__main__":
    s = "3, 2, 4, -4, 3, 2, 1, 3, #, #, #, #, #, #, #, 1, #, 3"
    root = make_tree(s)
    path_sum_up(root, 5)
