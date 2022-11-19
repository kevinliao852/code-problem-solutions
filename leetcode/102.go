/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	var traverse func(*TreeNode, int)

	if root == nil {
		return nil
	}

	traverse = func(node *TreeNode, count int) {

		if node == nil {
			return
		}

		if len(res) == count {
			res = append(res, []int{node.Val})
		} else {
			res[count] = append(res[count], node.Val)
		}

		count++
		traverse(node.Left, count)
		traverse(node.Right, count)
	}

	traverse(root, 0)
	return res
}
