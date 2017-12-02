class Solution(object):
    def countBits(self, n):
        cnt = 0
        while n > 0:
            cnt += (n & 0x1)
            n >>= 1
        return cnt

    def hammingWeight(self, n):
        """
        :type n: int
        :rtype: int
        """
        tbl = {}
        for i in xrange(0x100):
            tbl[i] = self.countBits(i)
        cnt = 0
        for i in xrange(4):
            k = n & 0xff
            cnt += tbl[k]
            n >>= 8
        return cnt

print Solution().hammingWeight(11)
