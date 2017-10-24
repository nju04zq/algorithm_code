class Solution(object):
    def largestRectangleArea(self, h):
        """
        :type heights: List[int]
        :rtype: int
        """
        h.append(0)
        n = len(h)
        stack = []
        i, max_area = 0, 0
        while i < n:
            if len(stack) == 0 or h[i] > h[stack[-1]]:
                stack.append(i)
                i += 1
            else:
                last = stack.pop()
                if len(stack) == 0:
                    width = i
                else:
                    width = i - stack[-1] - 1
                max_area = max(max_area, width * h[last])
        h.pop()
        return max_area

def test_largest(a):
    largest = solution.largestRectangleArea(a)
    print "{}, largest {}".format(a, largest)

solution = Solution()

a = [2,1,5,6,2,3, 1, 1, 1, 1]
test_largest(a)

a = []
test_largest(a)

a = [1]
test_largest(a)

a = [2, 1, 5, 6, 2, 3]
test_largest(a)

a = [2, 1, 2]
test_largest(a)
