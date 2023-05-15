class Solution:
    def isIsomorphic(self, s: str, t: str) -> bool:
        return True if [s.find(w) for w in s] == [t.find(w) for w in t] else False

# 华为