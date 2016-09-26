class Solution(object):
    def removeDuplicates(self, a):
        """
        :type nums: List[int]
        :rtype: int
        """
        if len(a) == 0:
            return 0

        j = 1
        dup_twice = False
        for i in xrange(1, len(a)):
            if a[i] == a[i-1]:
                if dup_twice == False:
                    a[j] = a[i]
                    j += 1
                dup_twice = True
            else:
                a[j] = a[i]
                j += 1
                dup_twice = False
        return j

def test_remove_dup(a):
    print "Before remove {}".format(a)
    n = solution.removeDuplicates(a)
    print "After remove {}".format(a[0:n])

def test_all(all_lists):
    for a in all_lists:
        test_remove_dup(a)

solution = Solution()

a = [[1],
     [1, 2],
     [1, 2, 3],
     [1, 1],
     [1, 1, 1],
     [1, 1, 2, 2, 3],
     [1, 1, 1, 2, 2, 2, 3, 3, 3],
     [1, 1, 1, 1, 2, 2, 2, 3, 3, 3, 3]]
test_all(a)
