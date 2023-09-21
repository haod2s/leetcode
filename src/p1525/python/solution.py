class Solution:
    def numSplits(self, s: str) -> int:
        ans = 0
        left, right = Counter(), Counter(s)
        leftCnt, rightCnt = len(left), len(right)
        for i in range(0, len(s)):
            if s[i] not in left:
                left[s[i]] = 1
                leftCnt += 1
            if right[s[i]] > 0:
                right[s[i]] -= 1
            if right[s[i]] == 0:
                rightCnt -= 1
            if leftCnt == rightCnt:
            	ans += 1
        return ans