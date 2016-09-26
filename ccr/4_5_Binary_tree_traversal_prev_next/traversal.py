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

def find_node(root, val):
    stack = [root]
    while len(stack) > 0:
        node = stack.pop()
        if node is None:
            continue
        if node.val == val:
            return node
        stack.append(node.left)
        stack.append(node.right)
    return None

def preorder_next(node):
    if node.left is not None:
        return node.left.val
    elif node.right is not None:
        return node.right.val

    while node.parent is not None:
        if node is node.parent.left and node.parent.right is not None:
            return node.parent.right.val
        node = node.parent

    return -1

def preorder_prev(node):
    if node.parent is None:
        return -1
    elif node is node.parent.left:
        return node.parent.val
    elif node.parent.left is None:
        return node.parent.val

    node = node.parent.left
    while node is not None:
        prev = node
        if node.right is not None:
            node = node.right
        else:
            node = node.left
    return prev.val

def left_most(root):
    node = root
    while node.left is not None:
        node = node.left
    return node

def inorder_next(node):
    if node.right is not None:
        lmost_node = left_most(node.right)
        return lmost_node.val

    while node.parent is not None:
        if node is node.parent.left:
            return node.parent.val
        node = node.parent

    return -1

def right_most(root):
    node = root
    while node.right is not None:
        node = node.right
    return node

def inorder_prev(node):
    if node.left is not None:
        rmost_node = right_most(node.left)
        return rmost_node.val

    while node.parent is not None:
        if node is node.parent.right:
            return node.parent.val
        node = node.parent

    return -1

def postorder_next(node):
    if node.parent is None:
        return -1
    if node is node.parent.right:
        return node.parent.val
    elif node.parent.right is None:
        return node.parent.val

    node = node.parent.right
    while node is not None:
        prev = node
        if node.left is None:
            node = node.right
        else:
            node = node.left
    return prev.val

def postorder_prev(node):
    if node.right is not None:
        return node.right.val
    elif node.left is not None:
        return node.left.val

    while node.parent is not None:
        if node is node.parent.right and node.parent.left is not None:
            return node.parent.left.val
        node = node.parent

    return -1

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

def preorder_recursive(root, result):
    if root is None:
        return
    result.append(root.val)
    preorder_recursive(root.left, result)
    preorder_recursive(root.right, result)

def inorder(root):
    result = []
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

def inorder_recursive(root, result):
    if root is None:
        return
    inorder_recursive(root.left, result)
    result.append(root.val)
    inorder_recursive(root.right, result)

def postorder(root):
    result = []
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

def postorder_recursive(root, result):
    if root is None:
        return
    postorder_recursive(root.left, result)
    postorder_recursive(root.right, result)
    result.append(root.val)

def test_preorder_next(samples):
    for s in samples:
        root = make_tree(s)
        a = preorder(root)
        for i in xrange(len(a)):
            node = find_node(root, a[i])
            result = preorder_next(node)
            if i == len(a) - 1:
                answer = -1
            else:
                answer = a[i+1]
            if result != answer:
                print "tree {}, preorder next for {}, get {}, shoule be {}".\
                      format(s, a[i], result, answer)

def test_preorder_prev(samples):
    for s in samples:
        root = make_tree(s)
        a = preorder(root)
        for i in xrange(len(a)):
            node = find_node(root, a[i])
            result = preorder_prev(node)
            if i == 0:
                answer = -1
            else:
                answer = a[i-1]
            if result != answer:
                print "tree {}, preorder prev for {}, get {}, shoule be {}".\
                      format(s, a[i], result, answer)

def test_inorder_next(samples):
    for s in samples:
        root = make_tree(s)
        a = inorder(root)
        for i in xrange(len(a)):
            node = find_node(root, a[i])
            result = inorder_next(node)
            if i == len(a) - 1:
                answer = -1
            else:
                answer = a[i+1]
            if result != answer:
                print "tree {}, inorder next for {}, get {}, shoule be {}".\
                      format(s, a[i], result, answer)

def test_inorder_prev(samples):
    for s in samples:
        root = make_tree(s)
        a = inorder(root)
        for i in xrange(len(a)):
            node = find_node(root, a[i])
            result = inorder_prev(node)
            if i == 0:
                answer = -1
            else:
                answer = a[i-1]
            if result != answer:
                print "tree {}, inorder prev for {}, get {}, shoule be {}".\
                      format(s, a[i], result, answer)

def test_postorder_next(samples):
    for s in samples:
        root = make_tree(s)
        a = postorder(root)
        for i in xrange(len(a)):
            node = find_node(root, a[i])
            result = postorder_next(node)
            if i == len(a) - 1:
                answer = -1
            else:
                answer = a[i+1]
            if result != answer:
                print "tree {}, postorder next for {}, get {}, shoule be {}".\
                      format(s, a[i], result, answer)

def test_postorder_prev(samples):
    for s in samples:
        root = make_tree(s)
        a = postorder(root)
        for i in xrange(len(a)):
            node = find_node(root, a[i])
            result = postorder_prev(node)
            if i == 0:
                answer = -1
            else:
                answer = a[i-1]
            if result != answer:
                print "tree {}, postorder prev for {}, get {}, shoule be {}".\
                      format(s, a[i], result, answer)

def test_preorder(root):
    result = preorder(root)
    answer = []
    preorder_recursive(root, answer)
    if result != answer:
        print "preorder for {}, get {}, shoule be {}".format(\
              dump_tree(root), result, answer)

def test_inorder(root):
    result = inorder(root)
    answer = []
    inorder_recursive(root, answer)
    if result != answer:
        print "inorder for {}, get {}, shoule be {}".format(\
              dump_tree(root), result, answer)

def test_postorder(root):
    result = postorder(root)
    answer = []
    postorder_recursive(root, answer)
    if result != answer:
        print "postorder for {}, get {}, shoule be {}".format(\
              dump_tree(root), result, answer)

def test_traversal():
    samples = ["1, #, 2", "1, #, 2, 3, #, 4", "1, 2, 3, 4, 5, 6, 7"]
    for s in samples:
        root = make_tree(s)
        test_preorder(root)
        test_inorder(root)
        test_postorder(root)

if __name__ == "__main__":
    test_traversal()
    samples = ["1, 2, #, 3", "1, #, 2, #, 3",  "1, 2, #, #, 3", "1, 2, 3, 4, 5",
               "1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15"]
    test_preorder_next(samples)
    test_preorder_prev(samples)
    test_inorder_prev(samples)
    test_inorder_next(samples)
    test_postorder_prev(samples)
    test_postorder_next(samples)
