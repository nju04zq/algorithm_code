class Solution(object):
    do_print = False
    def permute_internal(self, a, start, result):
        if start == len(a):
            result.append(a[:])
            return

        head = {}
        for i in xrange(start, len(a)):
            if a[i] in head:
                continue
            else:
                head[a[i]] = True
            a[i], a[start] = a[start], a[i]
            self.permute_internal(a, start+1, result)
            a[start], a[i] = a[i], a[start]

    def permuteUnique(self, a):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        a.sort()
        if len(a) == 0:
            return []

        result = []
        stack = []
        self.permute_internal(a, 0, result)
        return result

def test_case(a):
    result = solution.permuteUnique(a)
    print "permutation for {}, cnt {}\n{}".format(a, len(result), result)

solution = Solution()
test_case([1])
test_case([1, 1])
test_case([1, 1, 1])
test_case([1, 2, 3])
test_case([1, 2, 2])
test_case([1, 0, 0])
test_case([0, 1, 0, 0, 9])
