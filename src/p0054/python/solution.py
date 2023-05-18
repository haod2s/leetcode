class Solution:
    def spiralOrder(self, matrix: List[List[int]]) -> List[int]:
        m, n = len(matrix), len(matrix[0])
        visited = set()
        ans = []
        i, j = 0, 0
        while len(ans) < m*n:
            while j < n and (i, j) not in visited:
                ans.append(matrix[i][j])
                visited.add((i, j))
                j += 1
            i, j = i + 1, j - 1
            while i < m and (i, j) not in visited:
                ans.append(matrix[i][j])
                visited.add((i, j))
                i += 1
            i, j = i - 1, j - 1
            while j >= 0 and (i, j) not in visited:
                ans.append(matrix[i][j])
                visited.add((i, j))
                j -= 1
            i, j = i - 1, j + 1
            while i >= 0 and (i, j) not in visited:
                ans.append(matrix[i][j])
                visited.add((i, j))
                i -= 1
            i, j = i + 1, j + 1
        
        return ans
    
# 华为