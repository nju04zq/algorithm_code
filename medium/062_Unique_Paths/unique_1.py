class Solution(object):
    def unique_internal(self, a):
        m, n = len(a), len(a[0])

        f = [0 for i in xrange(n)]

        for j in xrange(n-1, -1, -1):
            if a[m-1][j] == 1:
                break
            f[j] = 1

        for i in xrange(m-2, -1, -1):
            if a[i][n-1] == 1:
                f[n-1] = 0
            for j in xrange(n-2, -1, -1):
                if a[i][j] == 1:
                    f[j] = 0
                else:
                    f[j] += f[j+1]
        return f[0]

    def uniquePathsWithObstacles(self, a):
        """
        :type obstacleGrid: List[List[int]]
        :rtype: int
        """
        if len(a) == 0:
            return 0

        result = self.unique_internal(a)
        return result

def test_unique(a):
    result = solution.uniquePathsWithObstacles(a);
    print "Obstacles {}, result {}".format(a, result)

def test_case_1():
    a = [[0, 0, 0], [0, 1, 0], [0, 0, 0]]
    test_unique(a)

def test_case_2():
    a = [[0, 0, 0], [0, 0, 0], [0, 0, 0]]
    test_unique(a)

solution = Solution()
test_case_1()
test_case_2()

