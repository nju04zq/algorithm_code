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
        i, j = 0, 0
        while i < m:
            if t[i] == s[j]:
                j += 1
            if j == n:
                return True
            i += 1
        return False
