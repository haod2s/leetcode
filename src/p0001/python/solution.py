class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        m = dict()
        ans = []
        for i, v in enumerate(nums):
            if v in m:
                ans = [i, m[v]]
                break
            else:
                m[target-v] = i
                
        return ans