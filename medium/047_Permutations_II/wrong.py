class Solution(object):
    def permute_unique_internal(self, a, start, result):
        if start == len(a):
            result.append(a[:])
            return

        for i in xrange(start, len(a)):
            if i > start and a[i] == a[i-1]:
                continue
            a[start], a[i] = a[i], a[start]
            self.permute_unique_internal(a, start+1, result)
            a[start], a[i] = a[i], a[start]


    def permuteUnique(self, a):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        a.sort()
        start, result = 0, []
        self.permute_unique_internal(a, start, result)
        return result

def test_permute_unique(a):
    result = solution.permuteUnique(a)
    print "Permutation unique for {}".format(a)
    for x in result:
        print x

solution = Solution()
test_permute_unique([3, 1, 1, 2]) # won't pass this case

