class Solution(object):
    def permute_internal(self, a, start, result):
        if start == len(a):
            result.append(a[:])
            return

        for i in xrange(start, len(a)):
            a[i], a[start] = a[start], a[i]
            self.permute_internal(a, start+1, result)
            a[i], a[start] = a[start], a[i]

    def permute(self, a):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        if len(a) == 0:
            return []

        result = []
        stack = []
        self.permute_internal(a, 0, result)
        return result

def test_case(a):
    result = solution.permute(a)
    print "Permutation for {}, cnt {}\n{}".format(a, len(result), result)

solution = Solution()
test_case([1])
test_case([1, 2, 3])
test_case([1, 2, 3, 4])
