class Solution:
    def numSubarrayProductLessThanK(self, nums: List[int], k: int) -> int:
        ans, product, i = 0, 1, 0
        for r in range(len(nums)):
            product *= nums[r]
            while i <= r and product >= k:
                product //= nums[i]
                i += 1
            ans += r - i + 1
                
        return ans