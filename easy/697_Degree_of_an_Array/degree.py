class Solution(object):
    def findShortestSubArray(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        tbl, minLen, degree = {}, len(nums), 0
        for i, num in enumerate(nums):
            if num not in tbl:
                tbl[num] = [i]
            else:
                tbl[num].append(i)
            rangeLen = i-tbl[num][0]+1
            if len(tbl[num]) > degree:
                degree, minLen = len(tbl[num]), rangeLen
            elif len(tbl[num]) == degree:
                minLen = min(minLen, rangeLen)
        return minLen
