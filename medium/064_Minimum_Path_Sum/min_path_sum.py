class Solution(object):
    def minPathSum(self, grid):
        """
        :type grid: List[List[int]]
        :rtype: int
        """
        if len(grid) == 0:
            return 0

        m, n = len(grid), len(grid[0])
        f = [0 for i in xrange(n)]

        f[0] = grid[0][0]
        for j in xrange(1, n):
            f[j] = f[j-1] + grid[0][j]

        for i in xrange(1, m):
            f[0] += grid[i][0]
            for j in xrange(1, n):
                f[j] = min(f[j-1], f[j]) + grid[i][j]
        return f[n-1]

def test_min_path_sum(a):
    result = solution.minPathSum(a)
    print "Grid {}, min {}".format(a, result)

def test_case_0():
    a = [[1]]
    test_min_path_sum(a)

def test_case_1():
    a = [[1, 2, 3]]
    test_min_path_sum(a)

def test_case_2():
    a = [[1], [2], [3]]
    test_min_path_sum(a)

def test_case_3():
    a = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
    test_min_path_sum(a)

solution = Solution()
test_case_0()
test_case_1()
test_case_2()
test_case_3()
