/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode) {
	s := traverseTree(root)
	var x, y *TreeNode
	for i, count := 0, 0; i < len(s)-1 && count < 2; i++ {
		if s[i].Val > s[i+1].Val {
			if count == 0 {
				x, y = s[i], s[i+1]
			} else if count == 1 {
				y = s[i+1]
			}
			count++
		}
	}
	x.Val, y.Val = y.Val, x.Val
}

func traverseTree(root *TreeNode) []*TreeNode {
	if root == nil {
		return []*TreeNode{}
	}
	left := traverseTree(root.Left)
	left = append(left, root)
	right := traverseTree(root.Right)
	return append(left, right...)
}