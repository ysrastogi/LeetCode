class Solution:
    def letsRob(self, ind, nums, dp):
        if ind == 0:
            return nums[0]
        if ind<0:
            return 0
        if dp[ind] != -1:
            return dp[ind]

        rob = nums[ind] + self.letsRob(ind-2, nums, dp)
        dontRob = 0 + self.letsRob(ind-1, nums, dp)
        dp[ind]=max(rob, dontRob)
        return dp[ind]
        
    def rob(self, nums: List[int]) -> int:
        dp = [-1]*(len(nums)+1)
        return self.letsRob(len(nums)-1, nums, dp)
        