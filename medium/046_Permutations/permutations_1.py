class Solution(object):
    def permute_internal(self, a, mask, path, result):
        if len(path) == len(a):
            result.append(path[:])
            return

        for i in xrange(0, len(a)):
            if mask[i] == 1:
                continue
            mask[i] = 1
            path.append(a[i])
            self.permute_internal(a, mask, path, result)
            path.pop()
            mask[i] = 0

    def permute(self, a):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        result, path = [], []
        mask = [0 for i in xrange(len(a))]
        self.permute_internal(a, mask, path, result)
        return result

def test_permute(a):
    result = solution.permute(a)
    print "Permutation for {}".format(a)
    for x in result:
        print x

solution = Solution()
test_permute([1, 2, 3])
