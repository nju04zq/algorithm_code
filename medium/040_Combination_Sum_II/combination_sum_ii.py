class Solution(object):
    def combination_sum_internal(self, a, target, start, result, path):
        total = sum(path)
        if total == target:
            result.append(path[:])
            return
        elif total > target:
            return

        for i in xrange(start, len(a)):
            if i > start and a[i] == a[i-1]: ##careful, i > start, not i > 0
                continue
            path.append(a[i])
            self.combination_sum_internal(a, target, i+1, result, path)
            path.pop()

    def combinationSum2(self, a, target):
        """
        :type candidates: List[int]
        :type target: int
        :rtype: List[List[int]]
        """
        a.sort()
        start = 0
        result, path = [], []
        self.combination_sum_internal(a, target, start, result, path)
        return result

def test_combination_sum(a, target):
    result = solution.combinationSum2(a, target)
    print "target {}, from {}".format(a, target)
    for x in result:
        print x

solution = Solution()
test_combination_sum([10,1,2,7,6,1,5], 8)
