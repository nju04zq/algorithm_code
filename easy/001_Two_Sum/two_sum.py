class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        a = [(nums[i], i) for i in xrange(0, len(nums))]
        a.sort(key=lambda x:x[0])
        i, j = 0, len(nums)-1

        while i < j:
            total = a[i][0] + a[j][0]
            if total == target:
                return [a[i][1], a[j][1]]
            elif total < target:
                i += 1
            else:
                j -= 1

        return [-1, -1]

solution = Solution()

a0 = [1, 2, 3, 50, 52, 55]
print solution.twoSum(a0, 5)
print solution.twoSum(a0, 3)
print solution.twoSum(a0, 55)

a1 = [3, 2, 4]
print solution.twoSum(a1, 6)
