class Solution(object):
    def add_to_result(self, result, total, target):
        if len(result) == 0:
            result.append(total)
        elif abs(total-target) < abs(result[0]-target):
            result[0] = total

    def two_sum_closest(self, a, start, x0, target, result):
        i, j = start, len(a)-1
        while i < j:
            total = x0 + a[i] + a[j]
            self.add_to_result(result, total, target)
            if total == target:
                return
            elif total > target:
                j -= 1
            else:
                i += 1

    def threeSumClosest(self, a, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """
        result = []
        a.sort()
        for i in xrange(len(a)):
            self.two_sum_closest(a, i+1, a[i], target, result)
        return result[0]

def test_three_sum_closest(a, target):
    result = solution.threeSumClosest(a, target)
    print "In {}, 3 sum closest to {} is {}".format(a, target, result)

solution = Solution()
test_three_sum_closest([-1, 2, 1, -4], 1)

