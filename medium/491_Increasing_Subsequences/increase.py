class Solution(object):
    def findSubsequences(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        res = self.helper(nums)
        final = []
        for a in res:
            if len(a) > 1:
                final.append(a)
        return final
    
    def helper(self, nums):
        if len(nums) == 0:
            return []
        num = nums[0]
        res = self.helper(nums[1:])
        appeared = set()
        for a in res:
            appeared.add(tuple(a))
        for i in xrange(len(res)):
            if num <= res[i][0]:
                b = [num] + res[i][:]
                if tuple(b) not in appeared:
                    res.append(b)
                    appeared.add(tuple(b))
        res.append([num])
        return res
