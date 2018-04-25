class Solution(object):
    def trap(self, height):
        """
        :type height: List[int]
        :rtype: int
        """
        leftMax = [0 for i in xrange(len(height))]
        maxHeight = 0
        for i in xrange(len(height)):
            leftMax[i] = maxHeight
            maxHeight = max(maxHeight, height[i])
        total, maxHeight = 0, 0
        for i in xrange(len(height)-1, -1, -1):
            minHeight = min(maxHeight, leftMax[i])
            if minHeight > height[i]:
                total += (minHeight-height[i])
            maxHeight = max(maxHeight, height[i])
        return total
