class Solution(object):
    def reverse(self, a):
        i, j = 0, len(a)-1
        while i < j:
            a[i], a[j] = a[j], a[i]
            i += 1
            j -= 1

    def rotate(self, matrix):
        """
        :type matrix: List[List[int]]
        :rtype: void Do not return anything, modify matrix in-place instead.
        """
        n = len(matrix)
        for i in xrange(n):
            for j in xrange(i+1, n):
                matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        for i in xrange(n):
            self.reverse(matrix[i])

solution = Solution()

a = [[]]
print "Before rotate, {}".format(a)
solution.rotate(a)
print "After rotate, {}".format(a)

a = [[1]]
print "Before rotate, {}".format(a)
solution.rotate(a)
print "After rotate, {}".format(a)

a = [[1, 2], [3, 4]]
print "Before rotate, {}".format(a)
solution.rotate(a)
print "After rotate, {}".format(a)

a = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
print "Before rotate, {}".format(a)
solution.rotate(a)
print "After rotate, {}".format(a)
