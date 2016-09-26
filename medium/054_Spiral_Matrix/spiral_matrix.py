class Solution(object):
    def insert_row(self, a, i, j, n, result, reverse):
        temp = []
        for k in xrange(j, j+n):
            temp.append(a[i][k])
        if reverse:
            result += temp[::-1]
        else:
            result += temp

    def insert_col(self, a, i, j, m, result, reverse):
        if m <= 0:
            return
        temp = []
        for k in xrange(i, i+m):
            temp.append(a[k][j])
        if reverse:
            result += temp[::-1]
        else:
            result += temp

    def spiral_order_internal(self, a, i, j, m, n, result):
        if m == 1:
            self.insert_row(a, i, j, n, result, reverse=False)
            return;
        if n == 1:
            self.insert_col(a, i, j, m, result, reverse=False)
            return;

        self.insert_row(a, i, j, n, result, reverse=False)
        self.insert_col(a, i+1, j+n-1, m-2, result, reverse=False)
        self.insert_row(a, i+m-1, j, n, result, reverse=True)
        self.insert_col(a, i+1, j, m-2, result, reverse=True)

        if m <= 2 or n <= 2:
            return

        self.spiral_order_internal(a, i+1, j+1, m-2, n-2, result);

    def spiralOrder(self, a):
        """
        :type matrix: List[List[int]]
        :rtype: List[int]
        """
        if len(a) == 0:
            return []

        m = len(a)
        n = len(a[0])
        result = []
        self.spiral_order_internal(a, 0, 0, m, n, result)
        return result

solution = Solution()

a = [[1]]
result = solution.spiralOrder(a)
print a
print result

a = [[1, 2, 3]]
result = solution.spiralOrder(a)
print a
print result

a = [[1], [2], [3]]
result = solution.spiralOrder(a)
print a
print result

a= [[ 1, 2, 3, 0 ],
    [ 4, 5, 6, 0 ],
    [ 7, 8, 9, 0 ]]
result = solution.spiralOrder(a)
print a
print result
