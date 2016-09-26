class Solution(object):
    def canJump(self, a):
        """
        :type nums: List[int]
        :rtype: bool
        """
        if len(a) == 0:
            return True

        can_jump = False
        i, max_cur = 0, 0
        while i < len(a):
            if max_cur >= (len(a) -1):
                can_jump = True
                break;
            max_next = 0
            while i <= max_cur:
                max_next = max(max_next, i+a[i])
                i += 1
            if max_cur == max_next:
                break
            max_cur = max_next
        return can_jump

solution = Solution()

a = [2,3,1,1,4]
result = solution.canJump(a)
print "{}, can jump {}".format(a, result)

a = [3,2,1,0,4]
result = solution.canJump(a)
print "{}, can jump {}".format(a, result)
        
