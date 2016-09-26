class Solution(object):
    def generate(self, result, path, a, start, n):
        for i in xrange(start, n):
            path.append(a[i])
            result.append(path[:])
            self.generate(result, path, a, i+1, n)
            path.pop()

    def subsets(self, a):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        a.sort()
        result = []
        path = []
        n = len(a)
        self.generate(result, path, a, 0, n)
        result.append([])
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
