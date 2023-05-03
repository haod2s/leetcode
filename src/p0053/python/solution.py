class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        pre, maxSum = nums[0], nums[0]
        for i in range(1, len(nums)):
            pre = max(pre + nums[i], nums[i])
            maxSum = max(maxSum, pre)
                
        return maxSum