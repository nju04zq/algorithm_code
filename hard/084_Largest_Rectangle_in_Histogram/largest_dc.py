class Solution(object):
    def largestRectangleArea(self, heights):
        """
        :type heights: List[int]
        :rtype: int
        """
        return self.helper(heights)
    
    def helper(self, heights):
        if len(heights) == 0:
            return 0
        minHeight, minIdx = -1, -1
        for i, height in enumerate(heights):
            if minHeight == -1 or height < minHeight:
                minHeight, minIdx = height, i
        minLeft = self.helper(heights[:minIdx])
        minRight = self.helper(heights[minIdx+1:])
        minMid = len(heights) * minHeight
        return max(minLeft, max(minRight, minMid))
