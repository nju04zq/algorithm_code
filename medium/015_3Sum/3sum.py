class Solution(object):
    def two_sum(self, result, a, target):
        i = 0
        j = len(a)-1
        while i < j:
            if i > 0 and a[i] == a[i-1]:
                i += 1
                continue
            if j < len(a)-1 and a[j] == a[j+1]:
                j -= 1
                continue

            total_2 = a[i] + a[j]
            if total_2 == target:
                result.append([-target, a[i], a[j]])
                i += 1
                j -= 1
            elif total_2 < target:
                i += 1
            else:
                j -= 1

    def threeSum(self, a):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        result = []
        a.sort()
        for i in xrange(0, len(a)-2):
            if i > 0 and a[i] == a[i-1]:
                continue
            self.two_sum(result, a[i+1:], -a[i])

        return result

def test_3sum(solution, a):
    result = solution.threeSum(a)
    print "Array {}".format(a)
    print result

solution = Solution()
test_3sum(solution, [0, 0, 0, 0, 0])
test_3sum(solution, [-1, 0, 1, 2, -1, -4])
test_3sum(solution, [-2, 0, 0, 2, 2])

