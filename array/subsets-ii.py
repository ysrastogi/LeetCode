class Solution:
    def backtrack(self, start, nums, current, result):
        result.append(list(current))
        for i in range(start, len(nums)):
            if i > start and nums[i] == nums[i - 1]:
                continue

            current.append(nums[i])

            self.backtrack(i + 1, nums, current, result)
            current.pop()

    def subsetsWithDup(self, nums):
        nums.sort()  # Sort to handle duplicates
        result = []
        self.backtrack(0, nums, [], result)
        return result