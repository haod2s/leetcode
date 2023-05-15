class Solution:
    def isSymmetric(self, root: Optional[TreeNode]) -> bool:
        def check(p: Optional[TreeNode], q: Optional[TreeNode]) -> bool:
            if p == None and q == None: return True
            if p == None or q == None: return False
            return p.val == q.val and check(p.left, q.right) and check(p.right, q.left)
        
        return check(root, root)

# 华为