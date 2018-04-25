class Solution(object):
    def removeKdigits(self, num, k):
        """
        :type num: str
        :type k: int
        :rtype: str
        """
        stack, i = [], 0
        while k > 0 and i < len(num):
            if len(stack) == 0 or stack[len(stack)-1] <= num[i]:
                stack.append(num[i])
                i += 1
            elif stack[len(stack)-1] > num[i]:
                stack = stack[:len(stack)-1]
                k -= 1
        while i < len(num):
            stack.append(num[i])
            i += 1
        while k > 0:
            stack = stack[:len(stack)-1]
            k -= 1
        i = 0
        while i < len(stack) and stack[i] == "0":
            i += 1
        if len(stack[i:]) == 0:
            return "0"
        else:
            return "".join(stack[i:])
