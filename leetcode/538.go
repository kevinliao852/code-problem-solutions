/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstToGst(root *TreeNode) *TreeNode {
	var sum int
	var traverseAndPlusSum func(*TreeNode)

	traverseAndPlusSum = func(root *TreeNode) {
		if root == nil {
			return
		}

		traverseAndPlusSum(root.Right)
		sum += root.Val
		root.Val = sum
		traverseAndPlusSum(root.Left)

	}

	traverseAndPlusSum(root)

	return root
}

// same as leetcode 1038
