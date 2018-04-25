# TLE
class Solution(object):
    def isSubsequence(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: bool
        """
        m, n = len(t), len(s)
        if n == 0:
            return True
        elif m == 0:
            return False
        dp = [[False for j in xrange(n)] for i in xrange(m)]
        if s[0] == t[0]:
            dp[0][0] = True
        for i in xrange(1, m):
            if dp[i-1][0] or t[i] == s[0]:
                dp[i][0] = True
            for j in xrange(1, n):
                if dp[i-1][j] or (dp[i-1][j-1] and t[i] == s[j]):
                    dp[i][j] = True
        return dp[m-1][n-1]
