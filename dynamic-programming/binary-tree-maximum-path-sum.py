# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def findMaxSum(self, root:Optional[TreeNode], pathSum:List) ->int:
        if not root:
            return 0
        
        lSum = max(0, self.findMaxSum(root.left, pathSum))
        rSum = max(0, self.findMaxSum(root.right, pathSum))
        pathSum[0] = max(pathSum[0], lSum+rSum+root.val)
        
        return max(lSum, rSum)+root.val
    
    def maxPathSum(self, root: Optional[TreeNode]) -> int:
        pathSum = [float('-inf')]
        self.findMaxSum(root, pathSum)
        return pathSum[0]

        
        