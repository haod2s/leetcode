class Solution:
    def pacificAtlantic(self, heights: List[List[int]]) -> List[List[int]]:
        m, n = len(heights), len(heights[0])
        directions = [(-1, 0), (1, 0), (0, -1), (0, 1)]
        def bfs(starts: List[Tuple[int, int]]) -> Set[Tuple[int, int]]:
            q = starts.copy()
            visited = set(starts)
            while len(q) > 0:
                i, j = q.pop(0)
                for d in directions:
                    if 0 <= i + d[0] < m and 0 <= j + d[1] < n and heights[i+d[0]][j+d[1]] >= heights[i][j] and \
                        (i+d[0], j+d[1]) not in visited:
                        q.append((i+d[0], j+d[1]))
                        visited.add((i+d[0], j+d[1]))
                        
            return visited
        
        pacific = [(0, i) for i in range(n)] + [(i, 0) for i in range(1, m)]
        atlantic = [(m - 1, i) for i in range(n)] + [(i, n - 1) for i in range(m - 1)]
        
        return list(map(list, bfs(pacific) & bfs(atlantic)))