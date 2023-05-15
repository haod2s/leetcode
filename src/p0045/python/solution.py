class Solution:
    def jump(self, nums: List[int]) -> int:
        maxPos, end, times = -math.inf, 0, 0
        for i in range(len(nums)-1):
            maxPos = max(maxPos, i+nums[i])
            if i == end:
                end = maxPos
                times += 1
        
        return times

# 华为