class Solution(object):
    def pow(self, x, n):
        base = x
        result = 1
        while n > 0:
            if n & 1 == 1:
                result *= base
            n >>= 1
            base *= base
        return result

    def myPow(self, x, n):
        """
        :type x: float
        :type n: int
        :rtype: float
        """
        if n == 0:
            return 1
        if x == 0:
            return 0

        if n < 0:
            negative = True
        else:
            negative = False

        n = abs(n)
        result = self.pow(x, n)
        if negative:
            result = 1/result

        return result

solution = Solution()

x, n = 0, 0
result = solution.myPow(x, n)
print "{}**{}, get {}, should be {}".format(x, n, result, x**n)

x, n = 0, 1
result = solution.myPow(x, n)
print "{}**{}, get {}, should be {}".format(x, n, result, x**n)

x, n = 2, 1
result = solution.myPow(x, n)
print "{}**{}, get {}, should be {}".format(x, n, result, x**n)

x, n = 2, 13
result = solution.myPow(x, n)
print "{}**{}, get {}, should be {}".format(x, n, result, x**n)

