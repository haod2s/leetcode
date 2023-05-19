class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        ans = 0
        w = {}
        i, j = 0, 0
        while j < len(s):
            if s[j] not in w:
                w[s[j]] = 1
            else:
                w[s[j]] += 1
            while s[j] in w and w[s[j]] > 1:
                w[s[i]] -= 1
                i += 1
            ans = max(ans, j-i+1)
            j += 1
        
        return ans
    
# 华为