class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        m, n = len(grid), len(grid[0])
        queue = []
        cnt = 0
        directions = [(-1, 0), (1, 0), (0, -1), (0, 1)]
        for i in range(m):
            for j in range(n):
                if grid[i][j] == '1':
                    cnt += 1
                    queue.append((i, j))
                    grid[i][j] = '0'
                    while queue:
                        x, y = queue.pop(0)
                        for d in directions:
                            nextX, nextY = x + d[0], y + d[1]
                            if 0 <= nextX < m and 0 <= nextY < n and grid[nextX][nextY] == '1':
                                queue.append((nextX, nextY))
                                grid[nextX][nextY] = '0'

        return cnt