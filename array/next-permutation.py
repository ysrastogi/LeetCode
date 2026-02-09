from typing import List

class Solution:
    def findBreakingPoint(self, nums: List[int]) -> int:
        for i in range(len(nums)-1, 0, -1):
            if nums[i] > nums[i-1]:
                return i-1
        return -1

    def nextPermutation(self, nums: List[int]) -> None:
        if len(nums) <= 1:
            return

        breaking_point = self.findBreakingPoint(nums)

        if breaking_point == -1:
            nums.reverse()
            return

        for i in range(len(nums)-1, breaking_point, -1):
            if nums[i] > nums[breaking_point]:
                nums[i], nums[breaking_point] = nums[breaking_point], nums[i]
                break

        nums[breaking_point+1:] = reversed(nums[breaking_point+1:])
