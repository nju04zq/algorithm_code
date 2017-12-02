class UndirectedGraphNode:
    def __init__(self, x):
        self.label = x
        self.neighbors = []

class Solution:
    def get_node(self, val, tbl):
        if val in tbl:
            return tbl[val]
        else:
            node = UndirectedGraphNode(val)
            tbl[val] = node
            return node

    # @param node, a undirected graph node
    # @return a undirected graph node
    def cloneGraph(self, node):
        if node is None:
            return None
        visited = set()
        nodes = [node]
        first_node = None
        new_tbl = {}
        while len(nodes) > 0:
            cnt = len(nodes)
            for i in xrange(cnt):
                node = nodes[i]
                val = node.label
                if val in visited:
                    continue
                visited.add(val)
                new_node = self.get_node(val, new_tbl)
                if first_node is None:
                    first_node = new_node
                for neighbor in node.neighbors:
                    new_neighbor = self.get_node(neighbor.label, new_tbl)
                    new_node.neighbors.append(new_neighbor)
                    nodes.append(neighbor)
            nodes = nodes[cnt:]
        return first_node

def get_node(nodes, val):
    if val in nodes:
        node = nodes[val]
    else:
        node = UndirectedGraphNode(val)
        nodes[val] = node
    return node

# "0,1,2#1,0,2#2,0,1"
def make_graph(s):
    first_node = None
    nodes = {}
    for toks in s.split("#"):
        vals = toks.split(",")
        node = get_node(nodes, vals[0])
        if first_node is None:
            first_node = node
        for val in vals[1:]:
            neighbor = get_node(nodes, val)
            node.neighbors.append(neighbor)
    return first_node

def dump_graph(node):
    visited = set()
    nodes = [node]
    while len(nodes) > 0:
        cnt = len(nodes)
        for i in xrange(cnt):
            val = nodes[i].label
            if val not in visited:
                s = "{0}: ".format(val)
                for neighbor in nodes[i].neighbors:
                    s += "{0} ".format(neighbor.label)
                    nodes.append(neighbor)
                visited.add(val)
                print s
        nodes = nodes[cnt:]

def test_clone(s):
    node = make_graph(s)
    node1 = Solution().cloneGraph(node)
    print "original graph"
    dump_graph(node)
    print "cloned graph"
    dump_graph(node1)

if __name__ == "__main__":
    test_clone("0,1,2#1,0,2,3#2,0,1,3#3,1,2")
