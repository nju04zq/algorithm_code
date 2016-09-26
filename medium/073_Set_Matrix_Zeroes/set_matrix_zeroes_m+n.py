class Solution(object):
    def set_matrix_row_zero(self, matrix, n, i):
        for j in xrange(n):
            matrix[i][j] = 0

    def set_matrix_col_zero(self, matrix, m, j):
        for i in xrange(m):
            matrix[i][j] = 0

    def setZeroes(self, matrix):
        """
        :type matrix: List[List[int]]
        :rtype: void Do not return anything, modify matrix in-place instead.
        """
        if len(matrix) == 0 or len(matrix[0]) == 0:
            return
        m, n = len(matrix), len(matrix[0])
        rows = [0 for i in xrange(m)]
        cols = [0 for i in xrange(n)]
        for i in xrange(m):
            for j in xrange(n):
                if matrix[i][j] == 0:
                    rows[i], cols[j] = 1, 1
        for i in xrange(m):
            if rows[i] == 1:
                self.set_matrix_row_zero(matrix, n, i)
        for j in xrange(n):
            if cols[j] == 1:
                self.set_matrix_col_zero(matrix, m, j)

def test_set_zeroes(a):
    print "Before set, {}".format(a)
    solution.setZeroes(a)
    print "After set, {}".format(a)

solution = Solution()

a = [[1, 2, 3], [4, 0, 6], [0, 8, 9]]
test_set_zeroes(a)

a = [[1]]
test_set_zeroes(a)

a = [[0]]
test_set_zeroes(a)

a = [[1, 2]]
test_set_zeroes(a)

a = [[1, 0]]
test_set_zeroes(a)

a = [[1], [2]]
test_set_zeroes(a)

a = [[1], [0]]
test_set_zeroes(a)
