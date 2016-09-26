class Solution(object):
    def combination_sum_internal(self, a, target, start, result, path):
        total = sum(path)
        if total == target:
            result.append(path[:])
            return
        elif total > target:
            return

        for i in xrange(start, len(a)):
            if i > start and a[i] == a[i-1]:
                continue
            path.append(a[i])
            self.combination_sum_internal(a, target, i, result, path)
            path.pop()

    def combinationSum(self, a, target):
        """
        :type candidates: List[int]
        :type target: int
        :rtype: List[List[int]]
        """
        a.sort()
        result, path = [], []
        start = 0
        self.combination_sum_internal(a, target, start, result, path)
        return result

def test_combination_sum(a, target):
    result = solution.combinationSum(a, target)
    print "Target {}, from {}".format(target, a)
    for x in result:
        print x

solution = Solution()
test_combination_sum([7, 3, 2, 2, 5], 7)
