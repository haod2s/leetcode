class Solution:
    def majorityElement(self, nums: List[int]) -> List[int]:
        m = dict()
        for n in nums:
            if n not in m:
                m[n] = 1
            else:
                m[n] += 1
        ans = list()
        for k, v in m.items():
            if 3 * v > len(nums):
                ans.append(k)
                
        return ans