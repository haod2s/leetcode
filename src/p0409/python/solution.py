class Solution:
    def longestPalindrome(self, s: str) -> int:
        m = dict()
        for c in s:
            if c in m:
                m[c] += 1
            else:
                m[c] = 1
        ans = 0
        for v in m.values():
            ans += v // 2 * 2
            if v % 2 == 1 and ans % 2 == 0:
                ans += 1
                
        return ans