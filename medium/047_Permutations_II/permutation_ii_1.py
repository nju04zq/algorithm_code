class Solution(object):
    def permute_unique_internal(self, a, mask, path, result):
        if len(path) == len(a):
            result.append(path[:])
            return

        prev = None
        for i in xrange(0, len(a)):
            if mask[i] == 1:
                continue
            if prev is None:
                prev = a[i]
            elif a[i] == prev:
                continue
            prev = a[i]

            mask[i] = 1
            path.append(a[i])
            self.permute_unique_internal(a, mask, path, result)
            path.pop()
            mask[i] = 0

    def permuteUnique(self, a):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        a.sort()
        path, result = [], []
        mask = [0 for i in xrange(len(a))]
        self.permute_unique_internal(a, mask, path, result)
        return result

def test_permute_unique(a):
    result = solution.permuteUnique(a)
    print "Permutation unique for {}".format(a)
    for x in result:
        print x

solution = Solution()
test_permute_unique([3, 1, 1, 2])

