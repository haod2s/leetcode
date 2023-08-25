/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
	var dfs func(p *TreeNode, prevMax int) int
	dfs = func(p *TreeNode, prevMax int) int {
		if p == nil {
			return 0
		}
		if p.Val >= prevMax {
			return 1 + dfs(p.Left, p.Val) + dfs(p.Right, p.Val)
		}
		return dfs(p.Left, prevMax) + dfs(p.Right, prevMax)
	}
	return 1 + dfs(root.Left, root.Val) + dfs(root.Right, root.Val)
}