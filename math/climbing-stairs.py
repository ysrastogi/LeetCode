class Solution:
    def climb(self, ind, dp):
        if ind == 0 or ind == 1:  # Base cases for 0 and 1 steps
            return 1
        if dp[ind] != -1:  # If already computed, return the value
            return dp[ind]
        # Correct the use of `ind` instead of `n`
        dp[ind] = self.climb(ind - 1, dp) + self.climb(ind - 2, dp)
        return dp[ind]
    
    def climbStairs(self, n: int) -> int:
        dp = [-1] * (n + 1)  # Initialize dp array with -1
        return self.climb(n, dp)
