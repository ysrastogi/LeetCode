class Solution:
    def orangesRotting(self, grid: List[List[int]]) -> int:
        rows, cols = len(grid), len(grid[0])
        queue = deque()  # Queue for BFS
        time = 0

        # Add all initial rotten oranges to the queue
        for i in range(rows):
            for j in range(cols):
                if grid[i][j] == 2:
                    queue.append((i, j, 0))  # (row, col, time)

        # Perform BFS
        while queue:
            row, col, time = queue.popleft()

            # Check all 4 directions
            for dr, dc in [(-1, 0), (1, 0), (0, -1), (0, 1)]:
                r, c = row + dr, col + dc
                if 0 <= r < rows and 0 <= c < cols and grid[r][c] == 1:
                    grid[r][c] = 2  # Mark as rotten
                    queue.append((r, c, time + 1))

        # Check if there are any fresh oranges left
        for row in grid:
            if 1 in row:
                return -1

        return time