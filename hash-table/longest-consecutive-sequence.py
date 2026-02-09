class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        if not nums:
            return 0
        s= set(nums)
        currentStreak =0
        longestSequence = 0
        for n in s:
            if not n-1 in s:
                currentNum = n
                currentStreak =1
            
                while currentNum +1 in s:
                    currentNum += 1
                    currentStreak +=1
            longestSequence = max(longestSequence, currentStreak)
        return longestSequence


        