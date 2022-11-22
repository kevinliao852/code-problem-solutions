/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	var res []int
	var stack []*Node

	if root == nil {
		return res
	}

	stack = append(stack, root)

	for len(stack) > 0 {
		n := stack[len(stack)-1]
		res = append(res, n.Val)
		stack = stack[:len(stack)-1]

		for i := len(n.Children) - 1; i >= 0; i-- {
			stack = append(stack, n.Children[i])
		}
	}

	return res
}
