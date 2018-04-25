class Solution(object):
    def jump(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        dp = [len(nums)+1 for i in xrange(len(nums))]
        dp[0] = 0
        for i in xrange(1, len(nums)):
            for j in xrange(0, i):
                if j+nums[j] >= i:
                    dp[i] = min(dp[i], dp[j]+1)
        return dp[len(nums)-1]
