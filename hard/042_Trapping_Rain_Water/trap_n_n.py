class Solution(object):
    def calc_left_max(self, a):
        left = []
        left_max = 0
        for i in xrange(len(a)):
            if a[i] > left_max:
                left_max = a[i]
            left.append(left_max)
        return left

    def calc_right_max(self, a):
        right = [0 for i in xrange(len(a))]
        right_max = 0
        for i in xrange(len(a)-1, -1, -1):
            if a[i] > right_max:
                right_max = a[i]
            right[i] = right_max
        return right

    def trap(self, a):
        """
        :type height: List[int]
        :rtype: int
        """
        if len(a) == 0:
            return 0

        left = self.calc_left_max(a)
        right = self.calc_right_max(a)
        total = 0
        for i in xrange(len(a)):
            barrier = min(left[i], right[i])
            if barrier > a[i]:
                total += (barrier - a[i])
        return total

def test_trap(a, answer):
    result = solution.trap(a)
    print a
    if result != answer:
        print "Get {}, should be {}".format(result, answer)
    else:
        print "Get {}".format(result)

solution = Solution()
a = [0,1,0,2,1,0,1,3,2,1,2,1]
test_trap(a, 6)
a = [4, 2, 3]
test_trap(a, 1)
a = [0]
test_trap(a, 0)
