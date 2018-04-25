class Solution(object):
    def coinChange(self, coins, amount):
        """
        :type coins: List[int]
        :type amount: int
        :rtype: int
        """
        dp = [-1 for i in xrange(amount+1)]
        dp[0], minCnt = 0, -1
        for i in xrange(len(coins)):
            for j in xrange(coins[i], amount+1):
                tmp = dp[j-coins[i]]
                if tmp >= 0:
                    if dp[j] == -1:
                        dp[j] = tmp + 1
                    else:
                        dp[j] = min(tmp+1, dp[j])
            if dp[amount] >= 0:
                if minCnt == -1 or dp[amount] < minCnt:
                    minCnt = dp[amount]
        return minCnt
        

class Solution(object):
    def coinChange(self, coins, amount):
        """
        :type coins: List[int]
        :type amount: int
        :rtype: int
        """
        dp = [amount+1 for i in xrange(amount+1)]
        dp[0] = 0
        for i in xrange(len(coins)):
            for j in xrange(coins[i], amount+1):
                dp[j] = min(dp[j-coins[i]]+1, dp[j])
        if dp[amount] > amount:
            return -1
        else:
            return dp[amount]
