class Solution:
    def findTarget(self, root: Optional[TreeNode], k: int) -> bool:
        l = []
        def dfs(n: Optional[TreeNode]):
            if n == None: return
            dfs(n.left)
            l.append(n.val)
            dfs(n.right)
        dfs(root)
        i, j = 0, len(l) - 1
        while i < j:
            if l[i] + l[j] == k:
                return True
            elif l[i] + l[j] > k:
                j -= 1
            else:
                i += 1
                
        return False
    
# 华为