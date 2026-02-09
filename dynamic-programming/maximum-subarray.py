class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        currentSum=0
        maxSum=-sys.maxsize

        for i in range(len(nums)):
            if currentSum>=0:
                currentSum+=nums[i]
            else:
                currentSum=nums[i]
            maxSum = max(maxSum, currentSum)
        return maxSum
        