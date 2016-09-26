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

def find_node(root, i):
    if root is None:
        return None
    if root.val == i:
        return root
    node = find_node(root.left, i)
    if node is not None:
        return node
    node = find_node(root.right, i)
    if node is not None:
        return node
    return None

def preorder(root):
    result = []
    stack = [root]
    while len(stack) > 0:
        node = stack.pop()
        if node is None:
            continue
        result.append(node.val)
        stack.append(node.right)
        stack.append(node.left)
    return result

def calc_depth(node):
    depth = 0
    while node is not None:
        depth += 1
        node = node.parent
    return depth

def slide_up(node, depth):
    while depth > 0:
        node = node.parent
        depth -= 1
    return node

def first_common_ancestor_via_parent(n1, n2):
    d1 = calc_depth(n1)
    d2 = calc_depth(n2)

    if d1 > d2:
        n1 = slide_up(n1, d1-d2)
    else:
        n2 = slide_up(n2, d2-d1)

    while n1 is not n2:
        n1 = n1.parent
        n2 = n2.parent

    return n1

def first_common_ancestor_recursive(root, n1, n2):
    if root is None:
        return None, None
    lp1, lp2 = first_common_ancestor_recursive(root.left, n1, n2)
    rp1, rp2 = first_common_ancestor_recursive(root.right, n1, n2)
    p1, p2 = None, None
    if root.left is n1:
        p1 = root
    elif root.left is n2:
        p2 = root
    if root.right is n1:
        p1 = root
    elif root.right is n2:
        p2 = root

def find_lca(root, n1, n2):
    if root is None:
        return None

    if root is n1 or root is n2:
        return root

    left_lca = find_lca(root.left, n1, n2)
    right_lca = find_lca(root.right, n1, n2)
    if left_lca is not None and right_lca is not None:
        return root
    elif left_lca is not None:
        return left_lca
    else:
        return right_lca

def test_first_common_ancestor():
    tree = "1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15"
    root = make_tree(tree)
    nodes = preorder(root)
    for i in xrange(len(nodes)):
        for j in xrange(i+1, len(nodes)):
            n1 = find_node(root, nodes[i])
            n2 = find_node(root, nodes[j])
            node = first_common_ancestor_via_parent(n1, n2)
            answer = find_lca(root, n1, n2)
            if node != answer:
                print "First common ancestor between {}/{}, get {}, should {}".\
                      format(n1.val, n2.val, node.val, answer.val)

if __name__ == "__main__":
    test_first_common_ancestor()

