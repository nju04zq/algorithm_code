# log(m*n) = log(m) + log(n)
class Solution(object):
    def get_val_via_idx(self, matrix, x, m, n):
        i, j = x/n, x%n
        return matrix[i][j]

    def search_matrix_internal(self, matrix, m, n, start, end, target):
        if start > end:
            return False

        mid = (start + end)/2
        val = self.get_val_via_idx(matrix, mid, m, n) 
        if val == target:
            return True
        elif val > target:
            end = mid - 1
        else:
            start = mid + 1

        result = self.search_matrix_internal(matrix, m, n, start, end, target)
        return result

    def searchMatrix(self, matrix, target):
        """
        :type matrix: List[List[int]]
        :type target: int
        :rtype: bool
        """
        m = len(matrix) 
        n = len(matrix[0])
        if m == 0 or n == 0:
            return False

        start, end = 0, m*n-1
        result = self.search_matrix_internal(matrix, m, n, start, end, target)
        return result

def test_search_matrix(a, target, answer):
    result = solution.searchMatrix(a, target)
    if result != answer:
        print "{}, {}, get {}, should be {}".format(a, target, result, answer)

solution = Solution()

a = [[1,   3,  5,  7],
     [10, 11, 16, 20],
     [23, 30, 34, 50]]
test_search_matrix(a, 3, True)
test_search_matrix(a, 1, True)
test_search_matrix(a, 50, True)
test_search_matrix(a, 20, True)
test_search_matrix(a, 21, False)
test_search_matrix(a, 2, False)
test_search_matrix(a, 51, False)
test_search_matrix(a, 0, False)
