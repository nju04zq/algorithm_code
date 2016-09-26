class Solution(object):
    def is_pair(self, left, right):
        pair = left + right
        if pair in ["()", "[]", "{}"]:
            return True
        else:
            return False

    def isValid(self, s):
        """
        :type s: str
        :rtype: bool
        """
        stack = []

        for ch in s:
            if ch in "[{(":
                stack.append(ch)
                continue
            if len(stack) == 0:
                return False
            left = stack.pop()
            if self.is_pair(left, ch) == False:
                return False

        if len(stack) > 0:
            return False
        else:
            return True

def test_solution(solution, s):
    print "{} is valid, {}".format(s, solution.isValid(s))

solution = Solution()
test_solution(solution, "");
test_solution(solution, "[]");
test_solution(solution, "[");
test_solution(solution, "]");
test_solution(solution, "[[]");
test_solution(solution, "[[{()}]](){}");

