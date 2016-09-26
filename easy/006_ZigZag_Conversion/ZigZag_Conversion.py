class Solution(object):
    def convert_one_line(self, s, n, nmax):
        round_interval = (nmax - 1) * 2
        if n == 0 or n == (nmax-1): #first line or last line
            step_interval = round_interval
        else:
            step_interval = (nmax - n - 1) * 2

        i, interval = n, step_interval
        line_output = ""
        while i < len(s):
            line_output += s[i]
            i += interval
            if n > 0 and n < (nmax-1): #except first line and last line
                interval = round_interval - interval

        return line_output


    def convert(self, s, numRows):
        """
        :type s: str
        :type numRows: int
        :rtype: str
        """
        if numRows == 1:
            return s

        output = ""
        for i in xrange(0, numRows):
            output += self.convert_one_line(s, i, numRows)
        return output

solution = Solution()
solution.convert("ABCD", 4)
        
