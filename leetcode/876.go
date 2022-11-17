/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	middle := head

	for head != nil && head.Next != nil {
		middle = middle.Next
		head = head.Next.Next
	}

	return middle
}
