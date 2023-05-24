class Solution:
    def rob(self, nums: List[int]) -> int:
        n = len(nums)
        if n == 1: return nums[0]
        if n == 2: return max(nums[0], nums[1])
        dp0, dp1 = [0] * n, [0] * n
        dp0[1] = nums[1]
        dp1[0], dp1[1] = nums[0], nums[0]
        for i in range(2, n):
            dp0[i] = max(dp0[i-2] + nums[i], dp0[i-1])
        for i in range(2, n-1):
            dp1[i] = max(dp1[i-2] + nums[i], dp1[i-1])
        
        return max(dp0[n-1], dp1[n-2])

# 华为