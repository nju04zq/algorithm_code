class Solution(object):
    def largestRectangleArea(self, heights):
        """
        :type heights: List[int]
        :rtype: int
        """
        n = len(heights)
        if n == 0:
            return 0
        max_area = 0
        for i in xrange(n):
            if i < n-1 and heights[i] <= heights[i+1]:
                continue
            min_height = heights[i]
            j = i
            while j >= 0:
                min_height = min(min_height, heights[j])
                area = min_height * (i - j + 1)
                max_area = max(max_area, area)
                j -= 1
        return max_area

def test_largest(a):
    largest = solution.largestRectangleArea(a)
    print "{}, largest {}".format(a, largest)

solution = Solution()

a = [2, 1, 5, 6, 2, 3]
test_largest(a)

a = [2, 1, 2]
test_largest(a)
