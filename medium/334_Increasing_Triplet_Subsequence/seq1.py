class Solution(object):
    def increasingTriplet(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        if len(nums) < 3:
            return False
        minPrefix = [0 for i in xrange(len(nums))]
        minLeft = 0
        for i, num in enumerate(nums):
            minPrefix[i] = minLeft
            if i == 0 or num < minLeft:
                minLeft = num
        maxRight = nums[-1]
        for i in xrange(len(nums)-2, 0, -1):
            if minPrefix[i] < nums[i] < maxRight:
                return True
            maxRight = max(maxRight, nums[i])
        return False
