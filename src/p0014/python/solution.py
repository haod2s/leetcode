class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        ans = ""
        if len(strs) == 0:
            return ans
        ans = strs[0]
        for s in strs[1:]:
            ans = self.commonPrefix(ans, s)
            
        return ans
            
    
    def commonPrefix(self, s1: str, s2: str) -> str:
        i, j = 0, 0
        while i < len(s1) and j < len(s2):
            if s1[i] != s2[j]:
                break
            i += 1
            j += 1
        
        return s1[0:i]