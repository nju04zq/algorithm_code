class Solution(object):
    def minDistance(self, s, t):
        """
        :type word1: str
        :type word2: str
        :rtype: int
        """
        if len(s) >= len(t):
            r = self.min_distance_internal(s, t)
        else:
            r = self.min_distance_internal(t, s)
        return r

    def min_distance_internal(self, s, t):
        dp = [i for i in xrange(len(t)+1)]
        for i in xrange(1, len(s)+1):
            prev = dp[0]
            dp[0] += 1
            for j in xrange(1, len(t)+1):
                temp = dp[j]
                if s[i-1] == t[j-1]:
                    dp[j] = prev
                else:
                    dp[j] = min(dp[j-1]+1, dp[j]+1)
                    dp[j] = min(dp[j], prev+1)
                prev = temp
        return dp[-1]

def test_edit_distance(a, b):
    r = solution.minDistance(a, b)
    print "{}, {}, edit distance {}".format(a, b, r)

solution = Solution()

a = ""
b = ""
test_edit_distance(a, b)

a = ""
b = "a"
test_edit_distance(a, b)

# Why could not apply LCS, see this example
a = "xxxab"
b = "abyyy"
test_edit_distance(a, b)
