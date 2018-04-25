class Solution(object):
    def maxSubArray(self, a):
        """
        :type nums: List[int]
        :rtype: int
        """
        if len(a) == 0:
            return 0
        else:
            return self.helper(a)
    
    def helper(self, a):
        if len(a) == 1:
            return a[0]
        low, high = 0, len(a)-1
        mid = low + (high-low)/2
        leftMax, rightMax = None, None
        if low < mid:
            leftMax = self.helper(a[low:mid])
        if mid < high:
            rightMax = self.helper(a[mid+1:high+1])
        res = self.cross(a, mid)
        if leftMax is not None:
            res = max(res, leftMax)
        if rightMax is not None:
            res = max(res, rightMax)
        return res
    
    def cross(self, a, mid):
        total, curMax = 0, 0
        i = mid -1
        while i >= 0:
            total += a[i]
            if total > 0:
                curMax = max(curMax, total)
            i -= 1
        res = a[mid] + curMax
        total, curMax = 0, 0
        i = mid + 1
        while i < len(a):
            total += a[i]
            if total > 0:
                curMax = max(curMax, total)
            i += 1
        return res + curMax
