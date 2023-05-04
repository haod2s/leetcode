class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        temp = sorted(intervals)
        ans = [temp[0]]
        for i in range(1, len(temp)):
            if ans[-1][1] < temp[i][0]:
                ans.append(temp[i])
            elif ans[-1][1] >= temp[i][0] and ans[-1][1] <= temp[i][1]:
                ans[-1][1] = temp[i][1]
                
        return ans