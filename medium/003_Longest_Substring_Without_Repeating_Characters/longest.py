class Solution(object):
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        tbl = {}
        start, maxLen = 0, 0
        for i, ch in enumerate(s):
            if ch in tbl and tbl[ch] >= start:
                start = tbl[ch]+1
            maxLen = max(maxLen, i-start+1)
            tbl[ch] = i
        return maxLen
