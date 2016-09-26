class Solution(object):
    def trap(self, a):
        """
        :type height: List[int]
        :rtype: int
        """
        if len(a) == 0:
            return 0

        stack = []
        left = -1
        total = 0
        for i in xrange(len(a)):
            if left == -1:
                left = a[i]
                stack.append(a[i])
                continue
            if a[i] < left:
                stack.append(a[i])
                continue
            x = len(stack)*left - sum(stack)
            total += x
            left = a[i]
            stack = [left]

        right = stack.pop()
        while len(stack) > 0:
            x = stack.pop()
            if x >= right:
                right = x
            else:
                total += (right - x)

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
