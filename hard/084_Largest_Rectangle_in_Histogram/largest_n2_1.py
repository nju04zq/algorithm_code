class Solution(object):
    def largestRectangleArea(self, heights):
        """
        :type heights: List[int]
        :rtype: int
        """
        n = len(heights)
        if n == 0:
            return 0
        l, r = 0, n-1
        largest = 0
        while l <= r:
            area = min(heights[l:r+1]) * (r - l + 1)
            largest = max(largest, area)
            if heights[l] < heights[r]:
                l += 1
            else:
                r -= 1

        return largest

def test_largest(a):
    largest = solution.largestRectangleArea(a)
    print "{}, largest {}".format(a, largest)

solution = Solution()

a = [2, 1, 5, 6, 2, 3]
test_largest(a)

a = [2, 1, 2]
test_largest(a)
