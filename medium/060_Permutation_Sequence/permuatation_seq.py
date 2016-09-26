class Solution(object):
    def make_factorial_tbl(self, n):
        tbl = [0 for i in xrange(n+1)]

        tbl[0] = 1
        for i in xrange(1, n+1):
            tbl[i] = i * tbl[i-1]
        return tbl

    def get_ith_num(self, nums, i):
        k = 0
        for j in xrange(len(nums)):
            if nums[j] == 1:
                continue
            if k == i:
                nums[j] = 1
                return j+1
            k += 1
        return 0

    def get_permutation_internal(self, n, k, tbl, nums, result):
        i = k/tbl[n-1]
        num = self.get_ith_num(nums, i)
        result.append(str(num))

        if n == 1:
            return

        self.get_permutation_internal(n-1, k%tbl[n-1], tbl, nums, result)

    def getPermutation(self, n, k):
        """
        :type n: int
        :type k: int
        :rtype: str
        """
        tbl = self.make_factorial_tbl(n)
        if k > tbl[n] or k <= 0:
            return ""

        nums = [0 for i in xrange(n)]

        result = []
        self.get_permutation_internal(n, k-1, tbl, nums, result)
        return "".join(result)

solution = Solution()

for i in xrange(24):
    result = solution.getPermutation(4, i+1)
    print result
