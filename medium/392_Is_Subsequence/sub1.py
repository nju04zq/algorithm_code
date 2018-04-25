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
        tbl = {}
        for i, ch in enumerate(t):
            if ch in tbl:
                tbl[ch].append(i)
            else:
                tbl[ch] = [i]
        i = -1
        for ch in s:
            if ch not in tbl:
                return False
            j = self.lowerbound(tbl[ch], i)
            if j == -1:
                return False
            i = j
        return True

    def lowerbound(self, a, target):
        low, high = 0, len(a)-1
        while low < high:
            mid = low + (high - low)/2
            if a[mid] <= target:
                low = mid + 1
            else:
                high = mid
        if low > high or a[low] < target:
            return -1
        else:
            return a[low]
