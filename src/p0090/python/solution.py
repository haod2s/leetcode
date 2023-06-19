class Solution:
    def subsetsWithDup(self, nums: List[int]) -> List[List[int]]:
        ans, t = [], []
        def dfs(i: int):
            ans.append(t.copy())
            for j in range(i, len(nums)):
                if j > i and nums[j-1] == nums[j]:
                    continue
                t.append(nums[j])
                dfs(j+1)
                t.pop()
        nums.sort()
        dfs(0)

        return ans