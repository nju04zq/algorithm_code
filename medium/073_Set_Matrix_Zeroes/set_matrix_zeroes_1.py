class Solution(object):
    def set_matrix_row_zero(self, matrix, n, i):
        for j in xrange(n):
            matrix[i][j] = 0

    def set_matrix_col_zero(self, matrix, m, j):
        for i in xrange(m):
            matrix[i][j] = 0

    def is_set_row0_zero(self, matrix, n):
        for j in xrange(n):
            if matrix[0][j] == 0:
                return True
        return False

    def is_set_col0_zero(self, matrix, m):
        for i in xrange(m):
            if matrix[i][0] == 0:
                return True
        return False

    def setZeroes(self, matrix):
        """
        :type matrix: List[List[int]]
        :rtype: void Do not return anything, modify matrix in-place instead.
        """
        if len(matrix) == 0 or len(matrix[0]) == 0:
            return
        m, n = len(matrix), len(matrix[0])
        row0_zero = self.is_set_row0_zero(matrix, n)
        col0_zero = self.is_set_col0_zero(matrix, m)
        for i in xrange(1, m):
            for j in xrange(1, n):
                if matrix[i][j] == 0:
                    matrix[0][j], matrix[i][0] = 0, 0
        for i in xrange(1, m):
            if matrix[i][0] == 0:
                self.set_matrix_row_zero(matrix, n, i)
        for j in xrange(1, n):
            if matrix[0][j] == 0:
                self.set_matrix_col_zero(matrix, m, j)
        if row0_zero:
            self.set_matrix_row_zero(matrix, n, 0)
        if col0_zero:
            self.set_matrix_col_zero(matrix, m, 0)

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
