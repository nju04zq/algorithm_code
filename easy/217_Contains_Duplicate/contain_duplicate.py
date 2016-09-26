class Solution(object):
    def containsDuplicate(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        if (len(nums) <= 1):
            return False

        nums.sort()
        for i in xrange(1, len(nums)):
            if (nums[i] == nums[i-1]):
                return True
        return False

solution = Solution()

a0 = [1, 2, 3, 4]
result = solution.containsDuplicate(a0)
print "{} contains duplicates, {}".format(a0, result)

a0 = [1, 2, 3, 3]
result = solution.containsDuplicate(a0)
print "{} contains duplicates, {}".format(a0, result)

a0 = [3, 3]
result = solution.containsDuplicate(a0)
print "{} contains duplicates, {}".format(a0, result)
