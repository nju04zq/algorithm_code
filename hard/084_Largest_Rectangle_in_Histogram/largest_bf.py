class Solution(object):
    def largestRectangleArea(self, heights):
        """
        :type heights: List[int]
        :rtype: int
        """
        maxArea = 0
        for i in xrange(len(heights)):
            minHeight = -1
            for j in xrange(i, len(heights)):
                height = heights[j]
                if minHeight == -1 or height < minHeight:
                    minHeight = height
                maxArea = max(maxArea, minHeight * (j-i+1))
        return maxArea
