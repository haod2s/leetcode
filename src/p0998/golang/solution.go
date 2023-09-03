/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	var dfs func(parent, cur *TreeNode, val int)
	dfs = func(parent, cur *TreeNode, val int) {
		if cur == nil {
			parent.Right = &TreeNode{
				Val: val,
			}
			return
		}
		if val > cur.Val {
			parent.Right = &TreeNode{
				Val:  val,
				Left: cur,
			}
			return
		}
		dfs(cur, cur.Right, val)
	}
	if val > root.Val {
		return &TreeNode{
			Val:  val,
			Left: root,
		}
	}
	dfs(root, root.Right, val)
	return root
}