class Solution(object):
    def maxProfit(self, prices):
        """
        :type prices: List[int]
        :rtype: int
        """
        if len(prices) == 0:
            return 0
        minPrice, maxProfit = prices[0], 0
        for price in prices:
            maxProfit = max(maxProfit, price-minPrice)
            minPrice = min(minPrice, price)
        return maxProfit
