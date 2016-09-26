# -*- coding: utf-8 -*-
############################################################
# Catalan number
# C(n+1) = ∑(i in [0, n]) C(i) * C(n-i)
# C1 = 1
# C2 = 2
# C3 = 5
# C4 = 14
# http://blog.sina.com.cn/s/blog_6aefe4250101asv5.html
############################################################
# C3 = C0 * C2 + C1 * C1 + C2 * C0
# Ck * Cj,  represents k parentheses in the n+1 parentheses
# C1 = (), C2 = (()), ()()
# C0 * C2 => ()/(()) ()/()()
# C1 * C1 => (())/()
# C2 * C0 => ((())), (()())
class Solution(object):
    def generate_internal(self, l, r, s, result):
        if l == 0 and r == 0:
            result.append(s)
            return

        if l > 0:
            self.generate_internal(l-1, r, s+"(", result)
        if l < r and r > 0:
            self.generate_internal(l, r-1, s+")", result)

    def generateParenthesis(self, n):
        """
        :type n: int
        :rtype: List[str]
        """
        if n <= 0:
            return []

        result = []
        self.generate_internal(n, n, "", result)
        return result

def test_generate(solution, n):
    result = solution.generateParenthesis(n)
    print "##{}/{}##\n{}".format(n, len(result), result)
    print "######"

solution = Solution()
test_generate(solution, 1)
test_generate(solution, 2)
test_generate(solution, 3)
test_generate(solution, 4)
        
