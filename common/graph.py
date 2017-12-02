class UndirectedGraphNode:
    def __init__(self, x):
        self.label = x
        self.neighbors = []

    def __str__(self):
        return str(self.label)

    def __repr__(self):
        return str(self.label)

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

if __name__ == "__main__":
    node = make_graph("0,1,2#1,0,2,3#2,0,1,3#3,1,2")
    dump_graph(node)
