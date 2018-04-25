class Solution(object):
    def findLength(self, a, b):
        """
        :type A: List[int]
        :type B: List[int]
        :rtype: int
        """
        m, n = len(a), len(b)
        dp = [[0 for i in xrange(n+1)] for i in xrange(m+1)]
        maxLen = 0
        for i in xrange(1, m+1):
            for j in xrange(1, n+1):
                if a[i-1] == b[j-1]:
                    dp[i][j] = dp[i-1][j-1] + 1
                    maxLen = max(maxLen, dp[i][j])
        return maxLen
