class Solution(object):
    def add_to_result(self, result, mask, n):
        temp = []
        for i in xrange(n):
            if mask[i] == 1:
                temp.append(i+1)
        result.append(temp)

    def combine_internal(self, result, mask, start, cnt, n, k):
        if cnt == k:
            self.add_to_result(result, mask, n)
            return

        for i in xrange(start, n):
            mask[i] = 1
            self.combine_internal(result, mask, i+1, cnt+1, n, k)
            mask[i] = 0

    def combine(self, n, k):
        """
        :type n: int
        :type k: int
        :rtype: List[List[int]]
        """
        if n <= 0 or k > n:
            return []
        result = []
        mask = [0 for i in xrange(n)]
        start, cnt = 0, 0
        self.combine_internal(result, mask, start, cnt, n, k)
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
