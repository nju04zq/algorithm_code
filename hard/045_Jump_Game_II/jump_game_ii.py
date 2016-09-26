class Solution(object):
    def jump(self, a):
        """
        :type nums: List[int]
        :rtype: int
        """
        i, steps, max_cur = 0, 0, 0
        while i < len(a):
            if max_cur >= len(a)-1:
                break
            max_next = 0
            while i <= max_cur:
                max_next = max(max_next, i+a[i])
                i += 1
            max_cur = max_next
            steps += 1
        return steps

solution = Solution()

a = [0]
steps = solution.jump(a)
print "Jump steps for {}, {}".format(a, steps)

a = [2, 3, 1, 1, 4]
steps = solution.jump(a)
print "Jump steps for {}, {}".format(a, steps)

#    0  1  2  3  4  5  6  7  8  9  a  b
a = [5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0]
steps = solution.jump(a)
print "Jump steps for {}, {}".format(a, steps)
