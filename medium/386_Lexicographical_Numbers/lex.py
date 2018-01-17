class Solution(object):
    def lexicalOrder(self, n):
        """
        :type n: int
        :rtype: List[int]
        """
        res = []
        for base in xrange(1, 10):
            x = base
            while x <= n:
                res.append(x)
                x *= 10
            x /= 10
            for y in xrange(1, x/base):
                z = x + y
                tmp = []
                while z%10 == 0:
                    tmp.append(z/10)
                    z /= 10
                res += tmp[::-1]
                if x+y <= n:
                    res.append(x+y)
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
