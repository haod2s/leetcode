class Solution:
    def restoreIpAddresses(self, s: str) -> List[str]:
        SEG_NUM = 4
        steps = [1, 2, 3]
        ans, ip = [], []
        def dfs(segId: int, i: int):
            if segId == SEG_NUM and i == len(s):
                ans.append('.'.join(ip))
                return
            if (segId == SEG_NUM and i < len(s)) or (segId <= SEG_NUM and i >= len(s)):
                return
            for d in steps:
                item = s[i:i+d]
                if len(item) > 1 and item[0] == '0':
                    return
                if int(item) > 255:
                    return
                ip.append(s[i:i+d])
                dfs(segId+1, i+d)
                ip.pop()
        dfs(0, 0)
        
        return ans