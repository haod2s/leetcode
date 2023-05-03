class Solution:
    def checkRecord(self, s: str) -> bool:
        na = 0
        for c in s:
            if c == 'A':
                na += 1
        
        return na < 2 and 'LLL' not in s