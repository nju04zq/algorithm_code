class Solution(object):
    def trap(self, a):
        """
        :type height: List[int]
        :rtype: int
        """
        if len(a) == 0:
            return 0

        left, right = 0, len(a)-1
        total = 0
        while left < right:
            if a[left] < a[right]:
                left_barrier = a[left]
                left += 1
                while a[left] < left_barrier:
                    total += (left_barrier - a[left])
                    left += 1
            else:
                right_barrier = a[right]
                right -= 1
                while a[right] < right_barrier:
                    total += (right_barrier - a[right])
                    right -= 1
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
