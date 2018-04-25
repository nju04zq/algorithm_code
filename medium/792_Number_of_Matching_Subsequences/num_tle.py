class Solution(object):
    def numMatchingSubseq(self, s, words):
        """
        :type S: str
        :type words: List[str]
        :rtype: int
        """
        tbl = {}
        for i, ch in enumerate(s):
            if ch in tbl:
                tbl[ch].append(i)
            else:
                tbl[ch] = [i]
        cnt = 0
        for word in words:
            if self.isSub(word, tbl):
                cnt += 1
        return cnt
    
    def isSub(self, word, tbl):
        cur = -1
        for ch in word:
            if ch not in tbl:
                return False
            j = self.upperbound(tbl[ch], cur)
            if j == -1:
                return False
            cur = j
        return True
    
    def upperbound(self, a, target):
        low, high = 0, len(a)-1
        while low < high:
            mid = low + (high-low)/2
            if a[mid] <= target:
                low = mid+1
            else:
                high = mid
        if low > high or a[low] <= target:
            return -1
        else:
            return a[low]
