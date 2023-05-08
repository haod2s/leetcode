class Solution:
    def climbStairs(self, n: int) -> int:
        pre, cur = 1, 2
        if n == 2:
            return cur
        elif n == 1:
            return pre
        i = 2
        while i < n:
            temp = cur
            cur += pre
            pre = temp
            i += 1
            
        return cur