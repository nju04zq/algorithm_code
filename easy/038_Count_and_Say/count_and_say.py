class Solution(object):
    def countAndSay(self, n):
        """
        :type n: int
        :rtype: str
        """
        if n <= 0:
            return ""

        say = "1"
        for i in xrange(1, n):
            say = self.analyse(say)
        return say

    def analyse(self, s):
        output = ""

        prev = s[0]
        cnt = 1

        for ch in s[1:]:
            if ch == prev:
                cnt += 1
            else :
                output += "{}{}".format(cnt, prev)
                prev = ch
                cnt = 1

        output += "{}{}".format(cnt, prev)
        return output

for i in xrange(0, 6):
    print "{}, {}".format(i,  Solution().countAndSay(i))
