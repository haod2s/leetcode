/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaves1 := make([]int, 0)
	leaves2 := make([]int, 0)
	leaves1 = leafSimilarHelper(root1, leaves1)
	leaves2 = leafSimilarHelper(root2, leaves2)
	if len(leaves1) != len(leaves2) {
		return false
	}
	for i := 0; i < len(leaves1); i++ {
		if leaves1[i] != leaves2[i] {
			return false
		}
	}
	return true
}

func leafSimilarHelper(root *TreeNode, leaves []int) []int {
	if root == nil {
		return leaves
	}
	if root.Left == nil && root.Right == nil {
		leaves = append(leaves, root.Val)
		return leaves
	}
	leaves = leafSimilarHelper(root.Left, leaves)
	leaves = leafSimilarHelper(root.Right, leaves)
	return leaves
}