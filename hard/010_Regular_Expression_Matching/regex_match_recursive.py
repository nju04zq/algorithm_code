class Solution(object):
    def is_char_same(self, s, p):
        if s == p:
            return True
        if p == ".":
            return True
        return False

    def is_match_internal(self, s, i, p, j):
        if j >= len(p):
            if i >= len(s):
                return True
            else:
                return False

        if (j+1) >= len(p) or p[j+1] != "*":
            if i == len(s):
                return False
            if self.is_char_same(s[i], p[j]) == False:
                return False
            return self.is_match_internal(s, i+1, p, j+1)

        result = self.is_match_internal(s, i, p, j+2)
        if result:
            return True

        for i in xrange(i, len(s)):
            if self.is_char_same(s[i], p[j]) == False:
                break
            result = self.is_match_internal(s, i+1, p, j+2)
            if result:
                return True

        return False

    def isMatch(self, s, p):
        """
        :type s: str
        :type p: str
        :rtype: bool
        """
        result = self.is_match_internal(s, 0, p, 0)
        return result

def test_is_match(solution, p, s, answer):
    result = solution.isMatch(s, p)
    if result != answer:
        print "P:{} S:{} result {} should {}".format(p, s, result, answer)

solution = Solution()
test_is_match(solution, ".*a*aa*.*b*.c*.*a*", "aabcbcbcaccbcaabc", True)
test_is_match(solution, "c*.*a*", "c", True)
