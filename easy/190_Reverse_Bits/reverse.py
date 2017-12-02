class Solution:
    # @param n, an integer
    # @return an integer
    def reverseBits(self, n):
        m = 0 
        for i in xrange(32):
            bit = n & 1
            m = (m << 1) | bit
            n >>= 1
        return m

print Solution().reverseBits(12)
print Solution().reverseBits(13)
