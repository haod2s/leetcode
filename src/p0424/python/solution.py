class Solution:
    def characterReplacement(self, s: str, k: int) -> int:
        count = [0] * 26
        i, j= 0, 0
        ans = 0
        while j < len(s):
            count[ord(s[j])-ord('A')] += 1
            j += 1
            if j - i - max(count) > k:
                count[ord(s[i])-ord('A')] -= 1
                i += 1
            ans = max(ans, j - i)	
            
        return ans