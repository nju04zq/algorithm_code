class Solution(object):
    def dfs(self, cur, n, res):
        if cur == 0:
            start = 1
        else:
            start = 0
        for i in xrange(start, 10):
            j = cur + i
            if j > n:
                break
            res.append(j)
            if j * 10 <= n:
                self.dfs(j*10, n, res)

    def lexicalOrder(self, n):
        """
        :type n: int
        :rtype: List[int]
        """
        res = []
        self.dfs(0, n, res)
        return res

def testSolution(sol, n):
    res = sol.lexicalOrder(n)
    print "{0}, get {1}, len {2}".format(n, res, len(res))
    #print "{0}, get len {1}".format(n, len(res))

sol = Solution()
testSolution(sol, 13)
testSolution(sol, 4)
testSolution(sol, 121)
testSolution(sol, 1003)
testSolution(sol, 193)
