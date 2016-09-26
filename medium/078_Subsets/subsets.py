class Solution(object):
    def add_result(self, r, a, mask, n):
        temp = []
        for i in xrange(n):
            if mask[i] == 1:
                temp.append(a[i])
        r.append(temp)

    def combination_internal(self, r, a, mask, start, cnt, n, k):
        if cnt == k:
            self.add_result(r, a, mask, n)
            return

        for i in xrange(start, n):
            mask[i] = 1
            self.combination_internal(r, a, mask, i+1, cnt+1, n, k)
            mask[i] = 0

    def combination(self, a, n, k):
        if k == 0:
            return [[]]
        r = []
        mask = [0 for i in xrange(n)]
        start, cnt = 0, 0
        self.combination_internal(r, a, mask, start, cnt, n, k)
        return r

    def subsets(self, a):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        a.sort()
        result = []
        n = len(a)
        for i in xrange(n+1):
            result += self.combination(a, n, i)
        return result

def test_subsets(a):
    result = solution.subsets(a)
    print a
    print "-"*10
    for s in result:
        print s
    print "\n"

solution = Solution()

test_subsets([5])
test_subsets([5, 2, 4])
