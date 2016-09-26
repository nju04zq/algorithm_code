class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

class MinHeap(object):
    def __init__(self, nodes):
        self.nodes = []
        self.size = 0
        self.init_nodes(nodes)
    
    def calc_parent(self, i):
        return (i+1)/2-1

    def calc_lchild(self, i):
        return 2*(i+1)-1

    def calc_rchild(self, i):
        return 2*(i+1)

    def get_val(self, i):
        return self.nodes[i].val

    def swap_node(self, i, j):
        self.nodes[i], self.nodes[j] = self.nodes[j], self.nodes[i]

    def slide_up(self, i):
        while (i > 0):
            parent = self.calc_parent(i)
            if self.get_val(i) >= self.get_val(parent):
                break
            self.swap_node(i, parent)
            i = parent

    def slide_down(self, i):
        while True:
            lchild = self.calc_lchild(i)
            rchild = self.calc_rchild(i)
            if lchild > (self.size-1):
                break
            if rchild > (self.size-1):
                if self.get_val(lchild) < self.get_val(i):
                    self.swap_node(lchild, i)
                break
            if self.get_val(lchild) < self.get_val(rchild):
                child = lchild
            else:
                child = rchild
            if self.get_val(child) >= self.get_val(i):
                break
            self.swap_node(i, child)
            i = child

    def init_nodes(self, nodes):
        for i in xrange(0, len(nodes)):
            if nodes[i] is not None:
                self.nodes.append(nodes[i])
                self.size += 1

        for i in xrange(0, self.size):
            self.slide_up(i)

    def get_min(self):
        if self.size == 0:
            return None
        else:
            return self.nodes[0]

    def reload_node(self, node):
        if node is None:
            node = self.nodes[self.size-1]
            self.size -= 1

        if self.size == 0:
            return

        self.nodes[0] = node
        self.slide_down(0)

class Solution(object):
    def mergeKLists(self, lists):
        """
        :type lists: List[ListNode]
        :rtype: ListNode
        """
        if lists is None or len(lists) == 0:
            return None

        prev = None
        head = None
        min_heap = MinHeap(lists)
        while True:
            node = min_heap.get_min()
            if node is None:
                break
            min_heap.reload_node(node.next)

            node.next = None
            if head is None:
                head = node
                prev = node
            else:
                prev.next = node
                prev = node

        return head

def dump_nodes(nodes):
    result = ""
    for node in nodes:
        result += "{} ".format(node.val)
    print "###{}###".format(result)

def test_merge_klists(lists):
    solution = Solution()
    p = solution.mergeKLists(lists)

    merged = ""
    while p is not None:
        merged += " {}".format(p.val)
        p = p.next
    print merged

def make_one_list(x):
    head = None
    prev = None
    for y in x:
        node = ListNode(y)
        if head is None:
            head = node
        else:
            prev.next = node
        prev = node
    return head

def make_klists(a):
    lists = []

    for x in a:
        y = make_one_list(x)
        lists.append(y)
    return lists

def test_case_0():
    a = [[0, 1], [5, 6, 7], [10], [15, 16], [20, 21, 22, 23]]
    lists = make_klists(a)
    test_merge_klists(lists)

def test_case_1():
    a = [[0, 1], [5, 6, 7], [10], [15, 16]]
    lists = make_klists(a)
    test_merge_klists(lists)

def test_case_2():
    a = [[-9,-7,-7],[-6,-4,-1,1],[-6,-5,-2,0,0,1,2],[-9,-8,-6,-5,-4,1,2,4],[-10],[-5,2,3]]
    lists = make_klists(a)
    test_merge_klists(lists)

#test_case_0()
#test_case_1()
test_case_2()
