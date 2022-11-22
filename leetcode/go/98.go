/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	// be carful of edge cases
	if root.Left == root.Right && root.Left == nil {
		return true
	}

	var isValid bool = true
	var isFirst bool
	var prev int

	traverse(root, &isValid, &prev, &isFirst)

	return isValid
}

func traverse(node *TreeNode, isValid *bool, prev *int, isFirst *bool) {
	if node == nil {
		return
	}

	traverse(node.Left, isValid, prev, isFirst)

	if *isFirst == false {
		*prev = node.Val
		*isFirst = true
	} else {
		if *prev >= node.Val {
			*isValid = false
		} else {
			*prev = node.Val
		}
	}

	traverse(node.Right, isValid, prev, isFirst)
}
