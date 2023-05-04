class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        m = dict()
        for v in nums:
            if v not in m:
                m[v] = 1
            else:
                m[v] += 1
        temp = sorted(m.items(), key=lambda x: -x[1])
        ans = list()
        for v in temp[0:k]:
            ans.append(v[0])
            
        return ans