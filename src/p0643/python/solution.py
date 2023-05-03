class Solution:
    def findMaxAverage(self, nums: List[int], k: int) -> float:
        preSumK = sum(nums[0:k])
        maxSumK = preSumK
        for i in range(k, len(nums)):
            preSumK = preSumK - nums[i-k] + nums[i]
            maxSumK = max(maxSumK, preSumK)
            
        return maxSumK / k