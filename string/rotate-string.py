class Solution:
    def rotateString(self, s: str, goal: str) -> bool:
        s= s+s
        if goal in s and len(goal) is len(s)//2:
            return True
        return False
        
        