class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        m = dict()
        for v in nums:
            if v not in m:
                m[v] = 1
            else:
                m[v] += 1
        for k, v in m.items():
            if len(nums) < 2 * v:
                return k
        return 0