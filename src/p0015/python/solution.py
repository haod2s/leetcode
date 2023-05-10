class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        ans = []
        nums.sort()
        n = len(nums)
        for k in range(n):
            if k > 0 and nums[k] == nums[k-1]: continue
            j = n - 1
            for i in range(k+1, n):
                if i > k+1 and nums[i-1] == nums[i]: continue
                while i < j and nums[k] + nums[i] + nums[j] > 0: j -= 1
                if i == j: break
                if nums[k] + nums[i] + nums[j] == 0:
                    ans.append([nums[k], nums[i], nums[j]])
            
        return ans