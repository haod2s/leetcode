class Solution:
    def numTilePossibilities(self, tiles: str) -> int:
        used = [False] * len(tiles)
        existed = set()
        t = ""
        def dfs(s: str):
            for j in range(0, len(tiles)):
                if not used[j] and (s + tiles[j]) not in existed:
                    s += tiles[j]
                    existed.add(s)
                    used[j] = True
                    dfs(s)
                    used[j] = False
                    s = s[:len(s)-1]
        dfs(t)
                
        return len(existed)