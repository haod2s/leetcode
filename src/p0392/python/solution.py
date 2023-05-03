class Solution:
    def isSubsequence(self, s: str, t: str) -> bool:
        if len(s) == 0 and len(t) == 0:
            return True
        start = 0
        for c in s:
            try:
                i = t.index(c, start)
                start = i + 1
            except ValueError:
                return False
            
        return True