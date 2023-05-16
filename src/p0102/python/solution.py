class Solution:
    def levelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        ans = []
        if root == None:
            return ans
        level = []
        q = [root, None]
        while len(q) > 0:
            x = q.pop(0)
            if x == None:
                ans.append(level)
                level = []
                if len(q) != 0: q.append(None)
            else:
            	level.append(x.val)
            	if x.left != None: q.append(x.left)
            	if x.right != None: q.append(x.right)
            
        return ans

# 华为