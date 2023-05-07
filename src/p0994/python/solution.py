class Solution:
    def orangesRotting(self, grid: List[List[int]]) -> int:
        ans, fresh, q = 0, 0, []
        m, n = len(grid), len(grid[0])
        for i in range(m):
            for j in range(n):
                if grid[i][j] == 2:
                    q.append((i, j))
                elif grid[i][j] == 1:
                    fresh += 1
        if len(q) == 0:
            if fresh > 0: return -1
            return ans
        q.append(None)
        while len(q) > 0:
            pos = q.pop(0)
            if pos == None and len(q) == 0:
                break
            if pos != None:
                if pos[0] > 0 and grid[pos[0]-1][pos[1]] == 1:
                    grid[pos[0]-1][pos[1]] = 2
                    fresh -= 1
                    q.append((pos[0]-1, pos[1]))
                if pos[1] < n - 1 and grid[pos[0]][pos[1]+1] == 1:
                    grid[pos[0]][pos[1]+1] = 2
                    fresh -= 1
                    q.append((pos[0], pos[1]+1))
                if pos[0] < m - 1 and grid[pos[0]+1][pos[1]] == 1:
                    grid[pos[0]+1][pos[1]] = 2
                    fresh -= 1
                    q.append((pos[0]+1, pos[1]))
                if pos[1] > 0 and grid[pos[0]][pos[1]-1] == 1:
                    grid[pos[0]][pos[1]-1] = 2
                    fresh -= 1
                    q.append((pos[0], pos[1]-1))
            else:
                q.append(None)
                ans += 1
        if fresh > 0:
            return -1
  
        return ans