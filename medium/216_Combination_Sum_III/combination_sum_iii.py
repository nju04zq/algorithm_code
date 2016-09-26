class Solution(object):
    def combination_sum3_internal(self, a, start, k, target, path, result):
        if len(path) > k:
            return
        total = sum(path)
        if len(path) == k:
            if total == target:
                result.append(path[:])
            else:
                return
        elif total >= target:
            return

        for i in xrange(start, len(a)):
            path.append(a[i])
            self.combination_sum3_internal(a, i+1, k, target, path, result)
            path.pop()

    def combinationSum3(self, k, n):
        """
        :type k: int
        :type n: int
        :rtype: List[List[int]]
        """
        a = xrange(1, 10)
        if n < 0 or k <= 0:
            return []
        start, path, result = 0, [], []
        self.combination_sum3_internal(a, start, k, n, path, result)
        return result

def test_combination_sum3(k, n):
    result = solution.combinationSum3(k, n)
    print "n {}, k {}".format(n, k)
    for x in result:
        print x

solution = Solution()
test_combination_sum3(3, 7)
test_combination_sum3(3, 9)
test_combination_sum3(2, 100)
test_combination_sum3(3, 15)
test_combination_sum3(4, 15)
