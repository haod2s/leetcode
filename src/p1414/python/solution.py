class Solution:
    def findMinFibonacciNumbers(self, k: int) -> int:
        if k == 1:
            return 1
        F = [1, 1]
        while F[-2] + F[-1] <= k:
            F.append(F[-2]+F[-1])
        ans = 0
        for i in range(len(F)-1, -1, -1):
            if k - F[i] >= 0:
                ans += 1
                k -= F[i]
        return ans 