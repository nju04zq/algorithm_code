# Given strings S and T, find the minimum (contiguous) substring W of S, so that T is a subsequence of W.
# 
# If there is no such window in S that covers all characters in T, return the empty string "". If there are multiple such minimum-length windows, return the one with the left-most starting index.
# 
# Example 1:
# 
# Input: 
# S = "abcdebdde", T = "bde"
# Output: "bcde"
# Explanation: 
# "bcde" is the answer because it occurs before "bdde" which has the same length.
# "deb" is not a smaller window because the elements of T in the window must occur in order.
# Note:
# 
# All the strings in the input will only contain lowercase letters. The length of S will be in the range [1, 20000]. The length of T will be in the range [1, 100]. 

import sys
import random

def dump(dp, t, s):
    m, n = len(t), len(s)
    sys.stdout.write("    ")
    for ch in s:
        sys.stdout.write("  " + ch)
    sys.stdout.write("\n")
    for i in xrange(m+1):
        if i == 0:
            sys.stdout.write(" ")
        else: 
            sys.stdout.write(t[i-1])
        for j in xrange(n+1):
            sys.stdout.write(" ")
            if dp[i][j] < 0:
                sys.stdout.write("{0}".format(dp[i][j]))
            else:
                sys.stdout.write(" {0}".format(dp[i][j]))
        sys.stdout.write("\n")

# dp[i][j] stores the minmum subsequence which contains t[i-1] and ends at s[j-1]
def minSub(s, t):
    if len(t) == 0 or len(s) == 0:
        return ""
    m, n = len(t), len(s)
    dp = [[0 for i in xrange(n+1)] for i in xrange(m+1) ]
    for i in xrange(1, m+1):
        dp[i][0] = -1
        for j in xrange(1, n+1):
            dp[i][j] = -1
            if dp[i-1][j-1] >= 0 and t[i-1] == s[j-1]:
                dp[i][j] = dp[i-1][j-1] + 1
            if dp[i][j-1] >= 0:
                if dp[i][j] == -1:
                    dp[i][j] = dp[i][j-1] + 1
                else:
                    dp[i][j] = min(dp[i][j], dp[i][j-1] + 1)
    #dump(dp, t, s)
    minLen, minEnd = -1, -1
    for j in xrange(1, n+1):
        if dp[m][j] == -1:
            continue
        if minLen == -1 or dp[m][j] < minLen:
            minLen, minEnd = dp[m][j], j-1
    if minLen == -1:
        return ""
    else:
        return s[minEnd+1-minLen:minEnd+1]

def bf(s, t):
    if len(t) == 0 or len(s) == 0:
        return ""
    minLen, minStart = -1, -1
    for i in xrange(len(s)):
        if s[i] != t[0]:
            continue
        k, j = i, 0
        while k < len(s):
            if s[k] == t[j]:
                j += 1
            if j == len(t):
                if minLen == -1 or k-i+1 < minLen:
                    minLen = k-i+1
                    minStart = i
                break
            k += 1
    if minLen == -1:
        return ""
    else:
        return s[minStart: minStart+minLen]
    return minLen

def testMinSub(s, t):
    ans = bf(s, t)
    res = minSub(s, t)
    print "s {0}, t {1}, get '{2}' '{3}'".format(s, t, ans, res)

def test():
    chs = "abcde"
    tLen = random.randint(1, 5)
    sLen = random.randint(1, 15)
    s, t = "", ""
    for i in xrange(tLen):
        j = random.randint(0, 3)
        t += chs[j]
    for i in xrange(sLen):
        j = random.randint(0, len(chs)-1)
        s += chs[j]
    ans = bf(s, t)
    res = minSub(s, t)
    if ans != res:
        raise Exception("s {0}, t {1}, get '{2}' '{3}'".format(s, t, ans, res))

testMinSub("abcdebdde", "bde")
for i in xrange(10000):
    test()


