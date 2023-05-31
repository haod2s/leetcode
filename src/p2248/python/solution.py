class Solution:
    def intersection(self, nums: List[List[int]]) -> List[int]:
        s = set(nums[0])
        for i in range(1, len(nums)):
            s.intersection_update(nums[i])
        ans = [v for v in s]
        ans.sort()
        
        return ans
    
# 腾讯