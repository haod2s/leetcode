class Solution:
    def countPairs(self, coordinates: List[List[int]], k: int) -> int:
        ans = 0
        cnt = Counter()
        for x, y in coordinates:
            for i in range(k+1):
                ans += cnt[i^x, (k-i)^y]
            cnt[x, y] += 1
        return ans