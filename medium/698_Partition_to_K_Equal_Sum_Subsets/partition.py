class Solution(object):
    def canPartitionKSubsets(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: bool
        """
        total = sum(nums)
        if total % k != 0:
            return False
        partSum = total/k
        nums.sort(reverse=True)
        mask = [True for i in xrange(len(nums))]
        for i in xrange(k):
            rc = self.partition(nums, mask, 0, partSum)
            if not rc:
                return False
        return True
    
    def partition(self, nums, mask, i, partSum):
        # with the prune if, time cost drops from 607ms to 44ms
        if partSum < 0:
            return False
        elif partSum == 0:
            return True
        elif i == len(nums):
            return False
        if mask[i]:
            mask[i] = False
            rc = self.partition(nums, mask, i+1, partSum-nums[i])
            if rc:
                return True
            else:
                mask[i] = True
        rc = self.partition(nums, mask, i+1, partSum)
        if rc:
            return True
        else:
            return False
