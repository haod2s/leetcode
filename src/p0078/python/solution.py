class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        ansList, ans = [], []
        def dfs(i: int):
            if i == len(nums):
                ansList.append(ans.copy())
                return
            ans.append(nums[i])
            dfs(i+1)
            ans.pop()
            dfs(i+1)
        dfs(0)
        
        return ansList