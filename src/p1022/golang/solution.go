/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumRootToLeaf(root *TreeNode) int {
	var (
		ans int
		dfs func(root *TreeNode, n int)
	)
	dfs = func(root *TreeNode, n int) {
		if root.Left == nil && root.Right == nil {
			ans += n*2 + root.Val
			return
		}
		if root.Left != nil {
			dfs(root.Left, n*2+root.Val)
		}
		if root.Right != nil {
			dfs(root.Right, n*2+root.Val)
		}
	}
	dfs(root, 0)
	return ans
}