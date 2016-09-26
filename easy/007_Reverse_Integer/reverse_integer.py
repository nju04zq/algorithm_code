class Solution(object):
    def is_overflow(self, i):
        hex_str = hex(i)[2:]
        if len(hex_str) > 8: #32bit, 8bytes
            return True
        if i & 0x80000000 != 0:
            return True
        return False

    def reverse(self, x):
        """
        :type x: int
        :rtype: int
        """
        s = str(x)
        prefix = ""
        if s.isalnum() == False:
            prefix = s[0]
            s = s[1:]

        s = s[::-1]
        i = int(s)

        if self.is_overflow(i):
            return 0

        if prefix == "-":
            return -i
        else:
            return i

print Solution().reverse(1563847412)
