class Solution:
    def letterCombinations(self, digits: str) -> List[str]:
        ans = []
        if len(digits) == 0:
            return ans
        m = {
            '2': 'abc',
			'3': 'def',
			'4': 'ghi',
			'5': 'jkl',
			'6': 'mno',
			'7': 'pqrs',
			'8': 'tuv',
			'9': 'wxyz'
		}
        t = []
        def dfs(i):
            if i == len(digits):
                ans.append(''.join(t))
                return
            for c in m[digits[i]]:
                t.append(c)
                dfs(i+1)
                t.pop()
        dfs(0)
        
        return ans