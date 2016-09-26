class Solution(object):
    def do_2sum(self, result, a, target, saved):
        i = 0
        j = len(a)-1
        while i < j:
            if i > 0 and a[i] == a[i-1]:
                i += 1
                continue
            if j < (len(a)-1) and a[j] == a[j+1]:
                j -= 1
                continue
            total = a[i] + a[j]
            if total == target:
                result.append(saved + [a[i], a[j]])
                i += 1
                j -= 1
            elif total < target:
                i += 1
            else:
                j -= 1

    def do_3sum(self, result, a, target, saved):
        for i in xrange(0, len(a)-2):
            if i > 0 and a[i] == a[i-1]:
                continue
            new_saved = saved + [a[i]]
            self.do_2sum(result, a[i+1:], target-a[i], new_saved)

    def fourSum(self, a, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[List[int]]
        """
        a.sort()
        result = []
        for i in xrange(0, len(a)-3):
            if i > 0 and a[i] == a[i-1]:
                continue
            self.do_3sum(result, a[i+1:], target-a[i], [a[i]])
        return result

def test_4sum(solution, a, target):
    result = solution.fourSum(a, target)
    print "4sum for {}".format(a)
    print result

solution = Solution()
test_4sum(solution, [1, 0, -1, 0, -2, 2], 0)
test_4sum(solution, [0, 0, 0, 0, 0, 0, 0], 0)
