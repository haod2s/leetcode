class Solution:
    def isPalindrome(self, s: str) -> bool:
        i, j = 0, len(s) - 1
        while i <= j:
            while i < len(s) and not (s[i].lower() >= 'a' and s[i].lower() <= 'z') and not (s[i] >= '0' and s[i] <= '9'):
                i += 1
            while j >= 0 and not (s[j].lower() >= 'a' and s[j].lower() <= 'z') and not (s[j] >= '0' and s[j] <= '9'):
                j -= 1
            if i > j: break
            if s[i].lower() != s[j].lower():
                return False
            i += 1
            j -= 1 
                
        return True