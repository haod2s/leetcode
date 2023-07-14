/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distributeCoins(root *TreeNode) int {
	var dfs func(n *TreeNode) (int, int)
	dfs = func(n *TreeNode) (int, int) {
		if n == nil {
			return 0, 0
		}
		leftCount, leftTimes := dfs(n.Left)
		rightCount, rightTimes := dfs(n.Right)
		count := n.Val + leftCount + rightCount - 1
		times := count
		if times < 0 {
			times *= -1
		}
		return count, times + leftTimes + rightTimes
	}
	_, ans := dfs(root)
	return ans
}