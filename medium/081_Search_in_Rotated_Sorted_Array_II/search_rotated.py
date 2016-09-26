class Solution(object):
    def search(self, a, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: bool
        """
        start, end = 0, len(a)-1
        while start <= end:
            mid = (start + end)/2
            if a[mid] == target or a[start] == target or a[end] == target:
                return True
            if a[start] < a[mid]:
                if a[start] < target < a[mid]:
                    end = mid - 1
                else:
                    start = mid + 1
            elif a[start] > a[mid]:
                if a[mid] < target < a[end]:
                    start = mid + 1
                else:
                    end = mid - 1
            else:
                start += 1
        return False

import random

def generate_seq():
    n = random.randint(0, 20)
    a = [0 for i in xrange(n)]
    for i in xrange(n):
        a[i] = random.randint(0, 12)
    a.sort()
    k = random.randint(0, n)
    b = a + a
    a = b[k: k+n]
    target = random.randint(0, 15)
    return a, target

def test_search(a, target):
    result = solution.search(a, target)
    answer = target in a
    if result != answer:
        print "search {} in {}, result {}, should be {}".format(\
              target, a, result, answer)
        return False
    else:
        return True

solution = Solution()
for i in xrange(1, 1000):
    a, target = generate_seq()
    rc = test_search(a, target)
    if rc == False:
        break
