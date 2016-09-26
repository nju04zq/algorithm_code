class Solution(object):
    def find_last_breaker(self, a, n):
        for i in xrange(n-1, 0, -1):
            if a[i] > a[i-1]:
                return i-1
        return -1

    def reverse(self, a, start, end):
        i, j = start, end
        while i < j:
            a[i], a[j] = a[j], a[i]
            i += 1
            j -= 1

    def find_next_larger(self, a, target, start, end):
        for i in xrange(start, end+1):
            if a[i] > target:
                return i
        return None

    def nextPermutation(self, a):
        """
        :type nums: List[int]
        :rtype: void Do not return anything, modify nums in-place instead.
        """
        n = len(a)
        if n == 0:
            return
        i = self.find_last_breaker(a, n)
        self.reverse(a, i+1, n-1)
        if i == -1:
            return
        j = self.find_next_larger(a, a[i], i+1, n-1)
        a[i], a[j] = a[j], a[i]

def test_next_perm(a):
    print "before do next: {}".format(a)
    solution.nextPermutation(a)
    print "after do next:  {}".format(a)

solution = Solution()
test_next_perm([5, 4, 3, 2, 1])
test_next_perm([2, 2, 3, 1, 5])
test_next_perm([2, 2, 3, 5, 1])
test_next_perm([2, 2, 5, 3, 1])

