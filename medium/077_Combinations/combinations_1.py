class Solution(object):
    def combine_internal(self, result, path, start, n, k):
        if len(path) == k:
            result.append(path[:])
            return

        for i in xrange(start, n):
            path.append(i+1)
            self.combine_internal(result, path, i+1, n, k)
            path.pop()

    def combine(self, n, k):
        """
        :type n: int
        :type k: int
        :rtype: List[List[int]]
        """
        if n <= 0 or k > n:
            return []
        start, path, result = 0, [], []
        self.combine_internal(result, path, start, n, k)
        return result

def test_combine(n, k):
    result = solution.combine(n, k)
    print "{}/{}".format(n, k)
    for a in result:
        print a

solution = Solution()
test_combine(1, 1)
test_combine(2, 1)
test_combine(2, 2)
test_combine(4, 1)
test_combine(4, 2)
test_combine(4, 3)
test_combine(4, 4)
