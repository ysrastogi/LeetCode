class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        buy=0
        sell=0
        profit=0
        maxProfit = float('-inf')
        while(buy<len(prices) and sell<len(prices)):
            profit=prices[sell]-prices[buy]
            maxProfit = max(maxProfit, profit)
            if profit<0:
                buy=sell
            sell+=1
        return maxProfit
        