class Solution:
    def findLHS(self, nums: List[int]) -> int:
        m = dict()
        for v in nums:
            if v not in m:
                m[v] = 1
            else:
                m[v] += 1
        maxLen = 0
        for k, v in m.items():
            if k+1 in m:
                maxLen = max(maxLen, v + m[k+1])
                
        return maxLen