class Solution:
    def minPathSum(self, grid):
        # Initialize the dynamic programming memoization table
        self.dp = [[-1] * len(grid[0]) for _ in range(len(grid))]
        return self.h(grid, 0, 0)
    
    def h(self, grid, i, j):
        # Base case: reached the bottom-right cell
        if i == len(grid) - 1 and j == len(grid[0]) - 1:
            return grid[i][j]
        
        # Out of bounds check (return a large value to prevent this path)
        if i >= len(grid) or j >= len(grid[0]):
            return float('inf')
        
        # If already computed, return memoized result
        if self.dp[i][j] != -1:
            return self.dp[i][j]
        
        # Recursively explore right and down paths
        right = self.h(grid, i, j + 1)
        down = self.h(grid, i + 1, j)
        
        # Memoize and return the minimum path sum
        self.dp[i][j] = grid[i][j] + min(right, down)
        return self.dp[i][j]