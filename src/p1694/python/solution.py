class Solution:
    def reformatNumber(self, number: str) -> str:
        temp = "".join([c for c in number if c.isdigit()])
        n = len(temp)
        ans = ''
        if n % 3 == 0:
            for i in range(0, n, 3):
                ans += temp[i:i+3] + '-'
            ans = ans[:-1]
        elif n % 3 == 2:
            for i in range(0, n-2, 3):
                ans += temp[i:i+3] + '-'
            if ans == '-': ans = ans[:-1]
            ans += temp[-2:]
        else:
            for i in range(0, n-4, 3):
                ans += temp[i:i+3] + '-'
            if ans == '-': ans = ans[:-1]
            ans += temp[-4:-2] + '-' + temp[-2:]
        
        return ans