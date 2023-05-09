class Solution:
    def threeSumClosest(self, nums: List[int], target: int) -> int:
        nums.sort()
        ans = math.inf
        for i in range(len(nums)):
            j, k = i + 1, len(nums) - 1
            while j < k:
                s = nums[i] + nums[j] + nums[k]
                if s == target:
                    return target
                if abs(s-target) < abs(ans-target):
                    ans = s
                if s > target: k -= 1
                if s < target: j += 1
                
        return ans