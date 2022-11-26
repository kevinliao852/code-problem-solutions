/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return findAncestor(root, p, q)
}

func findAncestor(root, p, q *TreeNode) *TreeNode {

	if root.Val > p.Val && root.Val > q.Val {
		return findAncestor(root.Left, p, q)
	} else if root.Val < p.Val && root.Val < q.Val {
		return findAncestor(root.Right, p, q)
	}

	return root
}
