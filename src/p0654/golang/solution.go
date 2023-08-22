/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	maxIdx, maxItem := 0, math.MinInt
	for i := range nums {
		if nums[i] > maxItem {
			maxItem = nums[i]
			maxIdx = i
		}
	}
	left := constructMaximumBinaryTree(nums[:maxIdx])
	right := constructMaximumBinaryTree(nums[maxIdx+1:])
	return &TreeNode{
		Val:   maxItem,
		Left:  left,
		Right: right,
	}
}