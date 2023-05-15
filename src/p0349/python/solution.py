class Solution:
    def intersection(self, nums1: List[int], nums2: List[int]) -> List[int]:
        ans = []
        s = set(nums1)
        for v in set(nums2):
            if v in s:
                ans.append(v)
                
        return ans
    
# 华为