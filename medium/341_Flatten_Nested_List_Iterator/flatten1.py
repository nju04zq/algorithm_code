# """
# This is the interface that allows for creating nested lists.
# You should not implement it, or speculate about its implementation
# """
#class NestedInteger(object):
#    def isInteger(self):
#        """
#        @return True if this NestedInteger holds a single integer, rather than a nested list.
#        :rtype bool
#        """
#
#    def getInteger(self):
#        """
#        @return the single integer that this NestedInteger holds, if it holds a single integer
#        Return None if this NestedInteger holds a nested list
#        :rtype int
#        """
#
#    def getList(self):
#        """
#        @return the nested list that this NestedInteger holds, if it holds a nested list
#        Return None if this NestedInteger holds a single integer
#        :rtype List[NestedInteger]
#        """

class NestedInteger(object):
    def __init__(self, val):
        self.val = val

    def isInteger(self):
        return isinstance(self.val, int)

    def getInteger(self):
        return self.val

    def getList(self):
        return self.val

    def __repr__(self):
        if self.isInteger():
            return str(self.val)
        else:
            return "L{0}".format(len(self.val))

class NestedIterator(object):
    def __init__(self, nestedList):
        """
        Initialize your data structure here.
        :type nestedList: List[NestedInteger]
        """
        self.stack = []
        if len(nestedList) == 0:
            return
        self.stack.append([nestedList, 0])

    def next(self):
        """
        :rtype: int
        """
        l, i = self.stack[-1][0], self.stack[-1][1]
        val = l[i].getInteger()
        self.stack[-1][1] += 1
        return val

    def hasNext(self):
        """
        :rtype: bool
        """
        while len(self.stack) > 0:
            l, i = self.stack[-1][0], self.stack[-1][1]
            if i >= len(l):
                self.stack.pop()
                continue
            if l[i].isInteger():
                return True
            self.stack[-1][1] += 1
            self.stack.append([l[i].getList(), 0])
        return False

def dfs(val):
    if isinstance(val, int):
        return NestedInteger(val)
    l = []
    for v in val:
        l.append(dfs(v))
    return NestedInteger(l)

def generate(val):
    l = []
    for v in val:
        l.append(dfs(v))
    return l

def test(val):
    print val
    nestedList = generate(val)
    i, v = NestedIterator(nestedList), []
    while i.hasNext():
        v.append(i.next())
    print v

# Your NestedIterator object will be instantiated and called as such:
# i, v = NestedIterator(nestedList), []
# while i.hasNext(): v.append(i.next())
test([1, [2, [3, 4], 5], 6])
test([])
test([[], [2, 3], [], []])
test([[], [], []])
