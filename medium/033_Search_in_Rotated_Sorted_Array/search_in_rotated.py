class Solution(object):
    def search(self, a, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """
        start, end = 0, len(a)-1
        while start <= end:
            mid = (start + end)/2
            if a[mid] == target:
                return mid
            # following two are necessary
            if a[start] == target:
                return start
            if a[end] == target:
                return end
            if a[start] < a[mid]:
                if a[start] < target < a[mid]:
                    end = mid - 1
                else:
                    start = mid + 1
            else:
                if a[mid] < target < a[end]:
                    start = mid + 1
                else:
                    end = mid - 1
        return -1

import random

def generate_seq():
    n = random.randint(0, 20)
    target = random.randint(0, 99)
    a = random.sample(xrange(100), n)
    a.sort()
    b = a + a
    k = random.randint(0, n)
    a = b[k:k+n]
    return a, target

def test_search():
    a, target = generate_seq()
    result = solution.search(a, target)
    try:
        answer = a.index(target)
    except:
        answer = -1
    if result != answer:
        print "search {} in {}, result {}, should be {}".format(\
               target, a, result, answer)
        return False
    return True

solution = Solution()
a = [3, 12, 38, 45, 62, 73, 76, 92, 97, 0]
target = 0
result = solution.search(a, target)

for i in xrange(0, 1000):
    rc = test_search()
    if rc == False:
        break
