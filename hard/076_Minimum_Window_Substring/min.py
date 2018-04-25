class Solution(object):
    def minWindow(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: str
        """
        tbl = {}
        for ch in t:
            tbl[ch] = tbl.get(ch, 0) + 1
        start, counter = 0, len(t)
        minStart, minLen = -1, len(s)+1
        for i in xrange(len(s)):
            ch = s[i]
            if ch not in tbl:
                continue
            tbl[ch] -= 1
            if tbl[ch] >= 0:
                counter -= 1
            while start <= i and counter == 0:
                winLen = i-start+1
                if winLen < minLen:
                    minStart, minLen = start, winLen
                ch_start = s[start]
                if ch_start in tbl:
                    if tbl[ch_start] >= 0:
                        counter += 1
                    tbl[ch_start] += 1
                start += 1
        if minStart == -1:
            return ""
        else:
            return s[minStart:minStart+minLen]
