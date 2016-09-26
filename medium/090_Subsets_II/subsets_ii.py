class Solution(object):
    def subsets_internal(self, a, start, path, result):
        result.append(path[:])
        if len(path) == len(a):
            return

        for i in xrange(start, len(a)):
            if i > start and a[i] == a[i-1]:
                continue
            path.append(a[i])
            self.subsets_internal(a, i+1, path, result)
            path.pop()

    def subsetsWithDup(self, a):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        a.sort()
        start, path, result = 0, [], []
        self.subsets_internal(a, start, path, result)
        return result

def test_subsets(a):
    result = solution.subsetsWithDup(a)
    print "Subsets of {}".format(a)
    for x in result:
        print x

solution = Solution()
test_subsets([1, 2, 2])
test_subsets([1, 2, 2, 3])
