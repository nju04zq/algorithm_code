class Solution(object):
    def lexicalOrder(self, n):
        """
        :type n: int
        :rtype: List[int]
        """
        res = []
        cur = 1
        for i in xrange(n):
            res.append(cur)
            if cur * 10 <= n:
                cur *= 10
            elif cur%10 != 9 and cur + 1 <= n:
                cur += 1
            else:
                while (cur/10) %10 == 9:
                    cur /= 10
                cur = cur/10 + 1
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
