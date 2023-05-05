class Solution:
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        ansList = list()
        ans = list()
        def dfs(i: int, remaining: int):
            if i == len(candidates):
                return
            if remaining == 0:
                ansList.append(ans.copy())
                return
            dfs(i+1, remaining)
            if remaining - candidates[i] >= 0:
                ans.append(candidates[i])
                dfs(i, remaining - candidates[i])
                ans.pop()
        dfs(0, target)
        
        return ansList