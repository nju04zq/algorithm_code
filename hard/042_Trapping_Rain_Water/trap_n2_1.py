class Solution(object):
    def find_left_max(self, a, end):
        left = 0
        for i in xrange(end):
            if a[i] > left:
                left = a[i]
        return left

    def find_right_max(self, a, start):
        right = 0
        for i in xrange(start, len(a)):
            if a[i] > right:
                right = a[i]
        return right

    def trap(self, a):
        """
        :type height: List[int]
        :rtype: int
        """
        if len(a) == 0:
            return 0

        total = 0
        for i in xrange(len(a)):
            left = self.find_left_max(a, i)
            right = self.find_right_max(a, i+1)
            barrier = min(left, right)
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

