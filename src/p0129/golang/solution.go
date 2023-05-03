/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sumNumbers(root *TreeNode) int {
	return sumNumbersHelper(root, 0)
}

func sumNumbersHelper(root *TreeNode, pre int) int {
	if root == nil {
		return pre
	}
	pre = pre*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return pre
	}
	ret := 0
	if root.Left != nil {
		ret += sumNumbersHelper(root.Left, pre)
	}
	if root.Right != nil {
		ret += sumNumbersHelper(root.Right, pre)
	}
	return ret
}